package main

import (
	"backend/src/configs/command"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Hello World")
	
	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}