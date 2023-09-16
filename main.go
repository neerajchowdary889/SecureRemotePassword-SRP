package main

import(
	"fmt"
	"srp/NG_values"
)
func SaltandNG_generation(){

	fmt.Println(">>> N and G value Generation")
	N,G := NG_values.GenerateNG()

	if N != nil && G != 0{
		fmt.Printf("--->N:\n %v\n--->G:\n %v\n",N,G)
		salt := NG_values.GenerateSalt()
		fmt.Printf("--->salt:\n %v\n",salt)
	}else{
		fmt.Println("Error: N or G is nil")
	}

}
func main(){
	fmt.Println("Hello World")
	SaltandNG_generation()
}