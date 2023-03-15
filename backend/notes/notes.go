package notes

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func (n *Note) ReadNotes() string {
	home, _ := os.UserHomeDir()
	filePath := filepath.Join(home, n.Title+".txt")
	body, err := os.ReadFile(filePath)
	if err != nil {
		log.Error(err)
		body = []byte("")
	}
	n.Body = string(body)
	return n.Body
}

func (n *Note) UpdateNotesBody(body string) error {
	n.Body = body
	return n.Save()
}

func (n *Note) UpdateNotesTitle(title string) error {
	err := os.Rename(n.Title+".txt", title+".txt")
	if err != nil {
		log.Error(err)
		return err
	}
	n.Title = title
	return nil
}

func (n *Note) Save() error {
	home, _ := os.UserHomeDir()
	saveFilePath := filepath.Join(home, n.Title+".txt")

	err := os.WriteFile(saveFilePath, []byte(n.Body), 0600)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
