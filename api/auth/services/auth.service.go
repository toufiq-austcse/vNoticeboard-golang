package services

import (
	"fmt"
	"gihub.com/toufiq-austcse/vNoticeboard/api/auth/dtos/req"
	"gihub.com/toufiq-austcse/vNoticeboard/api/institute/entities"
	"gihub.com/toufiq-austcse/vNoticeboard/api/institute/services"
)

type AuthService struct {
	InstituteService services.InstituteService
}

func New(instituteService services.InstituteService) AuthService {
	return AuthService{
		InstituteService: instituteService,
	}
}
func (service AuthService) LoginUser() string {
	return ""
}

func (service AuthService) RegisterUser(data req.RegisterReqDto) string {
	fmt.Println(data)
	return ""
}

func (service AuthService) IsDuplicateEmail(email string) bool {
	institute, _ := service.InstituteService.FindInstituteByEmail(email)
	return institute != entities.Institute{}
}
