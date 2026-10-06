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

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid Input")
	}
	return Todo{
		Text: content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	filename := "todo.json"
	content, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, content, 0644)
}
