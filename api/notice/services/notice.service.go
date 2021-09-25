package services

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/notice/repositories"
)

type INoticeService interface {
	Create()
	Update()
	Delete()
	Find()
}
type NoticeService struct {
	repository repositories.INoticeRepository
}

func NewNoticeService(repository repositories.INoticeRepository) INoticeService {
	return NoticeService{
		repository: repository,
	}
}

func (service NoticeService) Create() {
	service.repository.CreateNotice()

}

func (service NoticeService) Update() {
	service.repository.UpdateNotice()
}

func (service NoticeService) Delete() {
	service.repository.DeleteNotice()
}

func (service NoticeService) Find() {
	service.repository.FindNotice()
}
