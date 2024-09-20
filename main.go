package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalln("Invalid arguments\nUsage : go run . <filename>")
	}
	var colony Colony

	err := ParseFile(&colony, args[0])
	if err != nil {
		log.Fatalln(err)
	}
	colony.FindPaths()

}
