package server

import (
	"fmt"
	"math/big"
)
type ServerStoringDetails struct{
	Username string
	Salt uint64
	G uint8
	K string
	V *big.Int
}

type TempServerDetails struct{

}

func(Server_user *ServerStoringDetails)ServerStoredDetails()(bool){
	fmt.Println("Server Details:")
	fmt.Printf("Username: %v\n", Server_user.Username)
	fmt.Printf("Salt: %v\n", Server_user.Salt)
	fmt.Printf("G: %v\n", Server_user.G)
	fmt.Printf("K: %v\n", Server_user.K)
	fmt.Printf("V: %v\n", Server_user.V)
	return true
}