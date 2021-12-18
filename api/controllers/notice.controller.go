package controllers

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/dto/req"
	"gihub.com/toufiq-austcse/vNoticeboard/api/dto/res"
	"gihub.com/toufiq-austcse/vNoticeboard/api/entities"
	"gihub.com/toufiq-austcse/vNoticeboard/api/services"
	"gihub.com/toufiq-austcse/vNoticeboard/common/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type NoticeController struct {
	service services.INoticeService
}

func NewNoticeController(service services.INoticeService) NoticeController {
	return NoticeController{
		service: service,
	}
}

func (controller NoticeController) Insert(ctx *gin.Context) {
	var dto req.CreateNoticeReqDto
	errDto := ctx.ShouldBind(&dto)
	if errDto != nil {
		res := helper.BuildErrorResponse("Failed to process the response", errDto.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
		return
	}
	institute, _ := ctx.Get("institute")

	dto.InstituteID = institute.(entities.Institute).ID
	result := controller.service.Insert(dto)
	ctx.JSON(http.StatusOK, result)

}

func (controller NoticeController) FindOne(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("notice_id"), 0, 0)
	if err != nil {
		result := helper.BuildErrorResponse("No Param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, result)
	}
	institute, _ := ctx.Get("institute")
	result := controller.service.FindOneById(id, institute.(entities.Institute).ID)
	if (result == res.NoticeResponse{}) {
		response := helper.BuildErrorResponse("Data not found", "No Data with given Id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	response := helper.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, response)
	return
}

func (controller NoticeController) Update(ctx *gin.Context) {
	var dto req.UpdateNoticeReqDto
	errDto := ctx.ShouldBind(&dto)
	if errDto != nil {
		res := helper.BuildErrorResponse("Failed to process the response", errDto.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
		return
	}
	id, err := strconv.ParseUint(ctx.Param("notice_id"), 0, 0)
	if err != nil {
		result := helper.BuildErrorResponse("No Param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, result)
	}
	institute, _ := ctx.Get("institute")

	updatedNotice := controller.service.Update(id, dto, institute.(entities.Institute))
	response := helper.BuildResponse(true, "OK", updatedNotice)
	ctx.JSON(http.StatusOK, response)
	return
}

func (controller NoticeController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("notice_id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("No Param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	institute, _ := ctx.Get("institute")
	controller.service.DeleteOne(id, institute.(entities.Institute).ID)
	response := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (controller NoticeController) All(ctx *gin.Context) {
	institute, _ := ctx.Get("institute")
	results := controller.service.AllByInstitute(institute.(entities.Institute).ID)
	ctx.JSON(http.StatusOK, results)
}
