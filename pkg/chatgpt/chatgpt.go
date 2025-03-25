package chatgpt

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
	"github.com/wilsonllucena/chatgo/config"
	"github.com/wilsonllucena/chatgo/internal/entity"
)

type ChatGPT struct {
	Client *openai.Client
}

func NewChatGPT() *ChatGPT {
	return &ChatGPT{
		Client: openai.NewClient(config.Config("OPENAI_API_KEY")),
	}
}

func (c *ChatGPT) GenerateText(ctx context.Context, prompt entity.Chat) (string, error) {
	resp, err := c.Client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				Content: "Você é um professor de " + prompt.Subject + "que deseja criar exercícios personalizados para seus alunos." +
					" Forneça as seguintes informações para gerar uma lista de exercícios para alunos do(a) " + prompt.Grade + ", Quantidade de questões " + prompt.QuestionCount + ", Tipo de questões " + prompt.QuestionType + ", e considere as seguintes informações se houver " + prompt.AdditionalInfo +
					". Gere apenas os exerecios com gabarito não fale nada como 'Claro! vou criar exercícios', apenas gere os exercícios com gabarito o gabarito deve ser mostrado apenas final da pagina",
			},
		},
	})
	if err != nil {
		fmt.Println("Error generating text:", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (c *ChatGPT) GenerateTextPersonal(ctx context.Context, prompt entity.ChatPersonal) (string, error) {
	resp, err := c.Client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				Content: "Você é um personal trainer profissional especializado em criar programas de treinamento personalizados. " +
					"Baseado nas informações fornecidas pelo usuário: " + prompt.Body + ". " +
					"Crie um plano de treino detalhado e adequado às necessidades específicas, incluindo séries, repetições, " +
					"intervalos de descanso e dicas de execução para cada exercício. Considere quaisquer limitações, objetivos " +
					"e equipamentos disponíveis mencionados. Forneça também recomendações de aquecimento e alongamento.",
			},
		},
	})
	if err != nil {
		fmt.Println("Error generating text:", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
