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

var W int
var H int
var t [][]bool

// union find木を実装して解く
func solve(scanner *bufio.Scanner, writer *bufio.Writer) {
	WH := parseLineToInt(getNextLine(scanner))
	W = WH[0]
	H = WH[1]
	Q := parseLineToInt(getNextLine(scanner))[0]

	t = make([][]bool, H)
	for i := 0; i < H; i++ {
		t[i] = make([]bool, W)
	}

	initTree(W * H)
	for i := 0; i < Q; i++ {
		q := parseLineToInt(getNextLine(scanner))
		query(q, writer)
	}
}

func query(q []int, writer *bufio.Writer) {
	qti := 0
	if q[qti] == 1 {
		dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		a := point{q[1] - 1, q[2] - 1}
		t[a.y][a.x] = true
		for _, d := range dirs {
			next := point{a.x + d.x, a.y + d.y}
			if next.x < 0 || W <= next.x {
				continue
			}
			if next.y < 0 || H <= next.y {
				continue
			}
			if t[next.y][next.x] {
				unite(flatten(a), flatten(next))
			}
		}
	} else {
		a := point{q[1] - 1, q[2] - 1}
		b := point{q[3] - 1, q[4] - 1}
		if !t[a.y][a.x] || !t[b.y][b.x] {
			fmt.Fprintln(writer, "No")
			return
		}
		if isSameRoot(flatten(a), flatten(b)) {
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

func flatten(p point) int {
	return p.y*W + p.x
}

var parent []int
var rank []int

func initTree(size int) {
	parent = make([]int, size)
	rank = make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
}

func findRoot(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = findRoot(parent[x])
	return parent[x]
}

func isSameRoot(x, y int) bool {
	return findRoot(x) == findRoot(y)
}

func unite(x, y int) {
	x = findRoot(x)
	y = findRoot(y)
	if x == y {
		return
	}

	if rank[x] < rank[y] {
		parent[x] = y
	} else {
		parent[y] = x
		if rank[x] == rank[y] {
			rank[x]++
		}
	}
}
