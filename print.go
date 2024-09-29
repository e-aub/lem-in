package main

import (
	"fmt"
	"sort"
)

type Tunnel struct {
	Romms [2]string
}

func RunAnts(colony Colony, paths []Path) {
	ants := make([]Ant, colony.Ants)

	for n := 1; n <= colony.Ants; n++ {
		sort.Slice(paths, func(i, j int) bool {
			return (len(paths[i].Path)+paths[i].AntsIn <= len(paths[j].Path)+paths[j].AntsIn)
		})
		paths[0].AntsIn++
	}
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
	rooms := make(map[string]int)
	var result string
	for len(ants) > 0 {
		usedTunnels := make(map[Tunnel]bool)
		for i := 0; i < len(ants); i++ {
			ant := ants[i]
			if !usedTunnels[Tunnel{Romms: [2]string{ant.Path[ant.Next-1], ant.Path[ant.Next]}}] {
				if rooms[ant.Path[ant.Next]] == 0 {
					result += fmt.Sprintf("L%d-%s ", ant.Id, ant.Path[ant.Next])
					usedTunnels[Tunnel{Romms: [2]string{ant.Path[ant.Next-1], ant.Path[ant.Next]}}] = true
					if ant.Next < len(ant.Path)-1 {
						rooms[ant.Path[ant.Next]] = ant.Id
					}

					if ant.Next > 0 {
						rooms[ant.Path[ant.Next-1]] = 0
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

		}
		result += "\n"
	}
	fmt.Print(result)
}
