package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

// Display method to show the note details
func (todo Todo) Display() {
	fmt.Printf("\nYour note titled '%v' has the following content:\n\n\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

// New function to create a new note
func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("you should provide both a title and content")
	}
	return Todo{
		Text: text,
	}, nil
}
