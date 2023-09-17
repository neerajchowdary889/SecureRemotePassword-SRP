package client

import (
	"fmt"
)

func(user *ClientDetails) GenerateUsernamePassword(tempdetails *ClientTempDetails){
	var username,password string
	fmt.Println(">>> Username and Password Generation")
	fmt.Print("Enter Username: ")
	fmt.Scan(&username)
	user.Username = username
	fmt.Print("Enter Password: ")
	fmt.Scan(&password)
	status := user.computeK(tempdetails, password)
	fmt.Println(status)
}