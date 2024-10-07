package main

// FilterPaths selects the best combination of paths that can accommodate the given total number of ants.
// It returns the largest combination of non-overlapping paths that can fit within the total number of ants.
func FilterPaths(paths []Path, totalAnts int) []Path {
	bestCombo := []Path{}
	remainingAnts := totalAnts
	bestComboRemaining := 0
	for i := 0; i < len(paths); i++ {
		numOfrooms := 0
		selectedPaths := []Path{}
		remainingAnts = totalAnts
		path1 := paths[i]
		selectedPaths = append(selectedPaths, path1)
		remainingAnts -= getCapacity(path1)
		numOfrooms += getCapacity(path1)
		if remainingAnts > 0 {
			for j := 0; j < len(paths); j++ {
				if j != i {
					path2 := paths[j]
					if !PathsInterfear(selectedPaths, path2) {
						selectedPaths = append(selectedPaths, path2)
						remainingAnts -= getCapacity(path2)
						numOfrooms += getCapacity(path2)

						if remainingAnts <= 0 {
							return selectedPaths
						}
					}
				}
			}
		}
		if len(selectedPaths) >= len(bestCombo) {
			if len(selectedPaths) == len(bestCombo) && numOfrooms < bestComboRemaining {
				bestComboRemaining = numOfrooms
				bestCombo = selectedPaths
			} else if len(selectedPaths) > len(bestCombo) {
				bestCombo = selectedPaths
			}
		}
	}
	return bestCombo
}

// PathsInterfear checks if adding a new path interferes with already selected paths.
// It returns true if there are overlapping rooms in the paths, false otherwise.
func PathsInterfear(paths []Path, path2 Path) bool {
	occupiedRooms := make(map[string]bool)
	for _, path1 := range paths {
		for _, room1 := range path1.Path[1 : len(path1.Path)-1] {
			occupiedRooms[room1] = true
		}
	}
	for _, room2 := range path2.Path[1 : len(path2.Path)-1] {
		if occupiedRooms[room2] {
			return true
		}
	}
	return false
}

// getCapacity returns the capacity of a path, defined as the number of rooms it can accommodate
// (excluding the start and end rooms).
func getCapacity(path Path) int {
	return len(path.Path) - 2
}
