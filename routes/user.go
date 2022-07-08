package routes

import (
	"undangan_online_api/auth"
	"undangan_online_api/handler"
	"undangan_online_api/middleware"
	"undangan_online_api/repository"
	"undangan_online_api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userRoute struct {
	db          *gorm.DB
	group       *gin.RouterGroup
	mw          *middleware.MiddleWareService
	authService auth.Service
}

func NewUserRoute(db *gorm.DB, group *gin.RouterGroup, mw *middleware.MiddleWareService, authService auth.Service) *userRoute {
	return &userRoute{db, group, mw, authService}
}

func (r *userRoute) RunUserRoute() {
	userRepository := repository.NewUserRepository(r.db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, r.authService)

	r.group.POST("/auth/login", userHandler.Login)
	r.group.GET("/auth/users", r.mw.AuthMiddleware(), userHandler.FetchUser)

	r.group.POST("/users", r.mw.AuthMiddleware(), userHandler.CreateUser)
	r.group.GET("/users/:id", r.mw.AuthMiddleware(), userHandler.GetUser)
	r.group.GET("/users/", r.mw.AuthMiddleware(), userHandler.GetUsers)
	r.group.PUT("/users/:id", r.mw.AuthMiddleware(), userHandler.UpdateUser)
	r.group.DELETE("/users/:id", r.mw.AuthMiddleware(), userHandler.DeleteUser)

}
