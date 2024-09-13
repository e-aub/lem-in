package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	Colony struct { // Graph
		Rooms []*Room
	}
	Room struct { // Vertex
		Name        string
		Typ         string
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
		fmt.Printf("%s  :  (%d, %d)\n", room.Name, room.Coordinates[0], room.Coordinates[1])
	}
}

func main() {
	content, err := os.ReadFile("examples/example00.txt")
	if err != nil {
		log.Fatalln(err)
	}
	spliced := strings.Split(string(content), "\n")
	if len(spliced) > 1 {
		ants, err := strconv.Atoi(spliced[0])
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(ants)
	}
	var colony Colony
	for i := 0; i <= 5; i++ {
		colony.AddRoom(strconv.Itoa(i), [2]int{i, i + 1})
	}
	colony.AddTunnels("1", "2")

	colony.Print()
}
