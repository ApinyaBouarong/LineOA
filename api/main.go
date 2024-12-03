package main

import (
	"LineOA/pkg/hook"
	"LineOA/pkg/linebot"

	"github.com/gin-gonic/gin"
)

func main() {
	linebot.InitLineBot()

	router := gin.Default()
	router.POST("/webhook", hook.HandleLineWebhook)
	router.Run(":8080")

	// db, err := database.ConnectToDB()
	// if err != nil {
	// 	log.Fatalf("Could not connect to database: %v", err)
	// }
	// defer db.Close()

	// fmt.Println("Connected to database successfully!")
}
