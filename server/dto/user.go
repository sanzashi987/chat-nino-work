package dto

type UpdateUserConfigDto struct {
	ChatConfig       string `json:"chat_config"`
	PreferenceConfig string `json:"preference_config"`
}
