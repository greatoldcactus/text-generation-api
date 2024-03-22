package textgenerationapi

type LLMApi interface {
	AddMessage(msg Message) error

	//Continue message
	Continue(msg string) (string, error)
	//resulting message will not be appended to history
	Generate() (string, error)
	//answer one message
	Answer(msg Message) (string, error)

	ClearHistory()
	GetHistory() (History, error)

	SetSetting(setting string, val string) error
	GetSetting(setting string) (string, error)
}
