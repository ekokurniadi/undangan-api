package routes

import (
	"undangan_online_api/handler"
	"undangan_online_api/repository"
	"undangan_online_api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type guestBookRoute struct {
	db    *gorm.DB
	group *gin.RouterGroup
}

func NewGuestBookRoute(db *gorm.DB, group *gin.RouterGroup) *guestBookRoute {
	return &guestBookRoute{db, group}
}

func (r *guestBookRoute) RunGuestBookRoute() {
	guestBookRepository := repository.NewGuestBookRepository(r.db)
	guestBookService := service.NewGuestBookService(guestBookRepository)
	guestBookHandler := handler.NewGuestBookHandler(guestBookService)

	r.group.POST("/guest-book", guestBookHandler.CreateGuestBook)
	r.group.GET("/guest-book/:id/undangan", guestBookHandler.GetGuestBooks)
	r.group.GET("/guest-book/:id", guestBookHandler.GetGuestBook)
	r.group.DELETE("guest-book/:id", guestBookHandler.DeleteGuestBook)
	r.group.PUT("/guest-book/:id", guestBookHandler.UpdateGuestBook)

}
