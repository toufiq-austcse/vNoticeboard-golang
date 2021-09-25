package routes

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/institute/controllers"
	"gihub.com/toufiq-austcse/vNoticeboard/api/institute/repositories"
	"gihub.com/toufiq-austcse/vNoticeboard/api/institute/services"
	"github.com/gin-gonic/gin"
)

func Init(routerGroup *gin.RouterGroup) {
	repo := repositories.NewInstituteRepository()
	service := services.NewInstituteService(repo)
	controller := controllers.NewInstituteController(service)

	instituteRouterGroup := routerGroup.Group("institutes")

	instituteRouterGroup.GET("", controller.Find)
	instituteRouterGroup.POST("", controller.Create)
	instituteRouterGroup.PATCH("", controller.Update)
	instituteRouterGroup.DELETE("", controller.Delete)

}
