package main

import (
	"bufio"
	"fmt"
	"math"
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
	P := make([]point, N)
	for i := 0; i < N; i++ {
		x := getNextInt(scanner)
		y := getNextInt(scanner)

		P[i] = point{x, y}
	}

	ans := 0.0
	// 点を1つずつ原点として選び、そのときの最大の角度を探す
	for origin := range P {
		vec := make([]float64, 0, N)
		for i := 0; i < N; i++ {
			if i == origin {
				continue
			}
			vec = append(vec, P[origin].getAngle(&P[i]))
		}
		sort.Slice(vec, func(i, j int) bool { return vec[i] < vec[j] })

		maxDeg := 0.0
		for i := range vec {
			// 現在の角度から180度ずれた点が最大
			t := vec[i] + 180.0
			if t >= 360.0 {
				t -= 360.0
			}
			// 最大の角度に近い角度を探す
			c := sort.Search(len(vec), func(x int) bool { return vec[x] >= t })
			// 一致する値が見つからない場合はlen(vec)、または挿入される箇所を返すので
			// len(vec)のmodをとったものが候補
			l := len(vec)
			c1 := c % l
			c2 := (c + l - 1) % l

			cd1 := absAngle(vec[i], vec[c1])
			cd2 := absAngle(vec[i], vec[c2])

			m := math.Max(cd1, cd2)
			maxDeg = math.Max(m, maxDeg)
		}

		ans = math.Max(maxDeg, ans)
	}
	fmt.Fprintln(writer, ans)
}

// 座標pを原点として、対象の座標tへの角度を返す
func (p *point) getAngle(t *point) float64 {
	px := float64(t.x - p.x)
	py := float64(t.y - p.y)
	return atan2Degree(px, py)
}

type point struct {
	x int
	y int
}

func atan2Degree(x, y float64) float64 {
	c := math.Atan2(y, x) * 180.0 / math.Pi
	if c < 0.0 {
		return 360.0 + c
	}
	return c
}

func absAngle(d1, d2 float64) float64 {
	a := math.Abs(d1 - d2)
	if a >= 180.0 {
		return 360.0 - a
	}
	return a
}
