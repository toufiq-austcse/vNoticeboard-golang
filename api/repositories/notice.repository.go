package repositories

import (
	"fmt"
	"gihub.com/toufiq-austcse/vNoticeboard/api/entities"
	"gorm.io/gorm"
)

type INoticeRepository interface {
	InsertNotice(notice entities.Notice) entities.Notice
	UpdateNotice(updateNotice entities.Notice) entities.Notice
	FindNoticeByID(noticeId uint64, instituteId uint64) entities.Notice
	DeleteNotice(noticeId uint64, instituteId uint64)
	FindAllNoticesByInstitute(instituteId uint64) []entities.Notice
}
type NoticeRepository struct {
	connection *gorm.DB
}

func NewNoticeRepository(dbConn *gorm.DB) INoticeRepository {
	return NoticeRepository{
		connection: dbConn,
	}
}

func (repo NoticeRepository) UpdateNotice(updateNotice entities.Notice) entities.Notice {
	fmt.Println("updateNotice", updateNotice)
	repo.connection.Save(&updateNotice)
	return updateNotice
}

func (repo NoticeRepository) FindNoticeByID(noticeId uint64, instituteId uint64) entities.Notice {
	var notice entities.Notice
	repo.connection.Where(&entities.Notice{ID: noticeId, InstituteID: instituteId}).Find(&notice)
	fmt.Println(notice)
	return notice
}

func (repo NoticeRepository) DeleteNotice(noticeId uint64, instituteId uint64) {
	repo.connection.Delete(&entities.Notice{InstituteID: instituteId, ID: noticeId})
}

func (repo NoticeRepository) InsertNotice(notice entities.Notice) entities.Notice {
	repo.connection.Save(&notice)
	repo.connection.Preload("Institute").Find(&notice)
	return notice
}
func (repo NoticeRepository) FindAllNoticesByInstitute(instituteId uint64) []entities.Notice {
	var notices []entities.Notice
	repo.connection.Where(&entities.Notice{InstituteID: instituteId}).Find(&notices)
	return notices
}
