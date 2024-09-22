package main

import (
	"fmt"
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
	paths := colony.FindPaths()
	fmt.Println(paths)
	// fmt.Println(paths)
	// RunAnts(colony.Ants, paths)
	filtered := onlyUnique(paths)

	// fmt.Println(filtered)
	RunAnts(colony.Ants, filtered)
}
