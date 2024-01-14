package Nodes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"srp/Execution"
	"sync"
    "math/big"
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
			status := Login_101(ws, result.Message.(string))
                if !status{
                    ws.Close()
                }
		}

        // // Write message back to browser
        // if err = ws.WriteMessage(msgType, msg); err != nil {
        //     fmt.Println(err)
        //     return
        // }
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

func Login_101(ws *websocket.Conn, username string)(bool){
	fmt.Println("Login_101")

	ServerStoringDetails, status := Execution.Login(username)

	if status == false {
		SendTupleToClient(ws, Tuple{Status: 409, Message: "Username not found --> Dropping Connection"})
		return false
	}else{
		// Map := StructToMap(ServerStoringDetails)
		Map, err := json.Marshal(ServerStoringDetails)
        Map_str := string(Map)
		if err != nil{
			SendTupleToClient(ws, Tuple{Status: 409, Message: "conversion went wrong --> Dropping Connection"})
			return false
		}else{
            server_tempdetails := ServerStoringDetails.GenerateB()
            data := DataTuple{Message: Map_str, Metadata: server_tempdetails.B}
            dataJson, err := json.Marshal(data)
            if err != nil {
                log.Println(err)
                return false
            }
            dataStr := string(dataJson)
			SendTupleToClient(ws, Tuple{Status: 200, Message: dataStr})
			return true
		}

	}
}

func Server_Execution() {
    StartWebSocketServer()
}