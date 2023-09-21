package client

import(
	"math/big"
)

func FromServer(Map map[string]interface{}) (*ClientDetails){

	user := &ClientDetails{
		Username: Map["Username"].(string),
		Salt: Map["Salt"].(uint64),
		G: Map["G"].(uint8),
		K: Map["K"].(string),
		V: Map["V"].(*big.Int),
		N: Map["N"].(*big.Int),
	}

	return user
}