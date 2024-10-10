package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// IsValidName checks if the given name is valid for a room.
// A valid name cannot be empty, cannot start with '#' or 'L',
// and cannot contain whitespace.
func IsValidName(name string) bool {
	if name != "" {
		if name[0] == '#' || name[0] == 'L' {
			return false
		} else {
			for _, letter := range name {
				if unicode.IsSpace(letter) {
					return false
				}
			}
		}
		return true
	}
	return false
}

// ParseFile reads a file and populates the colony's properties based on its contents.
// It expects the file to contain the number of ants, start and end definitions,
// and room definitions or tunnels. Returns an error if the file format is invalid.
func ParseFile(colony *Colony, fileName string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) < 6 {
		return fmt.Errorf("insufficient data in file")
	}

	ants, err := strconv.Atoi(lines[0])
	if err != nil || ants <= 0 || ants > 1000_000 {
		return fmt.Errorf("invalid ants number : %s", lines[0])
	}
	colony.Ants = ants

	var start, end bool
	for _, line := range lines[1:] {
		if len(line) < 3 {
			return fmt.Errorf("invalid data format : %s", line)
		}

		switch {
		case line == "##start":
			if colony.Start != "" {
				return fmt.Errorf("start has already been defined as %s", colony.Start)
			} else if end {
				return fmt.Errorf("no room after '##end'")
			}
			start = true
		case line == "##end":
			if colony.End != "" {
				return fmt.Errorf("end has already been defined as %s", colony.End)
			} else if start {
				return fmt.Errorf("no room after '##start'")
			}
			end = true

		case line[0] == '#':
			continue

		default:
			if values := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' }); len(values) == 3 {
				name, xStr, yStr := values[0], values[1], values[2]
				if !IsValidName(name) {
					return fmt.Errorf("invalid room name : %s\nrooms names should never start with the letter L or with # and must have no spaces", name)
				}
				x, err := strconv.Atoi(xStr)
				y, err2 := strconv.Atoi(yStr)
				if err != nil || err2 != nil {
					return fmt.Errorf("invalid room coordinates: %s %s %s", name, xStr, yStr)
				}

				coords := [2]int{x, y}
				if start {
					if colony.Start != "" {
						return fmt.Errorf("start has already been defined as %s", colony.Start)
					}
					colony.Start = name
					start = false
				} else if end {
					if colony.End != "" {
						return fmt.Errorf("end has already been defined as %s", colony.End)
					}
					colony.End = name
					end = false
				}
				colony.AddRoom(name, coords)

			} else if values := strings.Split(line, "-"); len(values) == 2 {
				values[0] = strings.Trim(values[0], " ")
				values[1] = strings.Trim(values[1], " ")
				colony.AddTunnels(values[0], values[1])
			} else {
				return fmt.Errorf("invalid data format : %s", line)
			}
		}
	}

	if colony.Ants == 0 || colony.Start == "" || colony.End == "" {
		return fmt.Errorf("missing ants, start or end room")
	}
	return nil
}
