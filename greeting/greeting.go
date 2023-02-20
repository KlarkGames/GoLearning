package greeting

import "fmt"

func greeting(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
