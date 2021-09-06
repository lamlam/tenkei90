package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func configure(scanner *bufio.Scanner) {
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1000005), 1000005)
}
func getNextString(scanner *bufio.Scanner) string {
	scanned := scanner.Scan()
	if !scanned {
		panic("scan failed")
	}
	return scanner.Text()
}
func getNextInt(scanner *bufio.Scanner) int {
	i, _ := strconv.Atoi(getNextString(scanner))
	return i
}
func getNextInt64(scanner *bufio.Scanner) int64 {
	i, _ := strconv.ParseInt(getNextString(scanner), 10, 64)
	return i
}
func getNextFloat64(scanner *bufio.Scanner) float64 {
	i, _ := strconv.ParseFloat(getNextString(scanner), 64)
	return i
}
func main() {
	fp := os.Stdin
	wfp := os.Stdout

	scanner := bufio.NewScanner(fp)
	configure(scanner)
	writer := bufio.NewWriter(wfp)
	defer func() {
		r := recover()
		if r != nil {
			fmt.Fprintln(writer, r)
		}
		writer.Flush()
	}()
	solve(scanner, writer)
}

var N int
var pathMap map[int][]int

func solve(scanner *bufio.Scanner, writer *bufio.Writer) {
	N = getNextInt(scanner)
	pathMap = make(map[int][]int)
	for i := 0; i < N-1; i++ {
		A := getNextInt(scanner)
		B := getNextInt(scanner)
		pathMap[A] = append(pathMap[A], B)
		pathMap[B] = append(pathMap[B], A)
	}

	_, c := searchLongestPath(1)
	score, _ := searchLongestPath(c)

	fmt.Fprintln(writer, score+1)
}

// searchLongestPath search longest path from start city by bfs and return score and cityID
func searchLongestPath(start int) (int, int) {
	queue := make([]int, 0, N)
	INF := int(1 << 6)
	score := make([]int, N)
	for i := 0; i < N; i++ {
		score[i] = INF
	}

	queue = append(queue, start)
	score[getCityIndex(start)] = 0
	for len(queue) > 0 {
		cityID := queue[0]
		queue = queue[1:]

		currentScore := score[getCityIndex(cityID)]
		for _, v := range pathMap[cityID] {
			if score[getCityIndex(v)] == INF {
				score[getCityIndex(v)] = currentScore + 1
				queue = append(queue, v)
			}
		}
	}

	maxScore := -1
	farthestCity := -1
	for i, v := range score {
		if maxScore < v {
			maxScore = v
			farthestCity = i + 1
		}
	}

	return maxScore, farthestCity
}

func getCityIndex(cityID int) int {
	return cityID - 1
}
