package notes

type Note struct {
	Title string
	Body  string
}

func NewNote(title, body string) *Note {
	return &Note{
		Title: title,
		Body:  body,
	}
}
