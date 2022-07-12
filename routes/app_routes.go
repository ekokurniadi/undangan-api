package routes

import (
	"undangan_online_api/auth"
	"undangan_online_api/middleware"
	"undangan_online_api/repository"
	"undangan_online_api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type routesInjector struct {
	db *gorm.DB
}

func NewRoutesInjector(db *gorm.DB) *routesInjector {
	return &routesInjector{db}
}

func (inject *routesInjector) RunAppRouters() {

	authService := auth.NewService()
	userRepository := repository.NewUserRepository(inject.db)
	userService := service.NewUserService(userRepository)

	middleWares := middleware.NewMiddleWareService(authService, userService)

	router := gin.Default()
	router.Use(middleWares.CORSMiddleware())
	api := router.Group("/api/v1")

	NewUserRoute(inject.db, api, middleWares, authService).RunUserRoute()
	NewGaleriRoute(inject.db, api).RunGaleriRoute()
	NewUndanganRoute(inject.db, api).RunUndanganRoute()
	NewGuestBookRoute(inject.db, api).RunGuestBookRoute()
	NewHadiahRoute(inject.db, api).RunHadiahRoute()
	NewLoveStoryRoute(inject.db, api).RunLoveStoryRoute()
	NewUndanganDetailRoute(inject.db, api).RunUndanganDetailRoute()
	NewTurutMengundangRoute(inject.db, api).RunTurutMengundangRoute()

	router.Run()
}
