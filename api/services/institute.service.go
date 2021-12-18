package services

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/entities"
	"gihub.com/toufiq-austcse/vNoticeboard/api/repositories"
)

type InstituteService struct {
	repository repositories.InstituteRepository
}

func NewInstituteService(repository repositories.InstituteRepository) InstituteService {
	return InstituteService{
		repository: repository,
	}
}

func (service InstituteService) FindInstituteByEmail(email string) entities.Institute {
	return service.repository.FindByEmail(email)
}

func (service InstituteService) InsertInstitute(dto entities.Institute) entities.Institute {
	return service.repository.InsertUser(dto)

}

func (service InstituteService) FindInstituteById(id uint64) entities.Institute {
	return service.repository.FindInstituteById(id)
}
