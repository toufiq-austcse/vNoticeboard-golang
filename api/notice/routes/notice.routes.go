package routes

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/institute/controllers"
	"gihub.com/toufiq-austcse/vNoticeboard/api/notice/repositories"
	"gihub.com/toufiq-austcse/vNoticeboard/api/notice/services"
	"github.com/gin-gonic/gin"
)

func Init(routerGroup *gin.RouterGroup) {
	repo := repositories.NewNoticeRepository()
	service := services.NewNoticeService(repo)
	controller := controllers.NewInstituteController(service)

	noticeRouterGroup := routerGroup.Group("notices")

	noticeRouterGroup.GET("", controller.Find)
	noticeRouterGroup.POST("", controller.Create)
	noticeRouterGroup.PATCH("", controller.Update)
	noticeRouterGroup.DELETE("", controller.Delete)

}
