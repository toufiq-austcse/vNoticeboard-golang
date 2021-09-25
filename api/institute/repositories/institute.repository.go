package repositories

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/institute/entities"
	"gihub.com/toufiq-austcse/vNoticeboard/common/database"
)

type InstituteRepository struct {
}

func New () InstituteRepository{
	return InstituteRepository{}

}

func (repo InstituteRepository) FindByEmail(email string) (entities.Institute, error) {
	var institute entities.Institute
	err := database.DB.Where("email=?", email).Take(&institute).Error
	if err != nil {
		return institute, err
	}
	return institute, nil

}
