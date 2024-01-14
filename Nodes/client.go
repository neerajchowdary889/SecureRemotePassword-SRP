package Nodes

import (
	"encoding/json"
	// "fmt"
	"math/big"
	"net/http"
	"reflect"
	"srp/Execution"
	"srp/NG_values"
	"srp/client"
	"srp/server"

	"github.com/gin-gonic/gin"
)

type Map_Client struct {
    Username string   `json:"Username"`
    Salt     uint64   `json:"Salt,float64"`
    G        uint8    `json:"G,float64"`
    K        string   `json:"K"`
    V        *big.Int `json:"V,float64"`
    N        *big.Int `json:"N,float64"`
}

type ephemeral struct{
	A *big.Int `json:"A`
	Value_a uint64 `json:"a"`
}

func ConvertMaptoString(Map map[string]interface{})(string, error){

	jsonBytes, err := json.Marshal(Map)
    if err != nil {
        return "", err
    }

    return string(jsonBytes), nil

}

func StructToMap(item interface{}) map[string]interface{} {
    result := map[string]interface{}{}
    x := reflect.ValueOf(item).Elem()

    for i := 0; i < x.NumField(); i++ {
        field := x.Type().Field(i)
        value := x.Field(i).Interface()
        result[field.Name] = value
    }

    return result
}



func Client_Execution(){
	router := gin.Default()

	router.GET("/get_Hash",func(c *gin.Context){
		value := c.Query("value")
		Hash := NG_values.H(value)
		c.JSON(http.StatusOK,gin.H{
			"message":"Hash Value: "+Hash,
		})	
	})

	router.POST("/signup",func(c *gin.Context){
		var Username string = c.Query("username")
		var Password string = c.Query("password")
		user := &client.ClientDetails{}

		status := Execution.SaltandNG_generation(user)
		if !status{
			c.JSON(http.StatusConflict,gin.H{
				"message":"Error: N or G Creation Failure",
			})
			return
		}
		user.GenerateUsernamePassword(Username, Password)
		Map, err := user.SendToServer()

		if err != nil{
			c.JSON(http.StatusConflict,gin.H{
				"message":"Error: "+err.Error(),
			})
			return
		}

		Map_str, Map_err := ConvertMaptoString(Map)
		if Map_err != nil{
			c.JSON(http.StatusConflict,gin.H{
				"message":"Error: "+Map_err.Error(),
			})
			return
		}else{
			c.JSON(http.StatusOK,gin.H{
				"message": Map_str,
			})
		}
	})


	router.POST("/upload", func(c *gin.Context) {
		var Map_client Map_Client
	
		if err := c.ShouldBindJSON(&Map_client); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON provided",
			})
			return
		}

		inInterface := StructToMap(&Map_client)

		write_status := server.UserSignUp(inInterface)
	
		if write_status {
			c.JSON(http.StatusOK, gin.H{
				"message": "Signup Success",
			})
		} else {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Signup Failure",
			})
		}
	})

	router.POST("/computeA", func(c* gin.Context){
		var Map_client Map_Client
	
		if err := c.ShouldBindJSON(&Map_client); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON provided",
			})
			return
		}

		InInterface := StructToMap(&Map_client)
		user := client.FromServer(InInterface)
		if user == nil{
			c.JSON(http.StatusConflict,gin.H{
				"message":"Error: User not found",
			})
			return
		}else{
			user_tempdetails := user.GenerateA()
			a := user_tempdetails.Get_a()
			ephemeral := ephemeral{
				A: user_tempdetails.A,
				Value_a: a,
			}
			Map_str, Map_err := json.Marshal(ephemeral)
			if Map_err != nil{
				c.JSON(http.StatusConflict,gin.H{
					"message":"Error: "+Map_err.Error(),
				})
				return
			}else{
				c.JSON(http.StatusOK,gin.H{
					"message":Map_str,
				})
			}
		}
	})

	router.POST("computeA/compute-K_client", func(c *gin.Context){

	})

	router.Run("localhost:2002")
}