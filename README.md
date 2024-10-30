# LEM-IN : A digital version of an ant farm

[This project is from 01Talent program](https://github.com/01-edu/public/tree/master/subjects/lem-in)


A Go-based simulation of ant colony behavior within a graph of rooms and tunnels, focused on moving ants from a start room to an end room efficiently without overlap in room occupancy.

## Structs

- **Colony**: Manages rooms, the starting and ending points, and the number of ants.
- **Room**: Represents individual rooms with coordinates and adjacency relationships.
- **Ant**: Tracks each ant’s ID, assigned path, and next room.
- **Path**: Holds a path that can accommodate a certain number of ants.
- **Tunnel**: Defines a tunnel between two rooms, ensuring unique room occupancy by ants.

## Main Functions

- **AddRoom & AddTunnels**: Add rooms and tunnels to form the colony’s graph structure.
- **FindPaths**: Discovers paths from the start to the end room using a depth-first search approach, returning paths sorted by length.
- **FilterPaths**: Selects the optimal combination of paths that maximally accommodate the ants with a non-overlapping strategy.
- **PathsInterfere & getCapacity**: Helper functions to check for path overlap and calculate each path's capacity.
- **Go**: Simulates ant movement through selected paths, ensuring that no two ants occupy the same room simultaneously.

## Path Assignment and Movement

- **GroupAnts**: Allocates ants across paths based on the total ant count and path availability.
- **Go**: Manages each ant’s step-by-step progress, updating positions based on tunnel availability and tracking room occupancy.

## File Parsing

- **ParseFile**: Reads a file to initialize the colony, setting up ants, start/end rooms, and room/tunnel connections. Validations ensure correct data formatting, and constraints like room name restrictions (e.g., no whitespace, cannot start with 'L' or '#') maintain colony integrity.

---

This project provides a structured approach to simulating an ant colony's behavior within a network of connected rooms, aiming for optimal efficiency in ant traversal.

## How to Run
Make sure that Go is installed on your machine, compile the program, and run it with a valid colony file, passing its path as the first argument; for more details about file format, check the subject below the title.

## Collaboration

To collaborate on this project, feel free to fork the repository, submit pull requests, or open issues for any bugs, feature suggestions, or improvements. Contributions are welcome.
