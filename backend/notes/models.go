package notes

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func NewNote(title, body string) *Note {
	return &Note{
		Title: title,
		Body:  body,
	}
}
