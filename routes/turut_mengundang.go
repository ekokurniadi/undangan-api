package routes

import (
	"undangan_online_api/handler"
	"undangan_online_api/repository"
	"undangan_online_api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type turutMengundangRoute struct {
	db    *gorm.DB
	group *gin.RouterGroup
}

func NewTurutMengundangRoute(db *gorm.DB, group *gin.RouterGroup) *turutMengundangRoute {
	return &turutMengundangRoute{db, group}
}

func (r *turutMengundangRoute) RunTurutMengundangRoute() {
	turutMengundangRepository := repository.NewTurutMengundangRepository(r.db)
	turutMengundangService := service.NewTurutMengundangService(turutMengundangRepository)
	turutMengundangHandler := handler.NewTurutMengundangHandler(turutMengundangService)

	r.group.POST("/turut-mengundang", turutMengundangHandler.CreateTurutMengundang)
	r.group.GET("/turut-mengundang/:id/undangan", turutMengundangHandler.GetTurutMengundangs)
	r.group.GET("/turut-mengundang/:id", turutMengundangHandler.GetTurutMengundang)
	r.group.DELETE("/turut-mengundang/:id", turutMengundangHandler.DeleteTurutMengundang)
	r.group.PUT("/turut-mengundang/:id", turutMengundangHandler.UpdateTurutMengundang)
}
