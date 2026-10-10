package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	minPath         = 10
	maxDepth        = 20
	maxStepsPerEdge = 500_000
)

type edge struct{ a, b int }

var (
	graph map[int][]int
	steps int
)

func main() {
	file, err := os.Open("as20000102.txt")
	if err != nil {
		fmt.Println("Файла нет в этой директории:", err)
		return
	}
	defer file.Close()

	graph = make(map[int][]int)
	sc := bufio.NewScanner(file)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		p := strings.Fields(line)
		if len(p) < 2 {
			continue
		}
		a, e1 := strconv.Atoi(p[0])
		b, e2 := strconv.Atoi(p[1])
		if e1 != nil || e2 != nil {
			continue
		}
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	seen := map[edge]bool{}
	var edges []edge

	keys := make([]int, 0, len(graph))
	for u := range graph {
		keys = append(keys, u)
	}
	sort.Ints(keys)

	for _, u := range keys {
		for _, v := range graph[u] {
			a, b := u, v
			if a > b {
				a, b = b, a
			}
			e := edge{a, b}
			if !seen[e] {
				seen[e] = true
				edges = append(edges, e)
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		si := len(graph[edges[i].a]) + len(graph[edges[i].b])
		sj := len(graph[edges[j].a]) + len(graph[edges[j].b])
		if si != sj {
			return si > sj
		}
		if edges[i].a != edges[j].a {
			return edges[i].a < edges[j].a
		}
		return edges[i].b < edges[j].b
	})

	for _, e := range edges {
		steps = 0
		visited := map[int]bool{e.a: true}
		if pathLen, ok := dfs(e.a, e.b, e.a, 0, visited); ok {
			fmt.Printf("Слабосвязанные вершины: A=%d, B=%d\n", e.a, e.b)
			fmt.Println("Длина альтернативного пути:", pathLen)
			return
		}
	}
	fmt.Println("Слабо связанные вершины с путём >= 10 не найдены")
}

func dfs(s, t, c, depth int, visited map[int]bool) (int, bool) {
	steps++
	if steps > maxStepsPerEdge {
		return 0, false 
	}

	if c == t {
		if depth >= minPath {
			return depth, true
		}
		return 0, false 
	}
	if depth >= maxDepth {
		return 0, false
	}

	for _, nb := range graph[c] {
		if visited[nb] {
			continue
		}
		if depth == 0 && nb == t {
			continue 
		}
		visited[nb] = true
		if d, ok := dfs(s, t, nb, depth+1, visited); ok {
			return d, true
		}
		visited[nb] = false
	}
	return 0, false
}