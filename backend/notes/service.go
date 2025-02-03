package notes

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type NotesService struct {
	directory string
}

func NewNotesService(directory string) *NotesService {
	return &NotesService{
		directory: directory,
	}
}

func (n *NotesService) ReadNotes(title string) string {
	if title == "" {
		return ""
	}
	path := filepath.Join(n.directory, title+".txt")

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if _, err := os.Create(path); err != nil {
			log.Error(err)
		}
	}

	body, err := os.ReadFile(path)
	if err != nil {
		log.Error(err)
		body = []byte("")
	}
	return string(body)
}

func (n *NotesService) UpdateNotes(title, body string) error {
	path := filepath.Join(n.directory, title+".txt")

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if _, err := os.Create(path); err != nil {
			log.Error(err)
		}
	}

	if err := os.WriteFile(path, []byte(body), 0600); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (n *NotesService) RenameNotes(oldTitle, newTitle string) error {
	oldTitle = filepath.Join(n.directory, oldTitle+".txt")
	newTitle = filepath.Join(n.directory, newTitle+".txt")

	err := os.Rename(oldTitle, newTitle)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (n *NotesService) ChangeDirectory(newDirectory string) {
	if !filepath.IsAbs(newDirectory) {
		log.Error("Directory path must be absolute")
		return
	}
	n.directory = newDirectory
}
