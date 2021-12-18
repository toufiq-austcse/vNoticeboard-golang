package services

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/dto/req"
	"gihub.com/toufiq-austcse/vNoticeboard/api/dto/res"
	"gihub.com/toufiq-austcse/vNoticeboard/api/entities"
	"gihub.com/toufiq-austcse/vNoticeboard/api/repositories"
	"github.com/mashingan/smapping"
	"log"
)

type INoticeService interface {
	Insert(dto req.CreateNoticeReqDto) res.NoticeResponse
	Update(noticeId uint64, dto req.UpdateNoticeReqDto, institute entities.Institute) res.NoticeResponse
	DeleteOne(noticeId uint64, instituteId uint64)
	FindOneById(noticeId uint64, instituteId uint64) res.NoticeResponse
	AllByInstitute(instituteId uint64) []res.NoticeResponse
}
type NoticeService struct {
	repository repositories.INoticeRepository
}

func NewNoticeService(repository repositories.INoticeRepository) INoticeService {
	return NoticeService{
		repository: repository,
	}
}

func (service NoticeService) Insert(dto req.CreateNoticeReqDto) res.NoticeResponse {
	notice := entities.Notice{}
	err := smapping.FillStruct(&notice, smapping.MapFields(&dto))
	if err != nil {
		log.Fatalf("Fauled Map %v", err)
	}
	createdNotice := service.repository.InsertNotice(notice)

	return res.NoticeResponse{
		Id:          createdNotice.ID,
		Title:       createdNotice.Title,
		Description: createdNotice.Description,
	}
}

func (service NoticeService) Update(noticeId uint64, dto req.UpdateNoticeReqDto, institute entities.Institute) res.NoticeResponse {
	notice := service.repository.FindNoticeByID(noticeId, institute.ID)
	if (notice == entities.Notice{}) {
		return res.NoticeResponse{}
	}
	updatedNotice := service.repository.UpdateNotice(entities.Notice{ID: noticeId, Title: dto.Title, Description: dto.Description, Institute: institute})
	return res.NoticeResponse{
		Id:          updatedNotice.ID,
		Title:       updatedNotice.Title,
		Description: updatedNotice.Description,
	}
}

func (service NoticeService) DeleteOne(noticeId uint64, instituteId uint64) {
	service.repository.DeleteNotice(noticeId, instituteId)
}

func (service NoticeService) FindOneById(noticeId uint64, instituteId uint64) res.NoticeResponse {
	notice := service.repository.FindNoticeByID(noticeId, instituteId)
	return res.NoticeResponse{
		Id:          notice.ID,
		Title:       notice.Title,
		Description: notice.Description,
	}
}

func (service NoticeService) AllByInstitute(instituteId uint64) []res.NoticeResponse {
	results := service.repository.FindAllNoticesByInstitute(instituteId)

	var notices []res.NoticeResponse
	for _, notice := range results {
		notices = append(notices, res.NoticeResponse{
			Id:          notice.ID,
			Title:       notice.Title,
			Description: notice.Description,
		})
	}
	return notices
}
