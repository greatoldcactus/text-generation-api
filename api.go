package textgenerationapi

type LLMApi interface {
	AddMessage(msg Message) error

	//Continue message
	Continue(msg string) (string, error)
	//resulting message will not be appended to history
	Generate() (string, error)

	ClearHistory()
	GetHistory() (History, error)

	SetSettings(APISettings any) error
	GetSettings(APISettigs any) error
}
