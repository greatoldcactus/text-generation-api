package textgenerationapi

type Message struct {
	AuthorName string
	Message    string
}

type History struct {
	Messages []Message
}
