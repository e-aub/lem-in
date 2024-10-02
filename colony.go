package main

import (
	"errors"
	"log"
	"sort"
)

type (
	Colony struct { // Graph representing a collection of rooms and tunnels
		Rooms []*Room
		Start string
		End   string
		Ants  int
	}
	Room struct { // Vertex representing an individual room
		Name        string
		Coordinates [2]int
		Adjacent    []*Room
	}
	Ant struct { // Represents an ant with an ID, path, and next room index
		Id   int
		Path []string
		Next int
	}

	Path struct { // Represents a path taken by an ant or ants
		Path   []string
		AntsIn int
	}
)

// AddRoom adds a new room to the colony if it doesn't already exist.
// Logs an error if the room name or coordinates are already in use.
func (colony *Colony) AddRoom(name string, cord [2]int) {
	if !colony.ColonyContains(name, cord) {
		colony.Rooms = append(colony.Rooms, &Room{Name: name, Coordinates: cord})
		return
	}
	log.Fatalf("duh! tryna add an existing room : %s\n", name)
}

// GetRoom retrieves a room by its name from the colony.
// Returns nil if the room does not exist.
func (colony *Colony) GetRoom(name string) *Room {
	for _, room := range colony.Rooms {
		if room.Name == name {
			return room
		}
	}
	return nil
}

// AddTunnels creates a bidirectional tunnel between two rooms.
// Logs an error if either room does not exist.
func (colony *Colony) AddTunnels(from, to string) {
	fromRoom := colony.GetRoom(from)
	toRoom := colony.GetRoom(to)
	if fromRoom == nil {
		log.Fatalf("room : %s doesent exist to link it with : %s\n", from, to)
	} else if toRoom == nil {
		log.Fatalf("room : %s doesent exist to link it with : %s\n", to, from)
	}
	fromRoom.Adjacent = append(fromRoom.Adjacent, toRoom)
	toRoom.Adjacent = append(toRoom.Adjacent, fromRoom)
}

// ColonyContains checks if a room with the specified name or coordinates
// already exists in the colony. Returns true if it does.
func (colony *Colony) ColonyContains(name string, cord [2]int) bool {
	for _, room := range colony.Rooms {
		if (room.Coordinates[0] == cord[0]) && (room.Coordinates[1] == cord[1]) || room.Name == name {
			return true
		}
	}
	return false
}

// PathContainsRoom checks if a given room is already present in a path.
// Returns true if the room is found in the path.
func PathContainsRoom(path []*Room, adj *Room) bool {
	for _, room := range path {
		if room == adj {
			return true
		}
	}
	return false
}

// FindPaths finds all possible paths from the start room to the end room.
// Returns a slice of Path representing all found paths.
func (colony *Colony) FindPaths() ([]Path, error) {
	start := colony.GetRoom(colony.Start)
	end := colony.GetRoom(colony.End)
	if start == nil || end == nil {
		return nil, errors.New("start or end room not found")
	}

	paths := []Path{}
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
			paths = append(paths, Path{Path: temp})
			continue
		}
		for _, adj := range current.Adjacent {
			if !PathContainsRoom(path, adj) {
				newPath := append([]*Room{}, path...)
				newPath = append(newPath, adj)
				stack = append(stack, newPath)
			}
		}
	}

	if len(paths) == 0 {
		return nil, errors.New("there is no paths from start to end")
	}
	// Sort paths

	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i].Path) <= len(paths[j].Path)
	})

	return paths, nil
}
