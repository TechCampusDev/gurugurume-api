package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func hello(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func dbtest(c *gin.Context){
	db := dbConnect()
	fmt.Println(db)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func getting(c *gin.Context){
	db := dbConnect()
	var user User
	db.First(&user, 1)
	c.JSON(200, gin.H{
		"message": user,
	})
}

func post(c *gin.Context){
	db := dbConnect()
	db.Create(&User{Name: "huku", Email: "sample"})
	c.JSON(200, gin.H{
		"message": "success",
	})
}