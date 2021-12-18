package services

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/dto/req"
	"gihub.com/toufiq-austcse/vNoticeboard/api/entities"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService struct {
	InstituteService InstituteService
}

func NewAuthService(instituteService InstituteService) AuthService {
	return AuthService{
		InstituteService: instituteService,
	}
}
func (service AuthService) LoginUser() string {
	return ""
}

func (service AuthService) CreateInstitute(dto req.RegisterReqDto) entities.Institute {

	instituteToCreate := entities.Institute{}
	err := smapping.FillStruct(&instituteToCreate, smapping.MapFields(&dto))
	if err != nil {
		log.Println("called")
		log.Fatalf("Failed to map %v", err)
	}
	res := service.InstituteService.InsertInstitute(instituteToCreate)
	return res
}

func (service AuthService) IsDuplicateEmail(email string) bool {
	institute := service.InstituteService.FindInstituteByEmail(email)
	return institute != entities.Institute{}
}

func (service AuthService) VerifyCredentials(email string, password string) (bool, entities.Institute) {
	institute := service.InstituteService.FindInstituteByEmail(email)

	if institute.Email == "" {
		return false, entities.Institute{}
	}
	if comparePassword(institute.Password, []byte(password)) {
		return true, institute
	}
	return false, entities.Institute{}

}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
