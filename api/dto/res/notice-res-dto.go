package res

type NoticeResponse struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
