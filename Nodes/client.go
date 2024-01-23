package Nodes

import (
	"encoding/json"
	"fmt"
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
	Salt     uint64   `json:"Salt"`
	G        uint8    `json:"G"`
	K        string   `json:"K"`
	V        *big.Int `json:"V"`
	N        *big.Int `json:"N"`
}

type ClientTempDetails struct {
	A        *big.Int `json:"A"`
	B        *big.Int `json:"B"`
	a        uint64   `json:"a"`
	u        string   `json:"u"`
	K_client string   `json:"K_client"`
}

type U_generation struct{
	A *big.Int `json:"A"`
	B *big.Int `json:"B"`
}
type ephemeral struct {
	A       *big.Int `json:"A"`
	Value_a uint64   `json:"a"`
}
type priv_vars struct{
	Value_a uint64 `json:"a"`
	Value_u string `json:"u"`
}

func ConvertMaptoString(Map map[string]interface{}) (string, error) {

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

func Client_Execution() {
	router := gin.Default()

	router.GET("/get_Hash", func(c *gin.Context) {
		value := c.Query("value")
		Hash := NG_values.H(value)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hash Value: " + Hash,
		})
	})

	router.POST("/signup", func(c *gin.Context) {
		var Username string = c.Query("username")
		var Password string = c.Query("password")
		user := &client.ClientDetails{}

		status := Execution.SaltandNG_generation(user)
		if !status {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Error: N or G Creation Failure",
			})
			return
		}
		user.GenerateUsernamePassword(Username, Password)
		Map, err := user.SendToServer()

		if err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Error: " + err.Error(),
			})
			return
		}

		Map_str, Map_err := ConvertMaptoString(Map)
		if Map_err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Error: " + Map_err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
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

		fmt.Printf("%+v\n", Map_client)

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

	router.POST("/computeA", func(c *gin.Context) {
		var Map_client Map_Client

		if err := c.ShouldBindJSON(&Map_client); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON provided",
			})
			return
		}
		fmt.Printf("%+v\n", Map_client)
		InInterface := StructToMap(&Map_client)
		user := client.FromServer(InInterface)
		if user == nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Error: User not found",
			})
			return
		} else {
			user_tempdetails := user.GenerateA()
			a := user_tempdetails.Get_a()
			ephemeral := ephemeral{
				A:       user_tempdetails.A,
				Value_a: a,
			}
			Map_str, Map_err := json.Marshal(ephemeral)
			if Map_err != nil {
				c.JSON(http.StatusConflict, gin.H{
					"message": "Error: " + Map_err.Error(),
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": Map_str,
				})
			}
		}
	})

	router.POST("/computeU", func(c *gin.Context) {

		var user_tempdetails ClientTempDetails

		if err := c.ShouldBindJSON(&user_tempdetails); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON provided",
			})
			return
		}
		fmt.Printf("A-->%+v\n", user_tempdetails.A)
		fmt.Printf("B-->%+v\n", user_tempdetails.B)
		fmt.Printf("a-->%+v\n", user_tempdetails.a)
		u_client := server.Server_ComputeU_test(user_tempdetails.A, user_tempdetails.B)

		Map_str, Map_err := json.Marshal(u_client)
		if Map_err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Error: " + Map_err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": Map_str,
			})
		}
	})

	type Combined struct {
		UserTempDetails client.ClientTempDetails `json:"user_tempdetails"`
		User            client.ClientDetails     `json:"user"`
		Priv_vars		priv_vars				 `json:"priv_vars"`
		U_generation	U_generation			 `json:"U_generation"`
	}

	router.POST("computeA/compute-K_client", func(c *gin.Context) {

		var combined Combined
		if err := c.ShouldBindJSON(&combined); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON provided",
			})
			return
		}
		var Password string = c.Query("password")
		combined.UserTempDetails.Set_a(combined.Priv_vars.Value_a)
		combined.UserTempDetails.Set_u(combined.Priv_vars.Value_u)
		combined.UserTempDetails.A = combined.U_generation.A
		// fmt.Printf("A-->%+v\n", combined.UserTempDetails)
		// fmt.Printf("B-->%+v\n", combined.User)

		status := combined.User.Compute_K_client(&combined.UserTempDetails, Password)
		if !status {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Error: K_client Creation Failure",
			})
			return
		}

		Map_str, Map_err := json.Marshal(combined.UserTempDetails.K_client)
		if Map_err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Error: " + Map_err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": Map_str,
			})
		}
	})

	router.POST("computeA/compute-K_client/computeM1", func(c *gin.Context) {
		var combined Combined
		if err := c.ShouldBindJSON(&combined); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON provided",
			})
			return
		}

		combined.UserTempDetails.Set_a(combined.Priv_vars.Value_a)
		combined.UserTempDetails.Set_u(combined.Priv_vars.Value_u)
		combined.UserTempDetails.A = combined.U_generation.A
		// fmt.Printf("A-->%+v\n", combined.UserTempDetails)
		// fmt.Printf("B-->%+v\n", combined.User)

		M1 := combined.User.GenerateM1(&combined.UserTempDetails)
		Map_str, Map_err := json.Marshal(M1)
		if Map_err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Error: " + Map_err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": Map_str,
			})
		}
	})

	router.POST("computeA/compute-K_client/computeM", func(c *gin.Context) {
		var combined Combined
		if err := c.ShouldBindJSON(&combined); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON provided",
			})
			return
		}
		combined.UserTempDetails.Set_a(combined.Priv_vars.Value_a)
		combined.UserTempDetails.Set_u(combined.Priv_vars.Value_u)
		combined.UserTempDetails.A = combined.U_generation.A
		// fmt.Printf("A-->%+v\n", combined.UserTempDetails)
		// fmt.Printf("B-->%+v\n", combined.User)

		M_1 := c.Query("M_1")
		M := combined.User.GenerateM(&combined.UserTempDetails, M_1)
		Map_str, Map_err := json.Marshal(M)
		if Map_err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Error: " + Map_err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": Map_str,
			})
		}
	})

	router.GET("/get_Hash_2", func(c *gin.Context) {
		value := c.Query("value")
		Hash := NG_values.H(value)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hash Value: " + Hash,
		})
	})

	router.Run("localhost:2004")

}
