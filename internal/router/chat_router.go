package chat_router

import (
	"net/http"

	"github.com/wilsonllucena/chatgo/internal/handler"
)

func ChatRouter(res http.ResponseWriter, req *http.Request) {
	handler.ExecuteChat(res, req)
}

func ConfigureRoutes(mux *http.ServeMux) {
	chatHandler := http.HandlerFunc(handler.ExecuteChat)
	mux.Handle("/chat", chatHandler)
}
