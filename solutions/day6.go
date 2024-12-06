package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func Day6() (int, int) {
	matrix := parseInputDay6()

	return solutionDay6(matrix, true)

}

func parseInputDay6() [][]rune {
	file, err := os.Open("input/day6.test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	var matrix [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matrix
}

func solutionDay6(matrix [][]rune, part2 bool) (int, int) {
	current_y, current_x := findStart(matrix)
	direction := UP
	count := 1
	matrix[current_y][current_x] = '+'
	next_x, next_y := current_x, current_y
	part2Counter := 0
	for {
		next_y, next_x = getNextPos(next_y, next_x, direction)
		if !pointInBound(matrix, next_y, next_x) {
			break
		}
		if checkNewPos(matrix, next_y, next_x) {
			if matrix[next_y][next_x] == '.' {
				matrix[next_y][next_x] = getSymbol(direction, false)
				count++
			}
			current_x = next_x
			current_y = next_y
		} else {
			direction = nextDirection(direction)
			matrix[current_y][current_x] = '+'
			next_x, next_y = current_x, current_y
		}
	}

	return count, part2Counter
}

func findStart(matrix [][]rune) (int, int) {
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			if matrix[i][j] == '^' {
				return i, j
			}
		}
	}
	log.Fatal("Could not find start")
	return -1, -1
}

func nextDirection(dir Direction) Direction {
	return (dir + 1) % 4
}

func checkNewPos(matrix [][]rune, y int, x int) bool {
	return !(matrix[y][x] == '#')
}

func pointInBound(matrix [][]rune, y int, x int) bool {
	return !(y < 0 || y > len(matrix)-1) || (x < 0 || x > len(matrix[0])-1)
}

func getNextPos(y int, x int, direction Direction) (int, int) {
	switch direction {
	case UP:
		y--
	case DOWN:
		y++
	case LEFT:
		x--
	case RIGHT:
		x++
	}
	return y, x
}

func getSymbol(direction Direction, opposite bool) rune {
	if direction == UP || direction == DOWN {
		if opposite {
			return '-'
		} else {
			return '|'
		}
	}
	if opposite {
		return '|'
	} else {
		return '-'
	}
}

func printGrid(matrix [][]rune) {
	for _, row := range matrix {
		printArray(row)
	}
}

func printArray(array []rune) {
	for _, value := range array {
		fmt.Printf("%c ", value)
	}
	fmt.Println()
}
