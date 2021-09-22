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
	atcoder := "atcoder"
	mod := 1000000007

	N := getNextInt(scanner)
	S := getNextString(scanner)

	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, len(atcoder)+1)
	}

	dp[0][0] = 1
	for i := 0; i < len(S); i++ {
		for j := 0; j < len(atcoder); j++ {
			dp[i+1][j] += dp[i][j]
			if S[i] == atcoder[j] {
				dp[i+1][j+1] += dp[i][j]
			}
		}
		dp[i+1][len(atcoder)] += dp[i][len(atcoder)]
		for j := 0; j <= len(atcoder); j++ {
			dp[i+1][j] %= mod
		}
	}

	fmt.Fprintln(writer, dp[len(S)][len(atcoder)])
}
