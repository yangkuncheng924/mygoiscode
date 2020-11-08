package main

import (
	"net/http"

	"github.com/gin-gonic/gin")

func sayHello(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"hello golang!",
	})
}

func main() {
    r := gin.Default()

	r.GET("/hello",sayHello)
	
	r.POST("hello",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"method":"POST",
		})
	})

	r.PUT("hello",func(c *gin.Context){
		c.JSON(200,gin.H{
			"messger":"delect",
		})
	})
	r.Run(":9090")	
}
/*
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("D:/mycode/go/proijet/src/hello.txt")

	_, _ = fmt.Fprintln(w, string(b))

}
func main() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("error http list:%v\n", err)
		return
	}
}
*/