package greeting

import "fmt"

func Greeting(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
