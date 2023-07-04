package dto

type UpdateUserConfigDto struct {
	ChatConfig       string `json:"chat_config"`
	PreferenceConfig string `json:"preference_config"`
}

type UserConfig struct {
	Preference UserPreference `json:"preference"`
	Chat       ChatConfig     `json:"chat"`
}

type UserPreference struct {
	Avatar           string `json:"avatar"`
	SendKey          string `json:"send_key"`
	Theme            string `json:"theme"`
	Language         string `json:"language"`
	FontSize         string `json:"font_size"`
	SendPreviewubble bool   `json:"send_preview_bubble"`
	Mask             bool   `json:"mask"`
}
