package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

	ScoredPaths struct {
		Path  []string
		Score int
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
	finalPaths := onlyUnique(paths)

	// Print all paths found
	fmt.Println("All paths from start to end:")
	for _, path := range finalPaths {
		fmt.Printf("[%d] : ", path.Score)
		for _, room := range path.Path {
			fmt.Printf("%s --> ", room)
		}
		fmt.Println()
	}
}

// func onlyUnique(paths [][]string) [][]string {
// 	fmt.Println(paths)
// 	roomCount := make(map[string]int)
// 	for _, path := range paths {
// 		for _, room := range path {
// 			roomCount[room]++
// 		}
// 	}
// 	fmt.Println(roomCount)
// 	uniquePaths := [][]string{}
// 	count := make([]int, len(paths))
// 	for _, path := range paths {
// 		isUnique := true
// 		for i, room := range path {
// 			if i != 0 && i != len(path)-1 && roomCount[room] > 1 {
// 				count[i] += roomCount[room]
// 				isUnique = false
// 				break
// 			}
// 		}
// 		if isUnique {
// 			uniquePaths = append(uniquePaths, path)
// 		}
// 	}
// 	fmt.Println(count)
// 	return uniquePaths
// }

func onlyUnique(paths [][]string) []ScoredPaths {
	count := countOccurrences(paths)
	//sort
	for k := 0; k < len(count); k++ {
		for l := 0; l < len(count); l++ {
			if count[k] > count[l] {
				count[k], count[l] = count[l], count[k]
				paths[k], paths[l] = paths[l], paths[k]

			}
		}
	}

	for key, path := range paths {
		fmt.Printf("[%d]", count[key])
		for _, room := range path {
			fmt.Printf("%s --> ", room)
		}
		fmt.Println()
	}
	final := []ScoredPaths{}

	for n, path := range paths {
		ff := false
		for m, room := range path {
			if m != 0 && m != len(path)-1 {
				if flag, index := containInOtherPath(room, final); flag {
					ff = true
					final[index] = ScoredPaths{Path: path, Score: count[n]}
					break
				}
			}
		}
		if !ff {
			final = append(final, ScoredPaths{Path: path, Score: count[n]})
		}
	}

	return final
}

func countOccurrences(slices [][]string) []int {
	result := make([]int, len(slices))

	for i, slice := range slices {
		count := 0
		for k, otherSlice := range slices {
			if i == k {
				continue
			}
			for _, str := range slice {
				for _, otherStr := range otherSlice {
					if str == otherStr {
						count++
					}
				}
			}
		}
		result[i] = count
	}
	return result
}

func containInOtherPath(roomArg string, paths []ScoredPaths) (bool, int) {
	for i, path := range paths {
		for j, room := range path.Path {
			if room == roomArg && j != 0 && j != len(path.Path)-1 {
				return true, i
			}
		}
	}
	return false, 0
}

func containsRoom(path []*Room, room *Room) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}

func RunAnts() {

}
