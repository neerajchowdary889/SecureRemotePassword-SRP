package main

import(
	// "fmt"
	// "srp/Execution"
	"net/http"
	"github.com/gin-gonic/gin"
)

// func main(){
// 	Execution.Execution()
// }

func main(){
	router := gin.Default()

    router.GET("/test", func(c *gin.Context) {
        name := c.Query("name")
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello " + name,
        })
    })

	router.POST("/post_test",func(c *gin.Context){
		name := c.Query("name")
		c.JSON(http.StatusOK,gin.H{
			"message":"Hello "+name,
		})
	})

	router.Run("localhost:2002")
}