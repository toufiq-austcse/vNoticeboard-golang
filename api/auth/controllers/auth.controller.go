package controllers

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/auth/dtos/req"
	"gihub.com/toufiq-austcse/vNoticeboard/api/auth/services"
	"gihub.com/toufiq-austcse/vNoticeboard/common/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthController struct {
	authService services.AuthService
	jwtService  services.JwtService
}

func New(authService services.AuthService, jwtService services.JwtService) AuthController {
	return AuthController{authService: authService, jwtService: jwtService}
}

func (controller AuthController) Register(ctx *gin.Context) {
	var body req.RegisterReqDto
	if err := ctx.ShouldBind(&body); err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if controller.authService.IsDuplicateEmail(body.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)

	} else {
		createdInstitute := controller.authService.CreateInstitute(body)
		token := controller.jwtService.GenerateToken(strconv.FormatUint(createdInstitute.ID, 10))
		createdInstitute.Token = token
		response := helper.BuildResponse(true, "OK", createdInstitute)
		ctx.JSON(http.StatusCreated, response)
	}

}

func (controller AuthController) Login(ctx *gin.Context) {
	var loginDto req.LoginDto
	if err := ctx.ShouldBind(&loginDto); err != nil {
		response := helper.BuildErrorResponse("Failed to process the request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	isVerified, institute := controller.authService.VerifyCredentials(loginDto.Email, loginDto.Password)
	if !isVerified {
		response := helper.BuildErrorResponse("Please check again your credentials", "Invalid Credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	token := controller.jwtService.GenerateToken(strconv.FormatUint(institute.ID, 10))
	institute.Token = token
	response := helper.BuildResponse(true, "OK", institute)
	ctx.JSON(http.StatusCreated, response)

}

func (controller AuthController) Me(context *gin.Context) {

	institute, isExist := context.Get("institute")
	if !isExist {
		response := helper.BuildErrorResponse("Failed to process the request", "no institute found", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	respons := helper.BuildResponse(true, "", institute)
	context.JSON(http.StatusOK, respons)
}
