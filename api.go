package textgenerationapi

type LLMApi interface {

	//Continue message without using history
	Continue(msg Message) (string, error)
	//resulting message will not be appended to history
	Generate(history History) (string, error)
	//answer to one message without using history
	Answer(msg Message) (string, error)

	// AddMessage(msg Message) error
	// ClearHistory()
	// GetHistory() (History, error)

	SetSetting(setting string, val string) error
	GetSetting(setting string) (string, error)
}
