package main

import (
	"github.com/gin-gonic/gin"
	"github.com/loner-way/loner-gin-vue/controller"
)

func main()  {
	 r := gin.Default()
	 r.GET("/Ping", controller.Ping)

	 // 用户注册
	r.GET("/api/user/register",controller.Register)

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
