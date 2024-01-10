package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"srp/NG_values"
	"srp/Execution"
	"srp/client"
	"srp/server"
	"encoding/json"
	"fmt"
)

func ConvertMaptoString(Map map[string]interface{})(string, error){

	jsonBytes, err := json.Marshal(Map)
    if err != nil {
        fmt.Println(err)
        return "", err
    }

    return string(jsonBytes), nil

}

func ConvertStringtoMap(jsonStr string) (map[string]interface{}, error) {

    var Map map[string]interface{}

    err := json.Unmarshal([]byte(jsonStr), &Map)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    return Map, nil

}

func main(){
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

	router.POST("/signup/upload", func(c *gin.Context){
		Map_str := c.Query("map")
		Map, Map_err := ConvertStringtoMap(Map_str)

		if Map_err != nil{
			c.JSON(http.StatusConflict,gin.H{
				"message":"Error: "+Map_err.Error(),
			})
			return
		}
		var write_status bool = server.UserSignUp(Map)

		if write_status{
			c.JSON(http.StatusOK,gin.H{
				"message":"Signup Success",
			})
		}else{
			c.JSON(http.StatusConflict,gin.H{
				"message":"Signup Failure",
			})
		}
	})

	router.Run("localhost:2002")
}