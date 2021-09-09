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

// 小課題1のみ実装
func solve(scanner *bufio.Scanner, writer *bufio.Writer) {
	mod := 1000000007
	N := getNextInt(scanner)
	B := getNextInt(scanner)
	K := getNextInt(scanner)

	c := make([]int, K+1)
	for i := 1; i <= K; i++ {
		c[i] = getNextInt(scanner)
	}

	dp := make([][]int, 100001)
	for i := range dp {
		dp[i] = make([]int, 31)
	}

	// 桁DPは https://qiita.com/pinokions009/items/1e98252718eeeeb5c9ab を参考に学習
	// dp[上から何桁目][現時点でのBで割った余り] = 通り数
	// 上の桁から余りの数のパターンをカウントするのは、割り算の筆算で余りを持ち越すイメージを持った
	// 例：「XYZ1」が3で割切れるかどうかはXYZを3で割った余りが2 or 5 or 8のパターン

	// 余りが0のものを最終的に求めたいので初期値としては 現時点でのBで割った余り=1 を設定
	dp[0][0] = 1
	for i := 0; i < N; i++ {
		for j := 0; j < B; j++ {
			for k := 1; k <= K; k++ {
				nex := (10*j + c[k]) % B
				dp[i+1][nex] += dp[i][j]
				dp[i+1][nex] %= mod
				// fmt.Fprintf(writer, "dp[%d][%d]:%d, c[%d]:%d, nex:%d -> dp[%d+1][%d]:%d\n", i, j, dp[i][j], k, c[k], nex, i, nex, dp[i+1][nex])
			}
		}
	}

	fmt.Fprintln(writer, dp[N][0])
}
