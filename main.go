package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"unicode"
)

type (
	Colony struct { // Graph
		Rooms []*Room
		Start string
		End   string
		Ants  int
	}
	Room struct { // Vertex
		Name        string
		Coordinates [2]int
		Adjacent    []*Room
	}
)

func (colony *Colony) AddRoom(name string, cord [2]int) {
	if !colony.Contains(name, cord) {
		colony.Rooms = append(colony.Rooms, &Room{Name: name, Coordinates: cord})
	} else {
		fmt.Fprintln(os.Stderr, "existing room")
	}
}

func (colony *Colony) GetRoom(name string) *Room {
	for _, room := range colony.Rooms {
		if room.Name == name {
			return room
		}
	}
	return nil
}

func (colony *Colony) AddTunnels(from, to string) {
	fromRoom := colony.GetRoom(from)
	toRoom := colony.GetRoom(to)
	if fromRoom == nil {
		log.Fatalf("room %s doesent exist\n", from)
	} else if toRoom == nil {
		log.Fatalf("room %s doesent exist\n", to)
	}
	fromRoom.Adjacent = append(fromRoom.Adjacent, toRoom)
	toRoom.Adjacent = append(toRoom.Adjacent, fromRoom)
}

func (colony *Colony) Contains(name string, cord [2]int) bool {
	for _, room := range colony.Rooms {
		if (room.Coordinates[0] == cord[0]) && (room.Coordinates[1] == cord[1]) || room.Name == name {
			return true
		}
	}
	return false
}

func (colony *Colony) Print() {
	for _, room := range colony.Rooms {
		fmt.Printf("%s : (%d, %d) ", room.Name, room.Coordinates[0], room.Coordinates[1])
		for _, adj := range room.Adjacent {
			fmt.Printf("__%s", adj.Name)
		}
		fmt.Println()
	}
}

func (colony *Colony) FindPaths() {
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
			paths = append(paths, []string{})
			for _, room := range path {
				paths[len(paths)-1] = append(paths[len(paths)-1], room.Name)
			}
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
	// Print all paths found
	fmt.Println("All paths from start to end:")
	for _, path := range paths {
		for _, room := range path {
			fmt.Printf("%s ", room)
		}
		fmt.Println()
	}
	onlyUnique(paths)
}

// contains room string

func containsRoomString(paths *[][]string, toFindRoom string, current int) bool {
	for i, path := range *paths {
		if i != current {
			for _, room := range path {
				if room == toFindRoom {
					return true
				}
			}
		}
	}
	return false
}

// function to take only unique paths

func onlyUnique(paths [][]string) {
	len := len(paths)
	count := make([]int, len)
	for i, path := range paths {
		for _, room := range path {
			if containsRoomString(&paths, room, i) {
				count[i]++
			}
		}
	}
	fmt.Println(paths)
	fmt.Println(count)
}

// check if a room is already in the current path
func containsRoom(path []*Room, room *Room) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}

func IsValidName(name string) bool {
	if name != "" {
		if name[0] == '#' || name[0] == 'L' {
			return false
		} else {
			for _, letter := range name {
				if unicode.IsSpace(letter) {
					return false
				}
			}
		}
		return true
	} else {
		return false
	}
}

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
	colony.Print()
	colony.FindPaths()

}
