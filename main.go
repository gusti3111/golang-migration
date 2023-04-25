package main

import (
	// "log"

	// "github.com/golang-migrate/migrate/v4"
	// _ "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

func main() {
	r := gin.Default()

	// conn := "postgres://postgres:12345@localhost:5438/golangdb?sslmode=disable"
	conn := "postgresql://postgres:Fou8Fa145Pf0KNTemhHW@containers-us-west-59.railway.app:5575/railway"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	type accounts struct {
		Username string
		Password string
		Email    string
	}
	db.AutoMigrate(&accounts{})

	r.POST("/account", func(ctx *gin.Context) {
		var newaccount accounts
		if err := ctx.ShouldBindJSON(&newaccount); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&newaccount)
		ctx.JSON(http.StatusCreated, newaccount)

	})

	r.GET("/accounts", func(ctx *gin.Context) {
		var account []accounts
		db.Find(&account)
		ctx.JSON(200, gin.H{
			"account": account,
		})
	})
	r.Run(":8000")

}
