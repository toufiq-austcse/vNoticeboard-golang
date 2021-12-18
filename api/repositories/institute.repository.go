package repositories

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type InstituteRepository struct {
	connection *gorm.DB
}

func New(dbConn *gorm.DB) InstituteRepository {
	return InstituteRepository{
		connection: dbConn,
	}

}

func (repo InstituteRepository) FindByEmail(email string) entities.Institute {
	var institute entities.Institute
	repo.connection.Where("email=?", email).Take(&institute)
	return institute
}

func (repo InstituteRepository) InsertUser(institute entities.Institute) entities.Institute {
	institute.Password = hashAndSalt([]byte(institute.Password))
	repo.connection.Save(&institute)
	return institute

}

func (repo InstituteRepository) FindInstituteById(id uint64) entities.Institute {
	var institute entities.Institute
	repo.connection.Where("id=?", id).Take(&institute)
	return institute
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash password")
	}
	return string(hash)
}
