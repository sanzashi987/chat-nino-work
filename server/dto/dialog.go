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

type DialogConfig struct {
	Config          string `json:"config"`
	DialogTitle     string `json:"dialog_title"`
	UseGlobalConfig bool   `json:"use_global_config"`
}
