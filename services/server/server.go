package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"localhost.com/game_application"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/api", game_application.Game)
	r.Run(":8081")
}
