package client

import (
	"fmt"
	"math/big"
	"srp/NG_values"
	"encoding/binary"
)

func (user *ClientDetails) GenerateA() (*ClientTempDetails) {

	a := NG_values.Generate64BitNumber()

	A := new(big.Int)
	A = A.Exp(new(big.Int).SetUint64(uint64(user.G)), new(big.Int).SetUint64(a), user.N)
	   
    aBytes := make([]byte, 8)
    binary.LittleEndian.PutUint64(aBytes, a)

	ClientTempDetails := &ClientTempDetails{
		A: A,
		a: a,
	}

	return ClientTempDetails
}

func (user_tempdetails *ClientTempDetails) Client_ComputeU(B *big.Int) (string){
	// u = H(A | B)
	// fmt.Println("Client_A: ",user_tempdetails.A,"\n")
	// fmt.Printf("Client_a: %T\n",user_tempdetails.A)
	// fmt.Println("Client_B: ",B,"\n")
	// fmt.Printf("Client_B: %T\n",B)
	u := NG_values.H(append(user_tempdetails.A.Bytes(), B.Bytes()...))
	fmt.Println("Client_u: ",u)
	return u
}