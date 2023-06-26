package consts

import (
	openai "github.com/sashabaranov/go-openai"
)

var SupportModels = map[string]bool{
	openai.GPT4:              true,
	openai.GPT40314:          true,
	openai.GPT432K:           true,
	openai.GPT432K0314:       true,
	openai.GPT3Dot5Turbo0301: true,
	openai.GPT3Dot5Turbo:     true,
}
