package chat_router

import (
	"encoding/json"
	"net/http"

	"github.com/wilsonllucena/chatgo/internal/entity"
	"github.com/wilsonllucena/chatgo/internal/handler"
)

type ChatResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    string `json:"data,omitempty"`
}

func ChatRouter(res http.ResponseWriter, req *http.Request) {
	chat := new(entity.Chat)

	if err := json.NewDecoder(req.Body).Decode(chat); err != nil {
		response := ChatResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		}
		sendJSONResponse(res, http.StatusBadRequest, response)
	}

	response := handler.ExecuteChat(req.Context(), chat)

	successResponse := ChatResponse{
		Success: true,
		Data:    response,
	}

	sendJSONResponse(res, http.StatusOK, successResponse)
}

func PersonalRouter(res http.ResponseWriter, req *http.Request) {
	personal := new(entity.ChatPersonal)

	if err := json.NewDecoder(req.Body).Decode(personal); err != nil {
		response := ChatResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		}
		sendJSONResponse(res, http.StatusBadRequest, response)
	}

	response := handler.PersonalHandler(req.Context(), *personal)

	successResponse := ChatResponse{
		Success: true,
		Data:    response,
	}

	sendJSONResponse(res, http.StatusOK, successResponse)
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success":false,"message":"Error encoding response"}`))
	}
}
