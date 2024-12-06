package solutions

import (
	"bufio"
	"log"
	"os"
)

var matrix [][]rune

func Day6() (int, int) {
	parseInputDay6()

	return solutionDay6()

}

func parseInputDay6() {
	file, err := os.Open("input/day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func solutionDay6() (int, int) {
	start_row, start_col := -1, -1
	for row, _ := range matrix {
		for col, _ := range matrix[row] {
			if matrix[row][col] == '^' {
				start_row, start_col = row, col
			}
		}
	}
	if start_col == -1 {
		log.Fatalln("No start found")
	}

	state := State{Location{y: start_row, x: start_col}, UP}
	visitedLocations := Set[Location]{}
	for inGrid(state) {
		visitedLocations.Add(state.location)
		nState := nextState(state)
		if !inGrid(nState) {
			break
		}
		if isObstacle(nState) {
			state.direction = nextDir(state.direction)
		} else {
			state = nState
		}
	}

	count := 0
	for _, v := range visitedLocations.Items() {
		matrix[v.y][v.x] = '#'
		if isLoop(State{Location{y: start_row, x: start_col}, UP}) {
			count++
		}
		matrix[v.y][v.x] = '.'
	}
	return len(visitedLocations.Items()), count
}

func nextState(state State) State {
	switch state.direction {
	case UP:
		state.location.y--
	case DOWN:
		state.location.y++
	case LEFT:
		state.location.x--
	case RIGHT:
		state.location.x++
	}
	return state
}

func nextDir(dir Direction) Direction {
	return (dir + 1) % 4
}

func isObstacle(state State) bool {
	return matrix[state.location.y][state.location.x] == '#'
}

func inGrid(state State) bool {
	return state.location.x < len(matrix[0]) &&
		state.location.x >= 0 &&
		state.location.y < len(matrix) &&
		state.location.y >= 0
}

func isLoop(state State) bool {
	visited := Set[State]{}
	for inGrid(state) {
		if visited.Contains(state) {
			return true
		}
		visited.Add(state)
		nState := nextState(state)
		if !inGrid(nState) {
			return false
		}
		if isObstacle(nState) {
			state.direction = nextDir(state.direction)
		} else {
			state = nState
		}
	}
	return false
}
