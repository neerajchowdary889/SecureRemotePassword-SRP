package client

import (
	"fmt"
	"srp/NG_values"
)

func(user *ClientDetails) GenerateUsernamePassword(){
	var username,password string
	fmt.Println(">>> Username and Password Generation")
	fmt.Print("Enter Username: ")
	fmt.Scan(&username)
	user.Username = NG_values.H(username)
	fmt.Print("Enter Password: ")
	fmt.Scan(&password)
	user.computeK(password)
}