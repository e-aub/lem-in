package main

import (
	"fmt"
	"sort"
)

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
	for i, path := range paths {
		fmt.Printf("[%d]%v\n", count[i], path)
	}
	final := []ScoredPaths{}

	for n, path := range paths {
		ff := false
		for _, room := range path {
			if flag, index := containInOtherPath(room, final); flag {
				ff = true
				final[index] = ScoredPaths{Path: path, Score: count[n]}
				break
			}
		}
		if !ff {
			final = append(final, ScoredPaths{Path: path, Score: count[n]})
		}
	}
	sort.Slice(final, func(i, j int) bool {
		return len(final[i].Path)+final[i].Score > len(final[j].Path)+final[j].Score
	})
	return final
}

func countOccurrences(slices [][]string) []int {
	result := make([]int, len(slices))

	for i, slice := range slices {
		count := 0
		for _, otherSlice := range slices {
			// if i == k {
			// 	continue
			// }
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
