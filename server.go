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
		err := videoController.Save(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"rror": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Video input is valid"})
		}
	})
	router.Run()
}
