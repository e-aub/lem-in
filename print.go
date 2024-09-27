package main

import (
	"fmt"
	"sort"
)

func RunAnts(colony Colony, pathsSets [][][]string) {
	var TemPaths [][]string
	for _, set := range pathsSets {
		if colony.Ants > len(set) {
			if len(set) > len(TemPaths) {
				TemPaths = set
			}
		} else {
			TemPaths = pathsSets[0]
		}

	}

	var paths []Path
	for i, path := range TemPaths {
		paths = append(paths, Path{Path: path, Index: i})
	}
	// fmt.Println(paths)

	ants := make([]Ant, 0, colony.Ants)

	for n := 1; n <= colony.Ants; n++ {
		sort.Slice(paths, func(i, j int) bool {
			if len(paths[i].Path)+paths[i].AntsIn == len(paths[j].Path)+paths[j].AntsIn {
				return paths[i].Index > paths[j].Index
			}
			return len(paths[i].Path)+paths[i].AntsIn < len(paths[j].Path)+paths[j].AntsIn
		})

		ants = append(ants, Ant{Name: fmt.Sprintf("L%d", n), Next: 1, Path: paths[0].Path})
		paths[0].AntsIn++

	}

	rooms := make(map[string]string)
	for len(ants) > 0 {
		for i := 0; i < len(ants); i++ {
			ant := ants[i]
			if rooms[ant.Path[ant.Next]] == "" {
				fmt.Printf("%s-%s ", ant.Name, ant.Path[ant.Next])

				if ant.Next < len(ant.Path)-1 {
					rooms[ant.Path[ant.Next]] = ant.Name
				}

				if ant.Next > 0 {
					rooms[ant.Path[ant.Next-1]] = ""
				}

				if ant.Next < len(ant.Path)-1 {
					ants[i].Next++
				}
			}

			if ant.Path[ant.Next] == colony.End {
				ants = append(ants[:i], ants[i+1:]...)
				i--
			}
		}

		fmt.Print("\n")
	}
}
