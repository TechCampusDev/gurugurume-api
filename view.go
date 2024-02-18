package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"os"
)

// 匿名アカウントを作成し、そのトークンを返却する
func setup(c *gin.Context){
	godotenv.Load(".env")

	db := dbConnect()
	user := User{Name: "unknown", Email: "sample"}
	db.Create(&user)

	claims := jwt.MapClaims{
		"user_id": user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	c.JSON(200, gin.H{
		"token": tokenString,
	})
}

// トークンを解析し、user_idを調べる
func check_token(c *gin.Context){
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.Q66-3qLvTBhyI-14h1NDqkULwdcd3R7iKvAcaOL6VoU"
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	c.JSON(200, gin.H{
		"user_id": int64(claims["user_id"].(float64)),
	})
}