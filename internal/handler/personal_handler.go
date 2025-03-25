package handler

import (
	"context"

	"github.com/wilsonllucena/chatgo/internal/entity"
	"github.com/wilsonllucena/chatgo/pkg/chatgpt"
)

func PersonalHandler(ctx context.Context, prompt entity.ChatPersonal) string {

	chatgpt := chatgpt.NewChatGPT()

	response, err := chatgpt.GenerateTextPersonal(ctx, prompt)
	if err != nil {
		return ""
	}

	return response

}
