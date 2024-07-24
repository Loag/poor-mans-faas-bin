package main

import (
	"faas-bin/internal"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || len(args) > 1 {
		panic("too many arguments passed in")
	}

	event := internal.HttpEventFromString(args[0])

	// print that we received the event
	fmt.Printf("Received event: %v\n", event)

	os.Stdout.WriteString(fmt.Sprintf("successfully processed %s", event.Path))
	os.Exit(0)
}
