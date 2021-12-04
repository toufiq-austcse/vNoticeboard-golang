package main

import (
	authController "gihub.com/toufiq-austcse/vNoticeboard/api/auth/controllers"
	authService "gihub.com/toufiq-austcse/vNoticeboard/api/auth/services"
	"gihub.com/toufiq-austcse/vNoticeboard/common/middleware"

	"gihub.com/toufiq-austcse/vNoticeboard/api/institute/repositories"
	instituteService "gihub.com/toufiq-austcse/vNoticeboard/api/institute/services"
	"gihub.com/toufiq-austcse/vNoticeboard/common/database"
	"github.com/gin-gonic/gin"
)

var (
	db                  = database.SetupDBConnection()
	instituteRepository = repositories.New(db)
	myInstituteService  = instituteService.NewInstituteService(instituteRepository)
	//instituteController controllers.InstituteController   = controllers.NewInstituteController(myInstituteService)

	myAuthService = authService.NewAuthService(myInstituteService)
	myJwtService  = authService.NewJWtService()

	myAuthController = authController.New(myAuthService, myJwtService)
)

func main() {

	defer database.CloseDBConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/register", myAuthController.Register)
		authRoutes.POST("/login", myAuthController.Login)
		authRoutes.GET("/me", middleware.JwtAuthMiddleware(myJwtService, myInstituteService), myAuthController.Me)

	}
	r.Run()
}
