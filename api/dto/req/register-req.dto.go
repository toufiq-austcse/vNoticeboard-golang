package req

type RegisterReqDto struct {
	InstituteName string `json:"institute_name" binding:"required" `
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=3"`
}
