package greetings

import (
	"fmt"
	"errors"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hello, %v",name)
	return message,nil
}