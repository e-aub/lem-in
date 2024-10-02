package main

import (
	"log"
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
	Ant struct {
		Id   int
		Path []string
		Next int
	}

	Path struct {
		Path   []string
		AntsIn int
	}
)

func (colony *Colony) AddRoom(name string, cord [2]int) {
	if !colony.Contains(name, cord) {
		colony.Rooms = append(colony.Rooms, &Room{Name: name, Coordinates: cord})
		return
	}
	log.Fatalf("duh! tryna add an existing room : %s\n", name)
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
		log.Fatalf("room : %s doesent exist to link it with : %s\n", from, to)
	} else if toRoom == nil {
		log.Fatalf("room : %s doesent exist to link it with : %s\n", to, from)
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
