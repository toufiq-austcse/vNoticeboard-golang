package main

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/controllers"
	"gihub.com/toufiq-austcse/vNoticeboard/api/repositories"
	"gihub.com/toufiq-austcse/vNoticeboard/api/services"
	"gihub.com/toufiq-austcse/vNoticeboard/common/middleware"

	"gihub.com/toufiq-austcse/vNoticeboard/common/database"
	"github.com/gin-gonic/gin"
)

var (
	db                  = database.SetupDBConnection()
	instituteRepository = repositories.New(db)
	myInstituteService  = services.NewInstituteService(instituteRepository)
	//instituteController controllers.InstituteController   = controllers.NewInstituteController(myInstituteService)

	noticeRepository = repositories.NewNoticeRepository(db)
	myNoticeService  = services.NewNoticeService(noticeRepository)

	myAuthService = services.NewAuthService(myInstituteService)
	myJwtService  = services.NewJWtService()

	myAuthController   = controllers.NewAuthController(myAuthService, myJwtService)
	myNoticeController = controllers.NewNoticeController(myNoticeService)
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
	noticeRoutes := r.Group("api/notices", middleware.JwtAuthMiddleware(myJwtService, myInstituteService))
	{
		noticeRoutes.POST("/", myNoticeController.Insert)
		noticeRoutes.GET("/", myNoticeController.All)
		noticeRoutes.GET("/:notice_id", myNoticeController.FindOne)
		noticeRoutes.DELETE("/:notice_id", myNoticeController.Delete)
		noticeRoutes.PATCH("/:notice_id", myNoticeController.Update)
	}
	r.Run()
}
