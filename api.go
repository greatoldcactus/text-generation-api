package textgenerationapi

type LLMApi interface {

	//Continue message without using history
	Continue(msg Message) (Message, error)
	//resulting message will not be appended to history
	Generate(history History) (Message, error)
	//answer to one message without using history
	Answer(msg Message) (Message, error)

	// AddMessage(msg Message) error
	// ClearHistory()
	// GetHistory() (History, error)

	SetSetting(setting string, val string) error
	GetSetting(setting string) (string, error)

	SetAnswerTokens(tokens int) error
	GetAnswerTokens() int
}

type LLMApiModels interface {
	ListModels() (models []string, err error)
	SetModel(model string) error
	GetModel() string
}
