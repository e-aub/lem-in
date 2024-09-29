package main

import (
	"log"
	"os"
)

func main() {
	// start := time.Now()
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalln("Invalid arguments\nUsage : go run . <filename>")
	}
	fileName := args[0]
	// fileName := "examples/example01"
	var colony Colony

	err := ParseFile(&colony, fileName)
	if err != nil {
		log.Fatalln(err)
	}
	paths := colony.FindPaths()
	if len(paths) < 1 {
		log.Fatalln("There is no paths from start to end")
	}

	tyPaths := []Path{}
	for _, path := range paths {
		tyPaths = append(tyPaths, Path{Path: path})
	}
	// fmt.Println(tyPaths, "\n\n")
	subSets := FilterPaths(tyPaths, colony.Ants)
	RunAnts(colony, subSets)
	// RunAnts(colony, subSets)
	// fmt.Println(time.Since(start))
}
