package controller

import (
	"src/play.src/entity"
	"src/play.src/service"
	"github.com/gin-gonic/gin"
)

// VideoController . .
type VideoController interface {
	ListAll() []entity.Video
	Save(c *gin.Context) error
}

type controller struct {
	service service.VideoService
}

// New . .
func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) ListAll() []entity.Video {
	return c.service.ListAll()
}
func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}
