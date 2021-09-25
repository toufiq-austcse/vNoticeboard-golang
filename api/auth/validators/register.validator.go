package validators

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/auth/dtos/req"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func RegisterValidator(c *gin.Context) {
	var registerDto req.RegisterReqDto
	if err := c.ShouldBindJSON(&registerDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return

	}

	validate := validator.New()
	if err := validate.Struct(registerDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Next()

}
