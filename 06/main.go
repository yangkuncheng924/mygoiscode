package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	r :=gin.Default()

	r.GET("Web",func(c *gin.Context){

		name := c.Query("query")
		c.JSON(200,gin.H{
			"name":name,
		})
	})
	r.Run(":9090")

}