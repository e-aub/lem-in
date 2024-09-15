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

func (colony *Colony) AddRoom(name string, typ string, cord [2]int) {
	if !colony.Contains(name, cord) {
		colony.Rooms = append(colony.Rooms, &Room{Name: name, Typ: typ, Coordinates: cord})
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
		fmt.Printf("%s[%s]  :  (%d, %d) ", room.Name, room.Typ, room.Coordinates[0], room.Coordinates[1])
		for _, adj := range room.Adjacent {
			fmt.Printf("__%s", adj.Name)
		}
		fmt.Println()
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
	colony.AddRoom("1", "start", [2]int{23, 3})
	colony.AddRoom("2", "room", [2]int{16, 7})
	colony.AddRoom("3", "room", [2]int{16, 3})
	colony.AddRoom("4", "room", [2]int{16, 5})
	colony.AddRoom("5", "room", [2]int{9, 3})
	colony.AddRoom("6", "room", [2]int{1, 5})
	colony.AddRoom("7", "room", [2]int{4, 8})
	colony.AddRoom("0", "end", [2]int{9, 5})

	colony.AddTunnels("0", "4")
	colony.AddTunnels("0", "6")
	colony.AddTunnels("1", "3")
	colony.AddTunnels("4", "3")
	colony.AddTunnels("5", "2")
	colony.AddTunnels("3", "5")
	colony.AddTunnels("4", "2")
	colony.AddTunnels("2", "1")
	colony.AddTunnels("7", "6")
	colony.AddTunnels("7", "2")
	colony.AddTunnels("7", "4")
	colony.AddTunnels("6", "5")

	colony.Print()


	
}
