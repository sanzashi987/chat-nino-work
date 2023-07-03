package dto

type RequestMessagesDto struct {
	RequestPagingDto
	DialogID uint `json:"dialog_id" binding:"required"`
}

type ResponseSingleMessage struct {
	ID   uint   `json:"id"`
	Body string `json:"body"`
}

type ResponseMessages struct {
	ResponsePagingDto
	Value []ResponseSingleMessage `json:"value"`
}

type ResponseMessagesDto struct {
	Messages []ResponseMessages `json:"messages"`
}