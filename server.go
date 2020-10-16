package main

import (
	"log"
	"net/http"

	"src/play.src/controller"
	"src/play.src/middleware"
	"src/play.src/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	log.Println("Server Starter on :8080 . . .")

	router := gin.New()
	router.Use(gin.Recovery(), middleware.BasicAuth())

	router.GET("/videos", func(c *gin.Context) {
		c.JSON(http.StatusOK, videoController.ListAll())
	})

	router.POST("/videos", func(c *gin.Context) {
		c.JSON(http.StatusOK, videoController.Save(c))
	})

	router.Run()
}
