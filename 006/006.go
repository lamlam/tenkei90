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
	N := getNextInt(scanner)
	K := getNextInt(scanner)
	S := getNextString(scanner)

	NOT_FOUND := 1000000
	numChar := 26
	// 例: 文字列 abd のときある位置より後のみた該当する文字の場所
	// |   　　　　　| a | b | c | d | e | ... |
	// | 1文字目(a) | 0 | 1 | - | 2 | - | ... |
	// | 2文字目(b) | - | 1 | - | 2 | - | ... |
	// | 3文字目(d) | - | - | - | 2 | - | ... |
	positionMap := make([][]int, 100006)
	for i := 0; i <= N; i++ {
		positionMap[i] = make([]int, numChar)
	}
	for i := 0; i < numChar; i++ {
		positionMap[N][i] = NOT_FOUND
	}
	for i := N - 1; i >= 0; i-- {
		for j := 0; j < numChar; j++ {
			if int(S[i]-'a') == j {
				positionMap[i][j] = i
			} else {
				positionMap[i][j] = positionMap[i+1][j]
			}
		}
	}
	answer := ""
	current := 0
	for i := 1; i <= K; i++ {
		for j := 0; j < 26; j++ {
			next := positionMap[current][j]
			maxPossibleLen := N - next - 1 + i
			if maxPossibleLen >= K {
				answer += string('a' + j)
				current = next + 1
				break
			}
		}
	}

	fmt.Fprintln(writer, answer)
}
