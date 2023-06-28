package dto

type RequestPagingDto struct {
	PageSize  int `json:"page_size"`
	PageIndex int `json:"page_index"`
}

type ResponsePagingDto struct {
	RequestPagingDto
	TotalCount int `json:"total_count"`
}

type UpdateUserConfigDto struct {
	ChatConfig       string `json:"chat_config"`
	PreferenceConfig string `json:"preference_config"`
}

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

type RequestDialogsDto struct {
	RequestPagingDto
	// RequestMessageDto
	// UserId uint `json:"user_id"`
}

type ResponseSingleDialog struct {
}

type ResponseDialogs struct {
	ResponsePagingDto
	Value []ResponseSingleDialog `json:"value"`
}
type ResponseUserInfoDto struct {
	UpdateUserConfigDto
	Dialogs ResponseDialogs `json:"dialogs"`
}
