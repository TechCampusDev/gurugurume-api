package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//sample
	router.GET("/", hello)
	router.GET("/getNearBy", getNearBy)
	router.GET("/db", dbtest)
	router.GET("/get", getting)
	router.GET("/post", post)

	//init_app
	router.GET("/setenv", setEnv)

	//application
	router.GET("/setup", setup)
	router.GET("/checktoken", check_token)

	router.Run(":8080")
}
