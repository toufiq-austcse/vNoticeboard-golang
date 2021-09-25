package controllers

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/notice/services"
	"github.com/gin-gonic/gin"
)

type INoticeController interface {
	Create(ctx *gin.Context)
	Find(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
type  NoticeController struct {
	service services.INoticeService
}

func NewNoticeController(service services.INoticeService) INoticeController {
	return NoticeController{
		service: service,
	}
}

func (controller NoticeController) Create(ctx *gin.Context){
	controller.service.Create()
	ctx.JSON(201,gin.H{"data":"ok"})
}

func (controller NoticeController) Find(ctx *gin.Context){
	controller.service.Find()
	ctx.JSON(201,gin.H{"data":"ok"})
}

func (controller NoticeController) Update(ctx *gin.Context){
	controller.service.Update()
	ctx.JSON(201,gin.H{"data":"ok"})
}

func (controller NoticeController) Delete(ctx *gin.Context){
	controller.service.Delete()
	ctx.JSON(201,gin.H{"data":"ok"})
}