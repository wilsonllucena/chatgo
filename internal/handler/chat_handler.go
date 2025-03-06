package handler

import (
	"context"

	"github.com/wilsonllucena/chatgo/internal/entity"
	chatgpt "github.com/wilsonllucena/chatgo/pkg/chatgpt"
)

func ExecuteChat(ctx context.Context, chat *entity.Chat) string {
	chatGPT := chatgpt.NewChatGPT()

	response, err := chatGPT.GenerateText(ctx, *chat)
	if err != nil {
		return ""
	}

	return response
}
