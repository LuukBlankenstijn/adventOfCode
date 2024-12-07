package solutions

import (
	"fmt"
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

}

func solutionDayn() (int, int) {
	return 0, 0
}
