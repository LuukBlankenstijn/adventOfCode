package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Robot struct {
	pos   RobotPos
	speed RobotVelocity
}

type RobotPos struct {
	y int
	x int
}

type RobotVelocity struct {
	y int
	x int
}

var WIDTH int
var HEIGHT int
var robots = make([]Robot, 0)

func (r *Robot) move(steps int) {
	r.pos.x += r.speed.x * steps
	r.pos.y += r.speed.y * steps
	r.pos.x = r.pos.x % WIDTH
	r.pos.y = r.pos.y % HEIGHT
	if r.pos.x < 0 {
		r.pos.x += WIDTH
	}
	if r.pos.y < 0 {
		r.pos.y += HEIGHT
	}
}

type Quadrant int

const (
	TOP_LEFT Quadrant = iota
	TOP_RIGHT
	BOTTOM_LEFT
	BOTTOM_RIGHT
)

func (r *Robot) getQuadrant() Quadrant {
	var quadrant Quadrant = -1
	if r.pos.x <= (WIDTH-3)/2 {
		if r.pos.y <= (HEIGHT-3)/2 {
			quadrant = TOP_LEFT
		} else if r.pos.y >= (HEIGHT+1)/2 {
			quadrant = BOTTOM_LEFT
		}
	} else if r.pos.x >= (WIDTH+1)/2 {
		if r.pos.y <= (HEIGHT-3)/2 {
			quadrant = TOP_RIGHT
		} else if r.pos.y >= (HEIGHT+1)/2 {
			quadrant = BOTTOM_RIGHT
		}
	}
	return quadrant
}

func Day14() (int, int) {
	start := time.Now()
	parseInputDay14()
	part1, part2 := solutionDay14()
	elapsed := time.Since(start)
	fmt.Println("Solution day 14 time:", elapsed)
	return part1, part2
}

func parseInputDay14() {
	file, err := os.Open("input/2024/day14.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	pattern := "p=(-?\\d+),(-?\\d+) v=(-?\\d+),(-?\\d+)|(\\d+) (\\d+)"
	regex := regexp.MustCompile(pattern)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindStringSubmatch(line)
		if matches[4] != "" {
			x, err1 := strconv.Atoi(matches[1])
			y, err2 := strconv.Atoi(matches[2])
			xSpeed, err3 := strconv.Atoi(matches[3])
			ySpeed, err4 := strconv.Atoi(matches[4])
			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				log.Fatal(err1, err2, err3, err4)
			}
			robot := Robot{
				pos: RobotPos{
					y: y,
					x: x,
				},
				speed: RobotVelocity{
					y: ySpeed,
					x: xSpeed,
				},
			}
			robots = append(robots, robot)

		}
		if matches[6] != "" {
			width, err1 := strconv.Atoi(matches[5])
			height, err2 := strconv.Atoi(matches[6])
			if err1 != nil || err2 != nil {
				log.Fatal(err1, err2)
			}
			WIDTH, HEIGHT = width, height
		}

	}

}

func solutionDay14() (int, int) {
	var quadrants = make([]int, 4)
	uniquePos := 0
	part1 := 0
	counter := 1
	for uniquePos == 0 || part1 == 0 {
		for index := range robots {
			robots[index].move(1)

			// part 1
			if counter == 100 {
				quadrant := robots[index].getQuadrant()
				if quadrant != -1 {
					quadrants[quadrant]++
				}
			}

		}
		// part 2
		if checkUnique(robots) {
			uniquePos = counter
		}

		if counter == 100 {
			part1 = quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
		}
		counter++
	}

	return part1, uniquePos
}

func checkUnique([]Robot) bool {
	posses := Set[RobotPos]{}
	for _, robot := range robots {
		if posses.Contains(robot.pos) {
			return false
		}
		posses.Add(robot.pos)
	}
	return true
}

func printlnRobots([]Robot) {
	grid := make([][]int, HEIGHT) // Create rows
	for i := 0; i < HEIGHT; i++ {
		grid[i] = make([]int, WIDTH) // Create columns
	}
	for _, robot := range robots {
		grid[robot.pos.y][robot.pos.x] += 1
	}
	for _, row := range grid {
		for _, cell := range row {
			print(cell)
		}
		println("")
	}
}
