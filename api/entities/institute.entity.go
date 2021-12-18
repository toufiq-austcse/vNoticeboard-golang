package entities

type Institute struct {
	ID            uint64 `gorm:"primary_key:auto_increment" json:"id"`
	InstituteName string `gorm:"type:varchar(255)" json:"institute_name"`
	Email         string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password      string `gorm:"varchar(255)" json:"password"`
	Token         string `gorm:"-" json:"token,omitempty"`
}
