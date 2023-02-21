package main

import (
	"fmt"
	"github.con/KlarkGames/GoLearning/greeting"
	"log"
	"rsc.io/quote"
)

func main() {
	message, err := greeting.Greeting("Алексей")

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(message)
	fmt.Println(quote.Go())
}
