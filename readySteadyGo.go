package main

import (
	"fmt"
	"sort"
)

// Tunnel represents a connection between two rooms.
type Tunnel struct {
	Romms [2]string
}

// Go simulates the movement of ants through a network of paths in a colony.
// and a slice of Path objects representing available routes. The function assigns ants
// to paths based on their lengths and manages their movement while ensuring no two ants
// occupy the same room at the same time. It outputs the movements of the ants in a formatted string.
func Go(antsNumber int, paths []Path, endRoom string) {
	if len(paths) == 1 {
		paths[0].AntsIn += antsNumber
	} else {
		GroupAnts(&paths, antsNumber)
	}
	ants := make([]Ant, antsNumber)
	var n int
	for i := 0; i < len(ants); i++ {
		if n > len(paths)-1 {
			n = 0
		}
		if paths[n].AntsIn > 0 {
			ants[i] = Ant{Id: i + 1, Path: paths[n].Path, Next: 1}
			paths[n].AntsIn--
			n++
		} else {
			n++
			i--
		}
	}
	rooms := make(map[string]bool)
	var result string
	for len(ants) > 0 {
		usedTunnels := make(map[Tunnel]bool)
		for i := 0; i < len(ants); i++ {
			ant := ants[i]
			if !usedTunnels[Tunnel{Romms: [2]string{ant.Path[ant.Next-1], ant.Path[ant.Next]}}] {
				if !rooms[ant.Path[ant.Next]] {
					result += fmt.Sprintf("L%d-%s ", ant.Id, ant.Path[ant.Next])
					usedTunnels[Tunnel{Romms: [2]string{ant.Path[ant.Next-1], ant.Path[ant.Next]}}] = true
					if ant.Next < len(ant.Path)-1 {
						rooms[ant.Path[ant.Next]] = true
					}

					// if ant.Next > 1 {
					rooms[ant.Path[ant.Next-1]] = false
					// }

					// if ant.Next < len(ant.Path)-1 {
					ants[i].Next++
					// }
				}

				if ant.Path[ant.Next] == endRoom {
					ants = append(ants[:i], ants[i+1:]...)
					i--
					continue
				}
			}

		}
		result += "\n"
	}
	fmt.Print(result)
}

func GroupAnts(paths *[]Path, ants int) {
	for n := 1; n <= ants; n++ {
		// Sort paths based on the length of the path and the number of ants already in that path.
		sort.Slice(*paths, func(i, j int) bool {
			return (len((*paths)[i].Path)+(*paths)[i].AntsIn <= len((*paths)[j].Path)+(*paths)[j].AntsIn)
		})
		(*paths)[0].AntsIn++
	}
}
