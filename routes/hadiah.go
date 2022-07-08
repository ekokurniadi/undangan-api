package routes

import (
	"undangan_online_api/handler"
	"undangan_online_api/repository"
	"undangan_online_api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type hadiahRoute struct {
	db    *gorm.DB
	group *gin.RouterGroup
}

func NewHadiahRoute(db *gorm.DB, group *gin.RouterGroup) *hadiahRoute {
	return &hadiahRoute{db, group}
}

func (r *hadiahRoute) RunHadiahRoute() {
	hadiahRepository := repository.NewHadiahRepository(r.db)
	hadiahService := service.NewHadiahService(hadiahRepository)
	hadiahHandler := handler.NewHadiahHandler(hadiahService)

	r.group.POST("/hadiah", hadiahHandler.CreateHadiah)
	r.group.GET("/hadiah/:id/undangan", hadiahHandler.GetHadiahs)
	r.group.GET("/hadiah/:id", hadiahHandler.GetHadiah)
	r.group.DELETE("/hadiah/:id", hadiahHandler.DeleteHadiah)
	r.group.PUT("/hadiah/:id", hadiahHandler.UpdateHadiah)
}
