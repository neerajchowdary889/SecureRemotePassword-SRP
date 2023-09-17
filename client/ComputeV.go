package client

import(
	// "fmt"
	"srp/NG_values"
	"math/big"
	"encoding/binary"
)


func(user *ClientDetails) computeK(tempdetails *ClientTempDetails, password string) bool{
	
	Nbytes := tempdetails.N.Bytes()
	Gbytes := new(big.Int).SetUint64(uint64(user.G)).Bytes()

	K := NG_values.H(append(Nbytes,Gbytes...))
	user.K = K
	// fmt.Printf("K:\n %v\n",K)
	status := user.computeX(tempdetails, password)
	return status
}


func(user *ClientDetails) computeX(tempdetails *ClientTempDetails, password string) bool{
	// x = H(s | H ( I | ":" | p) )

	x := NG_values.H([]byte(user.Username + ":" + password))
	saltBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(saltBytes, user.Salt)
	X := NG_values.H(append(saltBytes, []byte(x)...))
	// fmt.Println("X: ",X)

	status := user.computeV(tempdetails, X)
	return status

}


func (user *ClientDetails) computeV(tempdetails *ClientTempDetails, X string) bool{
	// v = g^x (mod N)
	x := new(big.Int)
	x.SetString(X, 16)

	g := big.NewInt(int64(user.G)) 
	v := new(big.Int).Exp(g, x, tempdetails.N)
	user.V = v
	// fmt.Printf("v = g^x (mod N) = %s\n", v.String())
	return true
}