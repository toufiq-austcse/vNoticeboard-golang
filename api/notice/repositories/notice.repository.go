package repositories

import "fmt"

type INoticeRepository interface {
	CreateNotice()
	UpdateNotice()
	FindNotice()
	DeleteNotice()
}
type NoticeRepository struct {

}

func NewNoticeRepository() INoticeRepository {
	return NoticeRepository{}
}

func (repo NoticeRepository) UpdateNotice()  {
	fmt.Println("Notice Updated")
}


func (repo NoticeRepository) FindNotice()  {
	fmt.Println("Find Notice")
}

func (repo NoticeRepository) DeleteNotice()  {
	fmt.Println("Delete Notice")
}

func (repo NoticeRepository) CreateNotice()  {
	fmt.Println("Create Notice")
}
