package main

import (
	"github.com/gin-gonic/gin"
	"loner-gin-vue/controller"
)

func main()  {
	 r := gin.Default()
	 r.GET("/Ping", controller.Ping)
	// listen and serve on 0.0.0.0:8080
	r.Run()
}
