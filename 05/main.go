package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.GET("json",func(c *gin.Context) {
		// data := map[string]interface{}{
		// 	"名字":"小杨",
		// 	"年龄":18,
		// 	"爱好":"吃吃",
		// }
		data := gin.H{"name":"小杨","age":18,"angle":"吃",}
		c.JSON(200,data)
	})

	
	type msg struct{
		Name string  `json:"name"`
		Helg string
		Age  int
		
	}
	r.GET("json1",func(c *gin.Context) {

		data := msg{
			"小杨",
			"小李",
			26,
			
		}
		c.JSON(200,data)
	})

	r.Run(":9090")
	fmt.Println("hello world")
}