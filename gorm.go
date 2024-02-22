package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "github.com/joho/godotenv"
    "fmt"
  )

func dbConnect() *gorm.DB {
    godotenv.Load(".env")

    HOST := os.Getenv("DB_HOST")
    USER := os.Getenv("DB_USER")
    PASS := os.Getenv("DB_PASS")
    DBNAME := os.Getenv("DBNAME")
    dsn := "host=" + HOST + " user=" + USER + " password=" + PASS + " dbname=" + DBNAME + " port=5432" + " sslmode=disable TimeZone=Asia/Shanghai"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic(err.Error())
    }
    fmt.Println("db connected: ", &db)
    return db
}


/* func main() {
 db := gormConnect()

 defer db.Close()
 db.LogMode(true)
} */
