package main

import (
	"github/yang/controller"
	"github/yang/tool"

	"github.com/gin-gonic/gin"
)


func main() {
	cfg,err := tool.ParseConfig("./config/app.json")
	if err!=nil{
		panic(err.Error())
	}
	r := gin.Default()
	registerRouter(r)
	r.Run(cfg.AppHost+":"+cfg.AppPort)
}

func registerRouter(router *gin.Engine){
	new(controller.HelloController).Router(router)
}