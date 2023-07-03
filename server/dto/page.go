package dto

type RequestPagingDto struct {
	PageSize  int `json:"page_size"`
	PageIndex int `json:"page_index"`
}

type ResponsePagingDto struct {
	RequestPagingDto
	TotalCount int `json:"total_count"`
}
