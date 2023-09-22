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
	N *big.Int
}

type TempServerDetails struct{
	A *big.Int
	B *big.Int
	b uint64
	u string
}

func(Server_user *ServerStoringDetails)PrintDetails()(bool){
	fmt.Println("Server Details:")
	fmt.Printf("Username: %v\n", Server_user.Username)
	fmt.Printf("Salt: %v\n", Server_user.Salt)
	fmt.Printf("G: %v\n", Server_user.G)
	fmt.Printf("K: %v\n", Server_user.K)
	fmt.Printf("V: %v\n", Server_user.V)
	fmt.Printf("N: %v\n", Server_user.N)
	return true
}