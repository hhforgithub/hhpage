package main

import (
	"github.com/gin-gonic/gin"
	"hhpage/route"
)

func main(){
	engine := gin.Default()
	engine.LoadHTMLGlob("web/views/*")
	engine.Static("/resource","web/resource")

	setupRouter(engine)

	engine.Run(":8080")
}

func setupRouter(engine *gin.Engine){
	engine.GET("/",route.Home)
}