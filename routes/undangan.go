package routes

import (
	"undangan_online_api/handler"
	"undangan_online_api/repository"
	"undangan_online_api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type undanganRoute struct {
	db    *gorm.DB
	group *gin.RouterGroup
}

func NewUndanganRoute(db *gorm.DB, group *gin.RouterGroup) *undanganRoute {
	return &undanganRoute{db, group}
}

func (r *undanganRoute) RunUndanganRoute() {
	undanganRepository := repository.NewUndanganRepository(r.db)
	undanganService := service.NewUndanganService(undanganRepository)
	undanganHandler := handler.NewUndanganHandler(undanganService)

	r.group.POST("/undangan", undanganHandler.CreateUndangan)
	r.group.GET("/undangan", undanganHandler.GetUndangans)
	r.group.GET("/undangan/:id", undanganHandler.GetUndangan)
	r.group.DELETE("/undangan/:id", undanganHandler.DeleteUndangan)
	r.group.PUT("/undangan/:id", undanganHandler.UpdateUndangan)

}
