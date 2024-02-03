package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	router := gin.Default()

	router.GET("/", hello)
	router.GET("/db", dbtest)
	router.GET("/get", getting)
	router.GET("/post", post)

	router.Run(":8080")
}

func create_db(){
	db := dbConnect()
	db.AutoMigrate(&User{})
	fmt.Println("finished")
}
