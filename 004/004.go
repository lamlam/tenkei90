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

func solve(scanner *bufio.Scanner, writer *bufio.Writer) {
	H := getNextInt(scanner)
	W := getNextInt(scanner)

	A := make([][]int, H)
	B := make([][]int, H)
	for i := range A {
		A[i] = make([]int, W)
		B[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			A[i][j] = getNextInt(scanner)
		}
	}

	horizontalSum := make([]int, H)
	for i := 0; i < H; i++ {
		s := 0
		for j := 0; j < W; j++ {
			s += A[i][j]
		}
		horizontalSum[i] = s
	}
	verticalSum := make([]int, W)
	for i := 0; i < W; i++ {
		s := 0
		for j := 0; j < H; j++ {
			s += A[j][i]
		}
		verticalSum[i] = s
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			B[i][j] = horizontalSum[i] + verticalSum[j] - A[i][j]
			if j == W-1 {
				fmt.Fprintf(writer, "%d\n", B[i][j])
			} else {
				fmt.Fprintf(writer, "%d ", B[i][j])
			}
		}
	}
}
