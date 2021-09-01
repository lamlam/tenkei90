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

var N, L, K int
var A []int

func solve(scanner *bufio.Scanner, writer *bufio.Writer) {
	N = getNextInt(scanner)
	L = getNextInt(scanner)
	K = getNextInt(scanner)

	A = make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = getNextInt(scanner)
	}

	// 答えを二分探索で探す
	left := -1
	right := L + 1
	mid := -1
	for right-left > 1 {
		mid = left + (right-left)/2
		if isCuttable(mid) {
			left = mid
			//fmt.Fprintln(writer, "answer is higher", left, mid, right)
		} else {
			right = mid
			//fmt.Fprintln(writer, "answer is lower", left, mid, right)
		}
	}
	fmt.Fprintln(writer, left)
}
func isCuttable(l int) bool {
	cnt := 0
	pre := 0
	for i := 0; i < N; i++ {
		if (A[i]-pre >= l) && (L-A[i] >= l) {
			cnt++
			pre = A[i]
		}
	}
	return cnt >= K
}
