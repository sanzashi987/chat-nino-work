package services

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/cza14h/chat-nino-work/consts"
	"github.com/cza14h/chat-nino-work/dto"
	"github.com/cza14h/chat-nino-work/model/completion"
	"github.com/sashabaranov/go-openai"
)

func ReplyFromGPT(ctx context.Context, chatMessageDto *dto.ChatMessageDto) (string, uint64, error) {
	var gptRequest openai.ChatCompletionRequest
	messages := chatMessageDto.History
	json.Unmarshal([]byte(chatMessageDto.Content), &gptRequest)

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: chatMessageDto.Content,
	})

	gptRequest.Messages = messages
	gptRequest.Stream = false

	gptConfig := openai.DefaultConfig("")
	client := openai.NewClientWithConfig(gptConfig)
	if ok := consts.SupportModels[gptRequest.Model]; ok {
		response, err := client.CreateChatCompletion(ctx, gptRequest)
		if err != nil {
			return "", 0, err
		}

		userMessageId, _, err := completion.CreateMessagePair(response.Choices[0].Message.Content, chatMessageDto.Content, chatMessageDto.DialogId)
		if err != nil {
			return "", 0, err
		}
		return response.Choices[0].Message.Content, userMessageId, nil
	}

	return "", 0, errors.New("not supported model")

}
