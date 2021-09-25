package services

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/institute/entities"
	"gihub.com/toufiq-austcse/vNoticeboard/api/institute/repositories"
)


type InstituteService struct {
	repository repositories.InstituteRepository
}

func NewInstituteService(repository repositories.InstituteRepository) InstituteService {
	return InstituteService{
		repository: repository,
	}
}




func (service InstituteService) FindInstituteByEmail(email string) (entities.Institute,error) {
	return service.repository.FindByEmail(email)
}
