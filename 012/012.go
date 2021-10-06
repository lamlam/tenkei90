package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func configure(scanner *bufio.Scanner) {
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(make([]byte, 1000005), 1000005)
}
func getNextLine(scanner *bufio.Scanner) string {
	scanned := scanner.Scan()
	if !scanned {
		panic("scan failed")
	}
	return scanner.Text()
}
func parseLineToInt(s string) []int {
	sp := strings.Split(s, " ")
	ret := []int{}
	for i := range sp {
		v, _ := strconv.Atoi(sp[i])
		ret = append(ret, v)
	}

	return ret
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
	WH := parseLineToInt(getNextLine(scanner))
	W := WH[0]
	H := WH[1]
	Q := parseLineToInt(getNextLine(scanner))[0]

	t := make([][]bool, H)
	for i := range t {
		t[i] = make([]bool, W)
	}

	for i := 0; i < Q; i++ {
		q := parseLineToInt(getNextLine(scanner))
		query(q, t, writer)
	}
}

func query(q []int, t [][]bool, writer *bufio.Writer) {
	qti := 0
	if q[qti] == 1 {
		ax := q[1] - 1
		ay := q[2] - 1
		t[ay][ax] = true
	} else {
		a := point{q[1] - 1, q[2] - 1}
		b := point{q[3] - 1, q[4] - 1}
		if !t[a.y][a.x] || !t[b.y][b.x] {
			fmt.Fprintln(writer, "No")
			return
		}
		if bfs(t, a, b) {
			fmt.Fprintln(writer, "Yes")
		} else {
			fmt.Fprintln(writer, "No")
		}
	}
}

type point struct {
	x int
	y int
}

func bfs(t [][]bool, s point, g point) bool {
	H := len(t)
	W := len(t[0])
	dirs := []point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// goalの周りに道がなければfalseを返す
	hasPath := false
	for _, d := range dirs {
		next := point{g.x + d.x, g.y + d.y}

		if next.x < 0 || next.x >= W {
			continue
		}
		if next.y < 0 || next.y >= H {
			continue
		}
		if t[next.y][next.x] {
			hasPath = true
			break
		}
	}
	if !hasPath {
		return false
	}

	// 一度通った場所をマーク
	passed := make([][]bool, H)
	for i := 0; i < H; i++ {
		passed[i] = make([]bool, W)
	}
	passed[s.y][s.x] = true

	queue := []point{s}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if current.x == g.x && current.y == g.y {
			return true
		}

		for _, d := range dirs {
			next := point{current.x + d.x, current.y + d.y}

			if next.x < 0 || next.x >= W {
				continue
			}
			if next.y < 0 || next.y >= H {
				continue
			}
			if passed[next.y][next.x] {
				continue
			}
			if t[next.y][next.x] {
				passed[next.y][next.x] = true
				queue = append(queue, next)
			}
		}
	}

	return false
}
