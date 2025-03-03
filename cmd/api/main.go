package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	chat_router "github.com/wilsonllucena/chatgo/internal/router"
)

// Run inicia o servidor da API
func Run() error {
	app := mux.NewRouter()

	app.HandleFunc("/chat", chat_router.ChatRouter).Methods("POST")
	app.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Pong"))
	}).Methods("GET")

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Em produção, especifique os domínios permitidos
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		Debug:          true, // Habilite em desenvolvimento para debug
	})

	// Envolver o handler com CORS
	handler := c.Handler(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000" // porta padrão se não estiver definida
	}

	log.Printf("Servidor rodando na porta %s", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
