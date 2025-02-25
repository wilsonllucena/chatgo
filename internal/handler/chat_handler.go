package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wilsonllucena/teacher-agent/internal/entity"
	chatgpt "github.com/wilsonllucena/teacher-agent/pkg/chatgpt"
)

func ExecuteChat(res http.ResponseWriter, req *http.Request) {
	chat := new(entity.Chat)

	if err := json.NewDecoder(req.Body).Decode(chat); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(chat)
	chatGPT := chatgpt.NewChatGPT()

	response, err := chatGPT.GenerateText(req.Context(), *chat)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))
}
