package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Day24() (int, string) {
	start := time.Now()
	parseInputDay24()
	part1, part2 := solutionDay24()
	elapsed := time.Since(start)
	fmt.Println("Solution day 24 time:", elapsed)
	return part1, part2
}

var knownVariables map[string]int

type logicGate struct {
	var1, var2 string
	op         string
	res        string
}

var gates Set[logicGate]

func parseInputDay24() {
	file, err := os.Open("input/2024/day24.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	knownVariables = make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, ": ")
		value, _ := strconv.Atoi(parts[1])
		knownVariables[parts[0]] = value
	}
	regex := regexp.MustCompile(`(\w+)\s+(XOR|OR|AND)\s+(\w+)\s+->\s+(\w+)`)

	gates = Set[logicGate]{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		matches := regex.FindStringSubmatch(line)
		gates.Add(logicGate{matches[1], matches[3], matches[2], matches[4]})
	}
}

func solutionDay24() (int, string) {
	return part1(), part2()
}

func part1() int {
	var zs = make(map[int]int)
	gatesCopy := gates.Copy()
	for len(gatesCopy) > 0 {
		for _, gate := range gatesCopy.Items() {
			_, exists1 := knownVariables[gate.var1]
			_, exists2 := knownVariables[gate.var2]
			if exists1 && exists2 {
				knownVariables[gate.res] = gate.getResult()
				if strings.HasPrefix(gate.res, "z") {
					value, _ := strconv.Atoi(gate.res[1:])
					zs[value] = gate.getResult()
				}
				gatesCopy.Remove(gate)
			}
		}
	}

	bytesString := ""
	for i := 0; i < len(zs); i++ {
		bytesString = strconv.Itoa(zs[i]) + bytesString
	}
	binaryInt, _ := strconv.ParseInt(bytesString, 2, 64)
	return int(binaryInt)
}

func part2() string {
	gatesCopy := gates.Copy()
	wrongGates := Set[logicGate]{}
	for _, gate := range gatesCopy.Items() {
		if gate.op != "XOR" && strings.HasPrefix(gate.res, "z") && !strings.HasPrefix(gate.res, "z45") {
			wrongGates.Add(gate)
		}
		if !strings.HasPrefix(gate.res, "z") &&
			!strings.HasPrefix(gate.var1, "x") &&
			!strings.HasPrefix(gate.var1, "y") &&
			!strings.HasPrefix(gate.var2, "x") &&
			!strings.HasPrefix(gate.var2, "y") &&
			gate.op == "XOR" {
			wrongGates.Add(gate)
		}
	}

	for _, gate := range gatesCopy.Items() {
		if gate.op == "XOR" &&
			(strings.HasPrefix(gate.var1, "x") ||
				strings.HasPrefix(gate.var1, "y")) &&
			(strings.HasPrefix(gate.var2, "x") ||
				strings.HasPrefix(gate.var2, "y")) {
			foundGate := false
			for _, secondGate := range gatesCopy.Items() {
				if secondGate.op == "XOR" && (secondGate.var1 == gate.res || secondGate.var2 == gate.res) {
					foundGate = true
					break
				}
			}
			if !foundGate {
				if !wrongGates.Contains(gate) {

					println("test")
				}
				wrongGates.Add(gate)
			}
		} else if gate.op == "AND" {
			foundGate := false
			for _, secondGate := range gatesCopy.Items() {
				if secondGate.op == "OR" && (secondGate.var1 == gate.res || secondGate.var2 == gate.res) {
					foundGate = true
					break
				}
			}
			if !foundGate {
				if !wrongGates.Contains(gate) {
					println("test1234")
				}
				wrongGates.Add(gate)
			}
		}
	}
	outputList := []string{}
	for _, gate := range wrongGates.Items() {
		if gate.var1 == "x00" || gate.var2 == "x00" || gate.var1 == "y00" || gate.var2 == "y00" {
			continue
		}
		outputList = append(outputList, gate.res)
	}
	sort.Strings(outputList)
	output := ""
	for i, val := range outputList {
		output += val
		if i != len(outputList)-1 {
			output += ","
		}
	}
	return output
}

func (gate logicGate) getResult() int {
	if gate.op == "OR" {
		return OR(knownVariables[gate.var1], knownVariables[gate.var2])
	} else if gate.op == "AND" {
		return AND(knownVariables[gate.var1], knownVariables[gate.var2])
	} else if gate.op == "XOR" {
		return XOR(knownVariables[gate.var1], knownVariables[gate.var2])
	}
	return -1
}

func OR(x, y int) int {
	val := x == 1 || y == 1
	if val {
		return 1
	}
	return 0
}

func AND(x, y int) int {
	val := x == 1 && y == 1
	if val {
		return 1
	}
	return 0
}

func XOR(x, y int) int {
	val := x == y
	if val {
		return 0
	}
	return 1
}
