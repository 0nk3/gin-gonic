package service

import (
	"src/play.src/entity"
)

// VideoService . .
type VideoService interface {
	Save(entity.Video) entity.Video
	ListAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

// New . .
func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video

}
func (service *videoService) ListAll() []entity.Video {
	return service.videos
}
