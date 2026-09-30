package main

import (
	"bufio"
	"fmt"
	"miniProject/note"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Go!")
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
	}
	userNote.DisplayNote()
	err = userNote.Save()
	if err != nil {
		fmt.Println(err, "failed to save note")
		return
	}
	fmt.Println("Note saved successfully!")
	//fmt.Println(userNote.Title, userNote.Content)
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
