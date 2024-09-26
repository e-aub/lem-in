package main

func pathsInterfere(path1, path2 []string) bool {
	rooms1 := make(map[string]bool)
	for _, room := range path1[1 : len(path1)-1] {
		rooms1[room] = true
	}

	for _, room := range path2[1 : len(path2)-1] {
		if rooms1[room] {
			return true
		}
	}

	return false
}

func FindMaxNonInterferingPaths(paths [][]string) [][][]string {
	final := [][][]string{}
	n := len(paths)
	maxSet := [][]string{}

	queue := [][][]string{{}}

	for i := 0; i < n; i++ {
		currentPath := paths[i]
		newSubsets := [][][]string{}

		for _, subset := range queue {
			interferes := false

			for _, chosenPath := range subset {
				if pathsInterfere(currentPath, chosenPath) {
					interferes = true
					break
				}
			}

			if !interferes {
				newSubset := append([][]string{}, subset...)
				newSubset = append(newSubset, currentPath)
				newSubsets = append(newSubsets, newSubset)

				if len(newSubset) > len(maxSet) {
					maxSet = newSubset
					final = append(final, newSubset)
				}
			}
		}

		queue = append(queue, newSubsets...)
		// fmt.Println("\n\n\n", newSubsets, "\n\n\n")

	}

	return final
}
