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
	C := make([][]pair, 2)
	C[0] = make([]pair, 0, N)
	C[1] = make([]pair, 0, N)
	for i := 0; i < N; i++ {
		c := getNextInt(scanner) - 1
		p := getNextInt(scanner)
		if len(C[c]) == 0 {
			C[c] = append(C[c], pair{i, p})
		} else {
			// 自分より前の学籍番号までの合計に自分の点を足した値を保持
			C[c] = append(C[c], pair{i, p + C[c][len(C[c])-1].sum})
		}
	}

	Q := getNextInt(scanner)
	for q := 0; q < Q; q++ {
		l := getNextInt(scanner) - 1
		r := getNextInt(scanner) - 1

		s := []int{0, 0}
		for c := 0; c < len(C); c++ {
			if len(C[c]) == 0 {
				continue
			}
			// 学籍番号lより大きいクラス内のindexを求める
			li := sort.Search(len(C[c]), func(i int) bool { return C[c][i].n >= l })
			// 学籍番号rより大きいクラス内のindexを求める
			ri := sort.Search(len(C[c]), func(i int) bool { return C[c][i].n >= r })
			if ri == len(C[c]) {
				ri -= 1
			} else if C[c][ri].n > r {
				ri -= 1
			}

			// 合計値の前計算を利用して、範囲内の合計値を計算
			if li == 0 {
				s[c] = C[c][ri].sum
			} else {
				s[c] = C[c][ri].sum - C[c][li-1].sum
			}
		}

		fmt.Fprintf(writer, "%d %d\n", s[0], s[1])
	}
}

type pair struct {
	n   int
	sum int
}
