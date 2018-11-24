package main

import (
	"github.com/gin-gonic/gin"
	"hhpage/route"
)

func main(){
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/",route.Home)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}