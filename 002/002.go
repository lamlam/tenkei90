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

	// 奇数だと閉じれないカッコがあるので何も出力しない
	if N%2 == 1 {
		return
	}

	// ビット全探索
	// https://qiita.com/3x8tacorice/items/0b8341d7fd3ff3779111
	for bits := 0; bits < (1 << uint(N)); bits++ {
		opened := 0
		closed := 0
		isValidBrackets := true
		for i := N - 1; i >= 0; i-- {
			if bits>>uint(i)&1 == 1 { // そのbitが")"と判定
				if opened <= closed {
					isValidBrackets = false
					break
				}
				closed++
			} else { // そのbitが"("と判定
				if opened >= N/2 {
					isValidBrackets = false
					break
				}
				opened++
			}
		}
		if isValidBrackets {
			printBrackets(bits, N, writer)
		}
	}
}

func printBrackets(v int, n int, writer *bufio.Writer) {
	for i := n - 1; i >= 0; i-- {
		if v>>uint(i)&1 == 1 {
			fmt.Fprintf(writer, ")")
		} else {
			fmt.Fprintf(writer, "(")
		}
	}
	fmt.Fprintf(writer, "\n")
}
