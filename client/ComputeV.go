package client

import(
	// "fmt"
	"srp/NG_values"
	"math/big"
	"encoding/binary"
)


func(user *ClientDetails) computeK(password string) bool{
	
	Nbytes := user.N.Bytes()
	Gbytes := new(big.Int).SetUint64(uint64(user.G)).Bytes()

	K := NG_values.H(append(Nbytes,Gbytes...))
	user.K = K
	// fmt.Printf("K:\n %v\n",K)
	_,status := user.computeX(password, true)
	return status
}


func(user *ClientDetails) computeX(password string, val bool) (string, bool){
	// x = H(s | H ( I | ":" | p) )
	
	x := NG_values.H([]byte(user.Username + ":" + password))
	saltBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(saltBytes, user.Salt)
	X := NG_values.H(append(saltBytes, []byte(x)...))
	status := user.computeV(X)

	if val{
	return "", status
	}else{
		return X, status
	}
	
}


func (user *ClientDetails) computeV(X string) bool{
	// v = g^x (mod N)
	x := new(big.Int)
	x.SetString(X, 16)

	g := big.NewInt(int64(user.G)) 
	v := new(big.Int).Exp(g, x, user.N)
	user.V = v

	return true
}