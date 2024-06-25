package main

import (
	"net/http"

	docs "front_api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /

// PingExample godoc
// @Summary Dontpanic-Daily
// @Schemes
// @Description 42Labs server entrypoint
// @Tags Game
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router / [get]
func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	r.Static("/public", "./public")
	r.StaticFile("/styles.css", "./public/styles.css")
	r.StaticFile("/star.png", "./public/star.png")
	r.StaticFile("/main.js", "./public/main.js")
	r.LoadHTMLGlob("public/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
