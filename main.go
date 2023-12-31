package main

import (
	"fmt"
	"os"
	"srp/NG_values"
	"srp/client"
	"srp/server"
	"time"

	"github.com/briandowns/spinner"
	// "math/big"
)

func SaltandNG_generation(user *client.ClientDetails) bool {

	fmt.Println(">>> N and G value Generation")

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()

	N := NG_values.GenerateN()
	G := NG_values.GenerateG(N)

	s.Stop()

	if N != nil || G != 0 {
		user.N = N
		user.G = G
		salt := NG_values.Generate64BitNumber()
		user.Salt = salt
		return true
	} else {
		fmt.Println("Error: N or G is nil")
	}
	return false
}

func DisplayDetails(user *client.ClientDetails) {
	fmt.Println("Do you want to check the details before sending? (y/n)")
	var choice string
	fmt.Scan(&choice)
	if choice == "y" {
		user.ReadDetails()
		return
	} else {
		return
	}
}

func sendToserver(user *client.ClientDetails) {
	fmt.Println(">>> Sending to server")
	details := user.SendToServer()
	status := server.UserSignUp(details)
	if !status {
		fmt.Println("Error: User details not sent to server")
	} else {
		fmt.Println("User details sent to server...")
	}
}

func login() (*server.ServerStoringDetails, bool) {
	fmt.Print("Type out your username: ")
	var username string
	fmt.Scan(&username)
	FromServer, status := server.Searchcsv(username)
	if !status {
		return nil, false
	} else {
		fmt.Println("User found...")
		return FromServer, true
	}
}

func checkpermission(str string) bool {
	fmt.Printf("%v\n>>> ",str)
	var choice string
	fmt.Scan(&choice)
	if choice == "y" {
		return true
	} else {
		return false
	}
}


func main() {

	if checkpermission(`>>> Do you wan to Sign Up? (y/n) 
Note: If you're trying to signup your pc fans might kick in for few secs dont worry. 
Note: If an error occurs, please try again. The error might be due to the number not being found within 15 seconds.`) {
		user := &client.ClientDetails{}

		status := SaltandNG_generation(user)
		if !status {
			fmt.Println("Error: Salt and NG values not generated")
			os.Exit(1)
		}

		user.GenerateUsernamePassword()
		DisplayDetails(user)

		sendToserver(user)
	}

	// ---------------------------------------------------------------------

	if checkpermission(">>> Do you want to login? (y/n)") {
		ServerStoringDetails, status := login()
		if !status {
			fmt.Println("Error: User not found, Register First.")
		} else {
			Credentials := ServerStoringDetails.SendToClient()
			user := client.FromServer(Credentials)

			user_tempdetails := user.GenerateA()
			server_tempdetails := ServerStoringDetails.GenerateB()
			
			u_client := user_tempdetails.Client_ComputeU(server_tempdetails.B)
			u_server := server_tempdetails.Server_ComputeU(user_tempdetails.A)

			if u_client == u_server {
				user.Compute_K_client(user_tempdetails)
				server_tempdetails.Compute_K_server(ServerStoringDetails)
				M1 := user.GenerateM1(user_tempdetails)
				M2 := ServerStoringDetails.GenerateM2(server_tempdetails, M1)
				M := user.GenerateM(user_tempdetails, M1)

				if M == M2{
					fmt.Println("Permission Granted")
				}else{
					fmt.Println("Permission Denied")
				}
			}

		}
	}

}
