package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayhello(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"hello",
	})
}

func main() {
	r:=gin.Default();
	r.GET("/hello",sayhello)

	r.Run("9090")

}
