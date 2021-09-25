package controllers

import (
	"fmt"
	"gihub.com/toufiq-austcse/vNoticeboard/api/auth/dtos/req"
	"gihub.com/toufiq-austcse/vNoticeboard/api/auth/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService services.AuthService
}

func New(authService services.AuthService) AuthController {
	return AuthController{authService: authService}
}

func (controller AuthController) Register(ctx *gin.Context) {
	var body req.RegisterReqDto
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"data":  nil,
		})
		return
	}
	if controller.authService.IsDuplicateEmail(body.Email) {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "Duplicate Email",
			"data":  nil,
		})
		return
	} else {
		fmt.Println("called")
	}

}

func (controller AuthController) Login(ctx *gin.Context) {
	//controller.service.LoginUser()
}
