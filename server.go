package main

import (
	authController "gihub.com/toufiq-austcse/vNoticeboard/api/auth/controllers"
	authService "gihub.com/toufiq-austcse/vNoticeboard/api/auth/services"
	"gihub.com/toufiq-austcse/vNoticeboard/api/institute/repositories"
	instituteService "gihub.com/toufiq-austcse/vNoticeboard/api/institute/services"
	"gihub.com/toufiq-austcse/vNoticeboard/common/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

var (
	instituteRepository repositories.InstituteRepository  = repositories.New()
	myInstituteService  instituteService.InstituteService = instituteService.NewInstituteService(instituteRepository)
	//instituteController controllers.InstituteController   = controllers.NewInstituteController(myInstituteService)

	myAuthService    authService.AuthService       = authService.New(myInstituteService)
	myAuthController authController.AuthController = authController.New(myAuthService)
)

func setupApplicationRootRoutes(appRoutesConfig *gin.Engine) {
	routerGroup := appRoutesConfig.Group("api")
	{
		authRouterGroup := routerGroup.Group("auth")
		authRouterGroup.POST("register", myAuthController.Register)
	}

}
func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	loadEnv()
	database.Init()
	setupApplicationRootRoutes(r)
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
