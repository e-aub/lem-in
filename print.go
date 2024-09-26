package main

import (
	"fmt"
)

type (
	Ant struct {
		Name string
		Path []string
		Next int
	}
)

func RunAnts(colony Colony, pathsSets [][][]string) {
	var paths [][]string
	for _, set := range pathsSets {
		if colony.Ants > len(set) {
			if len(set) > len(paths) {
				paths = set
			}
		}
	}

	ants := make([]Ant, 0, colony.Ants)
	idx := 0
	for n := 1; n <= colony.Ants; n++ {
		if idx < len(paths) {
			ants = append(ants, Ant{Name: fmt.Sprintf("L%d", n), Next: 1, Path: paths[idx]})
		} else {
			idx = 0
			ants = append(ants, Ant{Name: fmt.Sprintf("L%d", n), Next: 1, Path: paths[idx]})
		}
		idx++
	}

	rooms := make(map[string]bool)
	for len(ants) > 0 {
		for i := 0; i < len(ants); i++ {
			ant := ants[i]
			if ant.Path[ant.Next] == colony.End {
				ants = append(ants[:i], ants[i+1:]...)
				i--
			}
			if occupied := rooms[ant.Path[ant.Next]]; !occupied {
				fmt.Printf("%s-%s ", ant.Name, ant.Path[ant.Next])
				if ant.Next < len(ant.Path)-1 {
					rooms[ant.Path[ant.Next]] = true
				}

				if ant.Next > 0 {
					rooms[ant.Path[ant.Next-1]] = false
				}

				if ant.Next < len(ant.Path)-1 {
					ants[i].Next++
				}

			}

		}

		fmt.Println()
	}
}
