package main

import (
	"log"
	"sort"
)

func (colony *Colony) FindPaths() [][]string {
	start := colony.GetRoom(colony.Start)
	end := colony.GetRoom(colony.End)
	if start == nil || end == nil {
		log.Fatalln("Start or end room not found")
	}

	paths := [][]string{}
	stack := [][]*Room{{start}}

	for len(stack) > 0 {
		path := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		current := path[len(path)-1]

		if current == end {
			temp := []string{}
			for _, room := range path {
				temp = append(temp, room.Name)
			}
			paths = append(paths, temp)
			continue
		}
		for _, adj := range current.Adjacent {
			if !containsRoom(path, adj) {
				newPath := append([]*Room{}, path...)
				newPath = append(newPath, adj)
				stack = append(stack, newPath)
			}
		}
	}
	//Sort paths

	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
	return paths
}

func containsRoom(path []*Room, adj *Room) bool {
	for _, room := range path {
		if room == adj {
			return true
		}
	}
	return false
}
