package Nodes

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"srp/Execution"
	"srp/server"
	"sync"
	"github.com/gorilla/websocket"
)
var mu sync.Mutex
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}
var status_101 bool = false
var status_201 bool = false
var status_301 bool = false
var status_401 bool = false
var M2 string
var ServerStoringDetails *server.ServerStoringDetails
var server_tempdetails *server.TempServerDetails
type Tuple struct {
    Status uint16 `json:"status"`
    Message interface{} `json:"message"`
}

type DataTuple struct{
    Message string `json:"message"`
    Metadata *big.Int `json:"metadata"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
    }

    defer ws.Close()

    for {
        // Read message from browser
        _, msg, err := ws.ReadMessage()
        if err != nil {
            fmt.Println(err)
            return
        }
		var result Tuple
		err = json.Unmarshal(msg, &result)
		if err != nil {
			fmt.Println("Error during Unmarshal:", err)
			return
		}

        // Print the message to the console
        fmt.Printf("%s sent: %T --> %T", ws.RemoteAddr(), result.Status, result.Message)

		if result.Status == 101{
			fmt.Println("Received 101")
			status_101, ServerStoringDetails, server_tempdetails = Login_101(ws, result.Message.(string))
                if !status_101{
                    ws.Close()
                }
                fmt.Printf("%+v\n", ServerStoringDetails)
                fmt.Printf("%+v\n", server_tempdetails)
		}

        if result.Status == 201 {

            if status_101{
                fmt.Println("Received 201 --> Generate U")
                if ServerStoringDetails == nil || server_tempdetails == nil {
                    fmt.Println("ServerStoringDetails or server_tempdetails is nil")
                    SendTupleToClient(ws, Tuple{409, "Nil Structs: error"})
                    ws.Close()
                }
                A_ := result.Message.(string)
                A := new(big.Int)
                _, ok := A.SetString(A_, 10) // 10 is the base
                if !ok {
                    fmt.Println("SetString: error")
                    SendTupleToClient(ws, Tuple{409, "SetString: error"})
                    ws.Close()
                }
                server_tempdetails.A = A
                fmt.Printf("Server_tempdetails.A: %v\n", server_tempdetails.A)
                fmt.Printf("Server_tempdetails.B: %v\n", server_tempdetails.B)
                B_u, err := Login_201(ws, ServerStoringDetails, server_tempdetails)
                if !err{
                    fmt.Println("U: error")
                    SendTupleToClient(ws, Tuple{409, "U: error"})
                    ws.Close()
                }
                status_201 = true
                server.SetU(server_tempdetails, B_u)
                fmt.Printf("Server_tempdetails.u: %v\n", server.GetU(server_tempdetails))
            }else{
            SendTupleToClient(ws, Tuple{409, "Bad req"})
            }
        }

        if result.Status == 301{
            if status_101 && status_201{
                fmt.Println("Received 301 --> Generate K")
                if ServerStoringDetails == nil || server_tempdetails == nil {
                    fmt.Println("ServerStoringDetails or server_tempdetails is nil")
                    SendTupleToClient(ws, Tuple{409, "Nil Structs: error"})
                    ws.Close()
                }
                status := server_tempdetails.Compute_K_server(ServerStoringDetails)
                if !status{
                    fmt.Println("K: error")
                    SendTupleToClient(ws, Tuple{409, "K: error"})
                    ws.Close()
                }
                status_301 = true
                SendTupleToClient(ws, Tuple{Status: 301, Message: true})
                fmt.Printf("K --> %s", server_tempdetails.K_server)
            }else{
                SendTupleToClient(ws, Tuple{409, "Bad req"})
            }
        }

        if result.Status == 401{
            if status_101 && status_201 && status_301{
                fmt.Println("Received 401 --> Verifying")
                if ServerStoringDetails == nil || server_tempdetails == nil {
                    fmt.Println("ServerStoringDetails or server_tempdetails is nil")
                    SendTupleToClient(ws, Tuple{409, "Nil Structs: error"})
                    ws.Close()
                }
                M_1 := result.Message.(string)
                // fmt.Printf("%+v\n", ServerStoringDetails)
                // fmt.Printf("%+v\n", server_tempdetails)

                // fmt.Println("M_1 ---> ", M_1)
                // fmt.Println("A ---> ", server_tempdetails.A)
                M2 = ServerStoringDetails.GenerateM2(server_tempdetails, M_1)
                fmt.Println("M2 ---> ", M2)
                status_401 = true
                SendTupleToClient(ws, Tuple{Status: 401, Message: true})
            }else{
                SendTupleToClient(ws, Tuple{Status: 409, Message: "Bad req"})
            }
        }

        if result.Status == 501{
            if status_101 && status_201 && status_301 && status_401{
                fmt.Println("Received 501 --> Final Verification")
                if ServerStoringDetails == nil || server_tempdetails == nil {
                    fmt.Println("ServerStoringDetails or server_tempdetails is nil")
                    SendTupleToClient(ws, Tuple{409, "Nil Structs: error"})
                    ws.Close()
                }
                M := result.Message.(string)
                // fmt.Printf("M: %s\n", M)
                // fmt.Printf("M2: %s\n", M2)
                if M == M2{
                    SendTupleToClient(ws, Tuple{Status: 501, Message: true})
                }else{
                    SendTupleToClient(ws, Tuple{Status: 409, Message: false})
                }
            }
        }
	}
}	
func StartWebSocketServer() {
    // WEB-SOCKETS
    http.HandleFunc("/ws", handleConnections)

    // Start listening on localhost:2004
    fmt.Println("WebSocket server started on :2002")
    err := http.ListenAndServe(":2002", nil)
    if err != nil {
        fmt.Println("ListenAndServe: ", err)
    }
    
}

func SendTupleToClient(ws *websocket.Conn, tuple Tuple) {
    // Serialize the tuple into JSON
    tupleJson, err := json.Marshal(tuple)
    if err != nil {
        log.Println(err)
        return
    }

    // Send the JSON over the WebSocket connection
	mu.Lock()
	defer mu.Unlock()
    if err := ws.WriteMessage(websocket.TextMessage, tupleJson); err != nil {
        log.Println(err)
        return
    }
}

func Login_101(ws *websocket.Conn, username string)(bool, *server.ServerStoringDetails, *server.TempServerDetails){
	fmt.Println("Login_101")

	ServerStoringDetails, status := Execution.Login(username)

	if status == false {
		SendTupleToClient(ws, Tuple{Status: 409, Message: "Username not found --> Dropping Connection"})
		return false, nil, nil
	}else{
		// Map := StructToMap(ServerStoringDetails)
		Map, err := json.Marshal(ServerStoringDetails)
        Map_str := string(Map)
		if err != nil{
			SendTupleToClient(ws, Tuple{Status: 409, Message: "conversion went wrong --> Dropping Connection"})
			return false, ServerStoringDetails, nil
		}else{
            server_tempdetails := ServerStoringDetails.GenerateB()
            data := DataTuple{Message: Map_str, Metadata: server_tempdetails.B}
            dataJson, err := json.Marshal(data)
            if err != nil {
                SendTupleToClient(ws, Tuple{Status: 409, Message: "Json marshal error --> Dropping Connection"})
                return false, ServerStoringDetails, nil
            }
            dataStr := string(dataJson)
			SendTupleToClient(ws, Tuple{Status: 200, Message: dataStr})
			return true, ServerStoringDetails, server_tempdetails
		}
	}
}

func Login_201(ws *websocket.Conn, ServerStoringDetails *server.ServerStoringDetails, server_tempdetails *server.TempServerDetails)(string, bool){
    U := server_tempdetails.Server_ComputeU(server_tempdetails.A)
    SendTupleToClient(ws, Tuple{Status: 201, Message: true})
    return U, true
}

func Server_Execution() {
    StartWebSocketServer()
}