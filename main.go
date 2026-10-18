package main

import (
	"bufio"
	"fmt"
	"miniProject/note"
	"miniProject/todo"
	"os"
	"reflect"
	"strings"
)

type saver interface {
	Save() error
}

//	type displayer interface {
//		Display()
//	}
type outputable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.4)
	printSomething("hello world")
	fmt.Println("Welcome to Go!")
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
	}

	todoText := getUserInput("Todo Text:")
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
	}

	//userNote.DisplayNote()
	//err = saveData(userNote)
	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
	}

	//userTodo.Display()
	//err = saveData(userTodo)
	err = outputData(userTodo)
	if err != nil {
		fmt.Println(err)
	}

	printSomething(userNote)
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println(intVal)
		return
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println(floatVal)
		return
	}
	stringVal, ok := value.(string)
	if ok {
		fmt.Println(stringVal)
		return
	}

	switch value.(type) {
	case int:
		fmt.Println("integer: ", value)
	case float64:
		fmt.Println("float: ", value)
	case string:
		fmt.Println("string: ", value)
	default:
		fmt.Println("unknown type: ", reflect.TypeOf(value))
	}
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("saving failed:", err)
		return err
	}
	fmt.Println("saved successfully")
	return nil
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func getNoteData() (string, string) {
	title := getUserInput("Note Tittle:")
	content := getUserInput("Note Content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
