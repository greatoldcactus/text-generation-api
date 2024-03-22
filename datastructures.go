package api

type Message struct {
	AuthorName string
	Message    string
}

type History struct {
	Messages []Message
}
