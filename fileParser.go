package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

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
	} else {
		return false
	}
}

func ParseFile(colony *Colony, fileName string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) < 6 {
		return errors.New("insufficient data in file")
	}

	ants, err := strconv.Atoi(lines[0])
	if err != nil || ants <= 0 {
		return errors.New("invalid ants number")
	}
	colony.Ants = ants

	var start, end bool
	for _, line := range lines {
		if len(line) < 3 {
			continue
		}

		switch {
		case line == "##start":
			if colony.Start != "" {
				return fmt.Errorf("start has already been defined as %s", colony.Start)
			} else if end {
				return errors.New("no room after 'end'")
			}
			start = true
		case line == "##end":
			if colony.End != "" {
				return fmt.Errorf("end has already been defined as %s", colony.End)
			} else if start {
				return errors.New("no room after 'start'")
			}
			end = true

		case line[0] == '#' || line[0] == 'L':
			continue

		default:
			if values := strings.Split(line, " "); len(values) == 3 {
				name, xStr, yStr := values[0], values[1], values[2]
				if !IsValidName(name) {
					continue
				}
				x, err := strconv.Atoi(xStr)
				y, err2 := strconv.Atoi(yStr)
				if err != nil || err2 != nil || x < 0 || y < 0 {
					return fmt.Errorf("invalid room coordinates: %s %d %d", name, x, y)
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
				colony.AddTunnels(values[0], values[1])
			}
		}
	}

	if colony.Ants == 0 || colony.Start == "" || colony.End == "" {
		return errors.New("missing ants or start/end room")
	}

	return nil
}
