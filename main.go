package main

import (
	"backend/src/configs/command"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("aplikasi berjalan pada port 8081")

	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
