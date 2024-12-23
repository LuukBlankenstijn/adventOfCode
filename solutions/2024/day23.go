package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func Day23() (int, string) {
	start := time.Now()
	parseInputDay23()
	part1, part2 := solutionDay23()
	elapsed := time.Since(start)
	fmt.Println("Solution day 23 time:", elapsed)
	return part1, part2
}

var adjacencyList map[string]Set[string]

func parseInputDay23() {
	file, err := os.Open("input/2024/day23.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	adjacencyList = make(map[string]Set[string])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nodes := strings.Split(line, "-")
		_, exists := adjacencyList[nodes[0]]
		if !exists {
			adjacencyList[nodes[0]] = Set[string]{}
		}
		adjacencyList[nodes[0]].Add(nodes[1])
		_, exists = adjacencyList[nodes[1]]
		if !exists {
			adjacencyList[nodes[1]] = Set[string]{}
		}
		adjacencyList[nodes[1]].Add(nodes[0])
	}
}

func bronKerbosch(graph map[string]Set[string], r, p, x Set[string], result *[][]string, part1 bool) {
	if part1 && (len(p) == 0 && len(x) == 0 || len(r) == 3) {
		if len(r) == 3 {
			*result = append(*result, r.Items())
		}
		return
	}

	if !part1 && (len(p) == 0 && len(x) == 0) {
		*result = append(*result, r.Items())
		return
	}
	for _, v := range p.Items() {
		neighbors := graph[v]
		singleTon := make(Set[string])
		singleTon.Add(v)
		rCopy := r.Copy()
		rCopy.Add(v)
		bronKerbosch(
			graph,
			rCopy,
			p.Intersection(neighbors),
			x.Intersection(neighbors),
			result,
			part1,
		)
		p.Remove(v)
		x.Add(v)
	}
}

func solutionDay23() (int, string) {

	allNodes := make(Set[string])
	for _, x := range adjacencyList {
		allNodes.Union(x)
	}
	var result [][]string

	bronKerbosch(adjacencyList, make(Set[string]), allNodes, make(Set[string]), &result, true)

	sumPart1 := 0
	for _, v := range result {
		for _, n := range v {
			if strings.HasPrefix(n, "t") {
				sumPart1++
				break
			}
		}
	}

	allNodesPart2 := make(Set[string])
	for _, x := range adjacencyList {
		allNodesPart2.Union(x)
	}
	var resultPart2 [][]string

	bronKerbosch(adjacencyList, make(Set[string]), allNodesPart2, make(Set[string]), &resultPart2, false)

	var minimum []string
	for _, n := range resultPart2 {
		if len(n) > len(minimum) {
			minimum = n
		}
	}

	sort.Strings(minimum)

	endstring := ""
	for i, n := range minimum {
		if i != 0 {
			endstring += ","
		}
		endstring += n
	}

	return sumPart1, endstring
}
