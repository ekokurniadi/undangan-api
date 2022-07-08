package routes

import (
	"undangan_online_api/handler"
	"undangan_online_api/repository"
	"undangan_online_api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type loveStoryRoute struct {
	db    *gorm.DB
	group *gin.RouterGroup
}

func NewLoveStoryRoute(db *gorm.DB, group *gin.RouterGroup) *loveStoryRoute {
	return &loveStoryRoute{db, group}
}

func (r *loveStoryRoute) RunLoveStoryRoute() {
	loveStoryRepository := repository.NewLoveStoryRepository(r.db)
	loveStoryService := service.NewLoveStoryService(loveStoryRepository)
	loveStoryHandler := handler.NewLoveStoryHandler(loveStoryService)

	r.group.POST("/love-story", loveStoryHandler.CreateLoveStory)
	r.group.GET("/love-story/:id/undangan", loveStoryHandler.GetLoveStorys)
	r.group.GET("/love-story/:id", loveStoryHandler.GetLoveStory)
	r.group.DELETE("/love-story/:id", loveStoryHandler.DeleteLoveStory)
	r.group.PUT("/love-story/:id", loveStoryHandler.UpdateLoveStory)
}
