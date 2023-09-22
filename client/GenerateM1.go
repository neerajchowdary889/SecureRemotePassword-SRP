package client

import(
	// "fmt"
	"math/big"
	"encoding/binary"
	"srp/NG_values"
)

func (user *ClientDetails) GenerateM1(user_tempdetails *ClientTempDetails) string{
	// M1 = H(H(N) xor H(g) | H(I) | s | A | B | K_client)
	HN := NG_values.H(user.N.Bytes())
	Hg := NG_values.H([]byte(string(user.G)))
	HI := NG_values.H([]byte(user.Username))

	saltBytes := make([]byte, 8)
    binary.BigEndian.PutUint64(saltBytes, user.Salt)

	HN_xor_Hg := XOR([]byte(HN), []byte(Hg))

	M1 := NG_values.H(append(HN_xor_Hg, append([]byte(HI), append(saltBytes, append(user_tempdetails.A.Bytes(), append(user_tempdetails.B.Bytes(), []byte(user_tempdetails.K_client)... )...)...)...)...))

	return M1

}
func XOR(a, b []byte) []byte {
    if len(a) != len(b) {
        panic("XOR: lengths of inputs do not match")
    }
    result := make([]byte, len(a))
    for i := 0; i < len(a); i++ {
        result[i] = a[i] ^ b[i]
    }
    return result
}

func (user *ClientDetails) GenerateM(user_tempdetails *ClientTempDetails, M_1 string) string{

	M1 := new(big.Int)
	M1.SetString(M_1, 16)

	K_client := new(big.Int)
	M1.SetString(user_tempdetails.K_client, 16)

	M2 := NG_values.H(append(user_tempdetails.A.Bytes(), append(M1.Bytes(), K_client.Bytes()...)...))
	return M2

}