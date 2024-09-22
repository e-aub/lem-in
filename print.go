package main

import (
	"fmt"
	"strings"
)

func RunAnts(ants int, paths []ScoredPaths) {
	var odd bool
	var groups int

	if ants > len(paths) {
		groups = ants / len(paths)
		if ants%len(paths) != 0 {
			odd = true
		}
	} else {
		groups = 1
		paths = paths[:ants] // Only consider the number of paths equal to ants
	}

	var walkPaths [][]RunRoom
	var antsAtStart int

	// Populate walkPaths with ants in the start room
	for i, scoredPath := range paths {
		temp := []RunRoom{}
		for j, roomName := range scoredPath.Path {
			if j == 0 {
				// Place ants at the start room
				var antsList []string
				if odd && i == 0 {
					groups++
				}
				for k := 1; k <= groups; k++ {
					antsAtStart++
					antsList = append(antsList, fmt.Sprintf("L%d", antsAtStart))
				}
				temp = append(temp, RunRoom{Name: roomName, Ants: antsList})
				if odd && i == 0 {
					groups--
				}
			} else {
				temp = append(temp, RunRoom{Name: roomName})
			}
		}
		walkPaths = append(walkPaths, temp)
	}

	// Track how many ants have reached the end room
	antsInEnd := 0

	for antsInEnd < ants { // Run until all ants reach the end
		moves := []string{} // Track all moves in this iteration

		for m := 0; m < len(walkPaths); m++ {
			path := walkPaths[m]

			// Move ants along the path in reverse order
			for n := len(path) - 2; n >= 0; n-- {
				if len(path[n].Ants) != 0 {
					move := MoveToNextRoom(path, n)
					if move != "" {
						moves = append(moves, move) // Track the move to print it later

						// If an ant reaches the end room, increment antsInEnd
						if n+1 == len(path)-1 {
							antsInEnd++
						}
					}
				}
			}
		}

		// Print all moves for this iteration
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}

// MoveToNextRoom moves an ant from the current room to the next if the next room is empty and returns the move string
func MoveToNextRoom(path []RunRoom, currentRoom int) string {
	if len(path[currentRoom].Ants) > 0 {
		// Check if the next room is empty or is the end room
		if len(path[currentRoom+1].Ants) == 0 || currentRoom+1 == len(path)-1 {
			// Move the first ant from the current room to the next
			movedAnt := path[currentRoom].Ants[0]
			path[currentRoom+1].Ants = append(path[currentRoom+1].Ants, movedAnt)

			// Remove the moved ant from the current room
			path[currentRoom].Ants = path[currentRoom].Ants[1:]

			// Return the formatted move string
			return fmt.Sprintf("%s-%s", movedAnt, path[currentRoom+1].Name)
		}
	}
	return ""
}
