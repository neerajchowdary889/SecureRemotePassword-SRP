package main

import (
	"fmt"
	"srp/NG_values"
	"srp/client"
	"srp/server"
	"os"
	"time"
	"github.com/briandowns/spinner"
	// "math/big"
)

func SaltandNG_generation(user *client.ClientDetails, tempdetails *client.ClientTempDetails) bool{

	fmt.Println(">>> N and G value Generation")

    s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
    s.Start()

	N := NG_values.GenerateN()
	G := NG_values.GenerateG(N)

	s.Stop()
	
	if N != nil || G != 0{
		tempdetails.N = N
		user.G = G
		salt := NG_values.GenerateSalt()
		user.Salt = salt
		return true
	}else{
		fmt.Println("Error: N or G is nil")
	}
	return false
}

func DisplayDetails(user *client.ClientDetails, tempdetails *client.ClientTempDetails){
	fmt.Println("Do you want to check the details before sending? (y/n)")
	var choice string
	fmt.Scan(&choice)
	if choice == "y"{
		user.ReadDetails(tempdetails)
		return
	}else{
		return
	}
}

func sendToserver(user *client.ClientDetails){
	fmt.Println(">>> Sending to server")
	details := user.SendToServer()
	status := server.UserSignUp(details)
	if !status{
		fmt.Println("Error: User details not sent to server")
	}else{
		fmt.Println("User details sent to server...")
	}
}

func main(){
	user := &client.ClientDetails{}
	tempdetails := &client.ClientTempDetails{}

	status := SaltandNG_generation(user, tempdetails)
	if !status{
		fmt.Println("Error: Salt and NG values not generated")
		os.Exit(1)
	}

	user.GenerateUsernamePassword(tempdetails)
	DisplayDetails(user, tempdetails)

	sendToserver(user)
}
