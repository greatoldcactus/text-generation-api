package textgenerationapi

type Message struct {
	AuthorName string
	Message    string
}

type History struct {
	Messages []Message
}

func (h *History) AddMessage(msg Message) {
	h.Messages = append(h.Messages, msg)
}

func (h *History) ClearHistory() {
	h.Messages = make([]Message, 0)
}
