package routes

import (
	"undangan_online_api/handler"
	"undangan_online_api/repository"
	"undangan_online_api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type galeriRoute struct {
	db    *gorm.DB
	group *gin.RouterGroup
}

func NewGaleriRoute(db *gorm.DB, group *gin.RouterGroup) *galeriRoute {
	return &galeriRoute{db, group}
}

func (r *galeriRoute) RunGaleriRoute() {
	galeriRepository := repository.NewGaleriRepository(r.db)
	galeriService := service.NewGaleriService(galeriRepository)
	galeriHandler := handler.NewGaleriHandler(galeriService)

	r.group.POST("/galeri", galeriHandler.CreateGaleri)
	r.group.GET("/galeri", galeriHandler.GetGaleris)
	r.group.GET("/galeri/:id", galeriHandler.GetGaleri)
	r.group.DELETE("/galeri/:id", galeriHandler.DeleteGaleri)
	r.group.GET("/galeri/:id/undangan", galeriHandler.GetGaleris)
}
