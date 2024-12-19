package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func Dayn() (int, int) {
	start := time.Now()
	parseInputDayn()
	part1, part2 := solutionDayn()
	elapsed := time.Since(start)
	fmt.Println("Solution day n time:", elapsed)
	return part1, part2
}

func parseInputDayn() {
	file, err := os.Open("input/2024/dayn.txt")
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
}

func solutionDayn() (int, int) {
	return 0, 0
}
