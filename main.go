package main

import (
	"github.com/gin-gonic/gin"
	"hhpage/route"
)

func main(){
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/",route.Home)
	r.Run(":80")
}