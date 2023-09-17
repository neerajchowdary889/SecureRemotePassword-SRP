package main

import (
	"fmt"
	"srp/NG_values"
	"srp/client"
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

func main(){
	user := &client.ClientDetails{}
	tempdetails := &client.ClientTempDetails{}

	status := SaltandNG_generation(user, tempdetails)
	if !status{
		fmt.Println("Error: Salt and NG values not generated")
		os.Exit(1)
	}
	user.GenerateUsernamePassword(tempdetails)
	user.ReadDetails(tempdetails)
}
