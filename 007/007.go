package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	N := getNextInt(scanner)
	A := make([]int, 0, N)
	for i := 0; i < N; i++ {
		A = append(A, getNextInt(scanner))
	}
	Q := getNextInt(scanner)
	B := make([]int, 0, Q)
	for i := 0; i < Q; i++ {
		B = append(B, getNextInt(scanner))
	}

	// 前準備として昇順にソート
	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
	for _, v := range B {
		fmt.Fprintln(writer, searchNearestRate(A, v))
	}
}

// 一番近いrateを探す
func searchNearestRate(A []int, rate int) int {
	if A[0] >= rate {
		return subAbs(A[0], rate)
	}
	if A[len(A)-1] <= rate {
		return subAbs(A[len(A)-1], rate)
	}

	// 二分探索
	l := -1
	r := len(A)
	for r-l > 1 {
		mid := l + (r-l)/2
		if A[mid] >= rate {
			r = mid
		} else {
			l = mid
		}
	}

	// 二分探索で見つけた位置の前後で差分が最小の値を返却
	minDiff := subAbs(rate, A[r])
	if r+1 < len(A) {
		d2 := subAbs(rate, A[r+1])
		if d2 < minDiff {
			minDiff = d2
		}
	}
	if r > 0 {
		d2 := subAbs(rate, A[r-1])
		if d2 < minDiff {
			minDiff = d2
		}
	}
	return minDiff
}

func subAbs(a int, b int) int {
	x := a - b
	if x < 0 {
		return -1 * x
	}
	return x
}
