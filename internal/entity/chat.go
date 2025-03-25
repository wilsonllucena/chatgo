package entity

type Chat struct {
	Grade          string `json:"grade"`
	Subject        string `json:"subject"`
	QuestionCount  string `json:"question_count"`
	QuestionType   string `json:"question_type"`
	AdditionalInfo string `json:"additional_info"`
}

type ChatPersonal struct {
	Body string `json:"body"`
}
