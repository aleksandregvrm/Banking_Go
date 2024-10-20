package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	todo "example.com/banking/interfaces"
	"example.com/banking/note"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}
type outputtable interface {
	saver
	displayer
}

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	task, err := getUserInput("Todo task: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	userTodo, err := todo.New(task)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userTodo)

}

func printSomething(value any) {
	fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Saving succeeded!")
	return nil
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title: ")
	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Note content: ")
	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("cannot read the text")
	}

	text = strings.TrimSpace(text)

	if text == "" {
		return "", errors.New("you should enter a value")
	}

	return text, nil
}
