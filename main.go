package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/wilsonllucena/teacher-agent/cmd/api"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// O canal para capturar sinais de interrupção é necessário para um desligamento gracioso do servidor.
	// Embora não seja estritamente obrigatório, é uma boa prática de programação pois:
	// 1. Permite que o servidor finalize suas operações pendentes antes de encerrar
	// 2. Evita a interrupção abrupta de conexões ativas
	// 3. Possibilita logging adequado do encerramento
	// Canal para capturar sinais de interrupção (SIGINT, SIGTERM)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Inicia o servidor em uma goroutine
	go func() {
		if err := api.Run(); err != nil {
			log.Fatalf("Erro ao iniciar o servidor: %v", err)
		}
	}()

	// Aguarda sinal de interrupção
	<-sigChan
	log.Println("Encerrando o servidor...")
}
