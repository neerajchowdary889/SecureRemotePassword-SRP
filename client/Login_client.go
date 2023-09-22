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
	u := NG_values.H(append(user_tempdetails.A.Bytes(), B.Bytes()...))
	fmt.Println("Client_u: ",u)
	return u
}

func (user *ClientDetails) Compute_K_client(B *big.Int, u string) (string){
	// S = (B - kg^x) ^ (a + ux) (mod N)

	var Password string
	fmt.Println("Enter Password: ")
	fmt.Scan(&Password)

	x, _ := user.computeX(Password, false)

	fmt.Println(x)

	return "Done..."
}