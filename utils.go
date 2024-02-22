package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func setEnv(c *gin.Context){
	create_table()
	c.JSON(200, gin.H{
		"message": "success setting environment",
	})
}

func create_table(){
	db := dbConnect()
	db.AutoMigrate(&User{})
	fmt.Println("create_table_finished")
}