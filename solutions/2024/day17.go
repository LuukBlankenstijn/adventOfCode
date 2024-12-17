package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Computer struct {
	A       int64
	B       int64
	C       int64
	IP      int64
	program []int64
	output  []int64
}

func (c *Computer) getInstruction() (int64, int64) {
	if c.IP+1 > int64(len(c.program)) {
		return -1, -1
	}
	return c.program[c.IP], c.program[c.IP+1]
}

func (c *Computer) getCombo(operand int64) int64 {
	if operand < 0 {
		return -1
	}
	if operand <= 3 {
		return operand
	}
	switch operand {
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	}
	return -1
}

func (c *Computer) adv(operand int64) {
	c.A = c.A / intPow(2, c.getCombo(operand))
	c.IP += 2
}

func (c *Computer) bxl(operand int64) {
	c.B = c.B ^ operand
	c.IP += 2
}

func (c *Computer) bst(operand int64) {
	c.B = c.getCombo(operand) % 8
	c.IP += 2
}

func (c *Computer) jnz(operand int64) {
	if c.A == 0 {
		c.IP += 2
		return
	}
	c.IP = operand
}

func (c *Computer) bxc() {
	c.B = c.B ^ c.C
	c.IP += 2
}

func (c *Computer) out(operand int64) {
	value := c.getCombo(operand) % 8
	c.output = append(c.output, value)
	c.IP += 2
}

func (c *Computer) bdv(operand int64) {
	c.B = c.A / intPow(2, c.getCombo(operand))
	c.IP += 2
}

func (c *Computer) cdv(operand int64) {
	c.C = c.A / intPow(2, c.getCombo(operand))
	c.IP += 2
}

func (c *Computer) printOutput() {
	for index, v := range c.output {
		print(v)
		if index != len(c.output)-1 {
			print(",")
		}
	}
	println("")
}

func (c *Computer) run() {
	instruction, operand := c.getInstruction()
	for instruction > -1 {
		switch instruction {
		case 0:
			c.adv(operand)
		case 1:
			c.bxl(operand)
		case 2:
			c.bst(operand)
		case 3:
			c.jnz(operand)
		case 4:
			c.bxc()
		case 5:
			c.out(operand)
		case 6:
			c.bdv(operand)
		case 7:
			c.cdv(operand)
		}
		instruction, operand = c.getInstruction()
	}
}

func (c *Computer) reset() {
	resetC := parseInputDay17()
	c.A = resetC.A
	c.B = resetC.B
	c.C = resetC.C
	c.IP = resetC.IP
	c.program = resetC.program
	c.output = resetC.output
}

var computer Computer

func Day17() (int64, int64) {
	start := time.Now()
	parseInputDay17()
	part1, part2 := solutionDay17()
	elapsed := time.Since(start)
	fmt.Println("Solution day 17 time:", elapsed)
	return part1, part2
}

func parseInputDay17() Computer {
	file, err := os.Open("input/2024/day17.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	computer = Computer{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Register A") {
			value, _ := strconv.Atoi(strings.Split(line, " ")[2])
			computer.A = int64(value)
		}
		if strings.HasPrefix(line, "Register B") {
			value, _ := strconv.Atoi(strings.Split(line, " ")[2])
			computer.B = int64(value)
		}
		if strings.HasPrefix(line, "Register C") {
			value, _ := strconv.Atoi(strings.Split(line, " ")[2])
			computer.C = int64(value)
		} else if strings.HasPrefix(line, "Program") {
			instructions := strings.Split(line, " ")[1]
			values := strings.Split(instructions, ",")
			for _, value := range values {
				intValue, _ := strconv.Atoi(value)
				computer.program = append(computer.program, int64(intValue))
			}
		}
	}
	computer.IP = 0
	return computer
}

func solutionDay17() (int64, int64) {
	computer.run()
	computer.printOutput()
	computer.reset()
	return 0, findPart2(computer)
}

func findPart2(c Computer) int64 {
	var a int64 = 1
	for {
		c.reset()
		c.A = a
		c.run()

		if reflect.DeepEqual(c.program, c.output) {
			return a
		}

		if len(c.program) > len(c.output) {
			a *= 2
			continue
		}

		if len(c.program) == len(c.output) {
			for j := len(c.program) - 1; j >= 0; j-- {
				if c.program[j] != c.output[j] {
					a += intPow(8, int64(j))
					break
				}
			}
		}

		if len(c.program) < len(c.output) {
			a /= 2
		}
	}
}

func intPow(base, exp int64) int64 {
	var result int64 = 1
	var i int64 = 0
	for ; i < exp; i++ {
		result *= base
	}
	return result
}
