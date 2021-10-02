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

	tasks := make([]task, N)
	for i := 0; i < N; i++ {
		d := getNextInt(scanner)
		c := getNextInt(scanner)
		s := getNextInt(scanner)
		tasks[i] = task{d, c, s}
	}

	// 締め切りの短い順にソート
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].d < tasks[j].d
	})

	// dp[どこまで仕事を見たか][合計仕事時間] = 報酬の最大値
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, 5009)
	}

	for i := 0; i < N; i++ {
		for j := 0; j <= 5000; j++ {
			// 仕事iをやらない場合
			dp[i+1][j] = maxInt(dp[i+1][j], dp[i][j])
			// 仕事iをやる場合
			if j+tasks[i].c <= tasks[i].d {
				sumWorkTime := j + tasks[i].c
				dp[i+1][sumWorkTime] = maxInt(dp[i+1][sumWorkTime], dp[i][j]+tasks[i].s)
			}
		}
	}

	ans := 0
	for i := 0; i <= 5000; i++ {
		ans = maxInt(ans, dp[N][i])
	}
	fmt.Fprintln(writer, ans)
}

type task struct {
	d int // 締め切り
	c int // 必要日数
	s int // 報酬
}

func maxInt(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
