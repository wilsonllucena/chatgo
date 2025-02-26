package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wilsonllucena/chatgo/internal/entity"
	chatgpt "github.com/wilsonllucena/chatgo/pkg/chatgpt"
)

// ChatResponse represents the structure of the JSON response
type ChatResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    string `json:"data,omitempty"`
}

func ExecuteChat(res http.ResponseWriter, req *http.Request) {
	chat := new(entity.Chat)

	if err := json.NewDecoder(req.Body).Decode(chat); err != nil {
		response := ChatResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		}
		sendJSONResponse(res, http.StatusBadRequest, response)
		return
	}

	fmt.Println(chat)
	chatGPT := chatgpt.NewChatGPT()

	response, err := chatGPT.GenerateText(req.Context(), *chat)
	if err != nil {
		errorResponse := ChatResponse{
			Success: false,
			Message: "Error generating response: " + err.Error(),
		}
		sendJSONResponse(res, http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := ChatResponse{
		Success: true,
		Data:    response,
	}
	sendJSONResponse(res, http.StatusOK, successResponse)
}

// sendJSONResponse is a helper function to send JSON responses
func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		// If encoding fails, send a simple error message
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success":false,"message":"Error encoding response"}`))
	}
}
