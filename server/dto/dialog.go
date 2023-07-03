package dto



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
