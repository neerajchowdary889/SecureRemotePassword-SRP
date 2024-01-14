package client
import(
	"math/big"
	"fmt"
)

type ClientDetails struct{
	Username string
	Salt uint64
	G uint8
	K string
	V *big.Int
	N *big.Int
}

type ClientTempDetails struct{
	A *big.Int
	B *big.Int
	a uint64
	u string
	K_client string
}

func(user *ClientDetails) ReadDetails(){

	fmt.Println("Client Details:")
    fmt.Printf("Username: %v\n", user.Username)
    fmt.Printf("Salt: %v\n", user.Salt)
    fmt.Printf("G: %v\n", user.G)
    fmt.Printf("K: %v\n", user.K)
    fmt.Printf("V: %v\n", user.V)
	
	fmt.Println("\nClient Temp Details:")
    fmt.Printf("N: %v\n", user.N)


}