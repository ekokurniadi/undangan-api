package routes

import (
	"undangan_online_api/handler"
	"undangan_online_api/repository"
	"undangan_online_api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type undanganDetailRoute struct {
	db    *gorm.DB
	group *gin.RouterGroup
}

func NewUndanganDetailRoute(db *gorm.DB, group *gin.RouterGroup) *undanganDetailRoute {
	return &undanganDetailRoute{db, group}
}

func (r *undanganDetailRoute) RunUndanganDetailRoute() {
	undanganDetailRepository := repository.NewUndanganDetailRepository(r.db)
	undanganDetailService := service.NewUndanganDetailService(undanganDetailRepository)
	undanganDetailHandler := handler.NewUndanganDetailHandler(undanganDetailService)

	r.group.POST("/undangan-detail", undanganDetailHandler.CreateUndanganDetail)
	r.group.GET("/undangan-detail/:id/undangan", undanganDetailHandler.GetUndanganDetailByKode)
	r.group.GET("/undangan-detail/:id", undanganDetailHandler.GetUndanganDetail)
	r.group.DELETE("/undangan-detail/:id", undanganDetailHandler.DeleteUndanganDetail)
	r.group.PUT("/undangan-detail/:id", undanganDetailHandler.UpdateUndanganDetail)
}
