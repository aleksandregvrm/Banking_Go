package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"created_at"`
}

// Display method to show the note details
func (note Note) Display() {
	fmt.Printf("\nYour note titled '%v' has the following content:\n\n%v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

// New function to create a new note
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("you should provide both a title and content")
	}
	return Note{
		Title:    title,
		Content:  content,
		CreateAt: time.Now(),
	}, nil
}
