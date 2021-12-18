package req

type CreateNoticeReqDto struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	InstituteID uint64 `json:"institute_id,omitempty"`
}
type UpdateNoticeReqDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
