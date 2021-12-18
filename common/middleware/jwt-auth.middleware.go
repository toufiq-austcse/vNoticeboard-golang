package middleware

import (
	"fmt"
	"gihub.com/toufiq-austcse/vNoticeboard/api/services"
	"gihub.com/toufiq-austcse/vNoticeboard/common/helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func JwtAuthMiddleware(jwtService services.JwtService, instituteService services.InstituteService) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process the request", "No token found", helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		headerArr := strings.Split(authHeader, " ")
		if len(headerArr) != 2 {
			response := helper.BuildErrorResponse("Failed to process the request", "Token must start with Bearer", helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(headerArr[1])
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			instituteId, err := strconv.ParseUint(fmt.Sprintf("%v", claims["institute_id"]), 10, 64)
			if err != nil {
				panic(err.Error())
			}
			institute := instituteService.FindInstituteById(instituteId)
			context.Set("institute", institute)
			context.Next()
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
