package greeting

import (
	"errors"
	"fmt"
	"math/rand"
)

func Greeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name found")
	}

	message := fmt.Sprintf(randFormat(), name)
	return message, nil
}

func randFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Hello, %v. Nice to meet you here!",
		"Здравствуй, %v. Рад тебя видеть",
	}

	return formats[rand.Intn(len(formats))]
}
