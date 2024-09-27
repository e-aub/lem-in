package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
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
	// fmt.Println(paths)
	// fmt.Println(paths)
	// RunAnts(colony.Ants, paths)
	// filtered := onlyUnique(paths)
	subSets := FindMaxNonInterferingPaths(paths)
	// for _, set := range subSets {
	// 	fmt.Println(set)
	// }
	fmt.Println(subSets)
	RunAnts(colony, subSets)
	fmt.Println(time.Since(start))
}
