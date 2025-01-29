package main

import (
	"fmt"
	"log"

	"github.com/Knoxknx/greetings"
)

func main() {
	log.SetPrefix("Main(): ")
	log.SetFlags(0)

	names := []string{"Knox", "Fort", "Nill"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	// message, err := greetings.Hello("Knox")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(messages)
}
