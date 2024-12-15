package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type Box struct {
	y int
	x int
}

type BigBox struct {
	left  Box
	right Box
}

type WareHouseRobot struct {
	y int
	x int
}

func (w *WareHouseRobot) move(d Direction) {
	newPos := Box{w.y, w.x}
	switch d {
	case UP:
		newPos.y--
	case DOWN:
		newPos.y++
	case RIGHT:
		newPos.x++
	case LEFT:
		newPos.x--
	}
	if obstacles.Contains(newPos) {
		return
	}
	if boxes.Contains(newPos) {
		if !newPos.move(d) {
			return
		}
	}
	w.y = newPos.y
	w.x = newPos.x
}

func (w *WareHouseRobot) canMove(d Direction) bool {
	newPos := Box{w.y, w.x}
	switch d {
	case UP:
		newPos.y--
	case DOWN:
		newPos.y++
	case RIGHT:
		newPos.x++
	case LEFT:
		newPos.x--
	}
	if obstaclesPart2.Contains(newPos) {
		return false
	}
	for _, v := range getBigBox(newPos, newPos).Items() {
		if !v.canMove(d) {
			return false
		}
	}
	return true
}

func (bb BigBox) print() {
	println(bb.left.y, bb.left.x, bb.right.y, bb.right.x)
}

func (w *WareHouseRobot) movePart2(d Direction) {
	newPos := Box{w.y, w.x}
	switch d {
	case UP:
		newPos.y--
	case DOWN:
		newPos.y++
	case RIGHT:
		newPos.x++
	case LEFT:
		newPos.x--
	}
	if obstaclesPart2.Contains(newPos) {
		return
	}
	for _, v := range getBigBox(newPos, newPos).Items() {
		v.move(d)
	}
	w.x = newPos.x
	w.y = newPos.y
}

func (b Box) move(d Direction) bool {
	newPos := Box{b.y, b.x}
	switch d {
	case UP:
		newPos.y--
	case DOWN:
		newPos.y++
	case RIGHT:
		newPos.x++
	case LEFT:
		newPos.x--
	}
	if obstacles.Contains(newPos) {
		return false
	}
	if boxes.Contains(newPos) {
		if !newPos.move(d) {
			return false
		}
	}
	boxes.Remove(b)
	boxes.Add(newPos)
	return true
}

func (bb BigBox) canMove(d Direction) bool {
	newPos := BigBox{Box{bb.left.y, bb.left.x}, Box{bb.right.y, bb.right.x}}
	switch d {
	case UP:
		newPos.left.y--
		newPos.right.y--
	case DOWN:
		newPos.left.y++
		newPos.right.y++
	case RIGHT:
		newPos.left.x++
		newPos.right.x++
	case LEFT:
		newPos.left.x--
		newPos.right.x--
	}

	if obstaclesPart2.Contains(newPos.left) || obstaclesPart2.Contains(newPos.right) {
		return false
	}
	bigBoxesInTheWay := getBigBox(newPos.left, newPos.right)
	for _, v := range bigBoxesInTheWay.Items() {
		//if v == newPos {
		//	continue
		//}
		if v != bb && !v.canMove(d) {
			return false
		}
	}
	return true
}

func (bb BigBox) move(d Direction) {
	newPos := BigBox{Box{bb.left.y, bb.left.x}, Box{bb.right.y, bb.right.x}}
	switch d {
	case UP:
		newPos.left.y--
		newPos.right.y--
	case DOWN:
		newPos.left.y++
		newPos.right.y++
	case RIGHT:
		newPos.left.x++
		newPos.right.x++
	case LEFT:
		newPos.left.x--
		newPos.right.x--
	}
	if obstaclesPart2.Contains(newPos.left) || obstaclesPart2.Contains(newPos.right) {
		return
	}
	bigBoxesInTheWay := getBigBox(newPos.left, newPos.right)
	for _, v := range bigBoxesInTheWay.Items() {
		if v != bb {
			v.move(d)
		}
	}
	boxesPart2.Remove(bb)
	boxesPart2.Add(newPos)
}

func getBigBox(a Box, b Box) Set[BigBox] {
	temp := Set[BigBox]{}
	for _, v := range boxesPart2.Items() {
		if v.left == a || v.right == a {
			temp.Add(v)
		} else if v.left == b || v.right == b {
			temp.Add(v)
		}
	}
	return temp
}

var boxes = Set[Box]{}
var obstacles = Set[Box]{}
var robot = WareHouseRobot{}
var movements []Direction

var boxesPart2 = Set[BigBox]{}
var obstaclesPart2 = Set[Box]{}
var robotPart2 = WareHouseRobot{}
var zeroBox = BigBox{Box{0, 0}, Box{0, 0}}

var height int
var width int

func Day15() (int, int) {
	start := time.Now()
	parseInputDay15()
	part1, part2 := solutionDay15()
	elapsed := time.Since(start)
	fmt.Println("Solution day 15 time:", elapsed)
	return part1, part2
}

func parseInputDay15() {
	file, err := os.Open("input/2024/day15.txt")
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
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		row := []rune(line)
		for i, v := range row {
			if v == '#' {
				obstacles.Add(Box{counter, i})
				obstaclesPart2.Add(Box{counter, i * 2})
				obstaclesPart2.Add(Box{counter, (i * 2) + 1})
			} else if v == 'O' {
				boxes.Add(Box{counter, i})
				boxesPart2.Add(BigBox{Box{counter, i * 2}, Box{counter, (i * 2) + 1}})
			} else if v == '@' {
				robot = WareHouseRobot{counter, i}
				robotPart2 = WareHouseRobot{counter, i * 2}
			}
			width = i
		}
		counter++
	}

	height = counter

	movements = []Direction{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		for _, v := range row {
			switch v {
			case '^':
				movements = append(movements, UP)
			case 'v':
				movements = append(movements, DOWN)
			case '>':
				movements = append(movements, RIGHT)
			case '<':
				movements = append(movements, LEFT)
			}
		}
	}
}

func solutionDay15() (int, int) {
	for _, direction := range movements {
		robot.move(direction)
		if robotPart2.canMove(direction) {
			robotPart2.movePart2(direction)
		}
		//printWarehouse()
	}
	score := 0
	for _, box := range boxes.Items() {
		score += 100*box.y + box.x
	}
	scorePart2 := 0
	for _, box := range boxesPart2.Items() {
		scorePart2 += 100*box.left.y + box.left.x
	}
	return score, scorePart2
}

func isBox(b Box) Direction {
	for _, v := range boxesPart2.Items() {
		if v.left == b {
			return LEFT
		}
		if v.right == b {
			return RIGHT
		}
	}
	return DOWN
}
