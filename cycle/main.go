package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("square_graph_dataset.txt")
	if err != nil {
		fmt.Println("Файл нет в текущей директории:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	var tokens []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		tokens = append(tokens, strings.Fields(line)...)
	}
	idx := 0
	T, err := strconv.Atoi(tokens[idx])
	if err != nil {
		fmt.Println("Первый токен не число:", tokens[idx])
		return
	}
	idx++

	results := make([]int, 0, T)

	for t := 0; t < T; t++ {
		if idx+1 >= len(tokens) {
			break
		}

		_, _ = strconv.Atoi(tokens[idx])
		E, err := strconv.Atoi(tokens[idx+1])
		if err != nil {
			fmt.Println("Ошибка в заголовке графа:", tokens[idx], tokens[idx+1])
			return
		}
		idx += 2

		edges := make([][2]int, 0, E)
		for i := 0; i < E; i++ {
			if idx+1 >= len(tokens) {
				break
			}
			a, _ := strconv.Atoi(tokens[idx])
			b, _ := strconv.Atoi(tokens[idx+1])
			idx += 2
			if a == b {
				continue
			}
			edges = append(edges, [2]int{a, b})
		}

		results = append(results, hasCycle4DFS(edges))
	}
	fileOut, _ := os.Create("output.txt")
	defer fileOut.Close()
	writer := bufio.NewWriter(fileOut)

	for _, num := range results {
		_, err := writer.WriteString(strconv.Itoa(num) + "\n")
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
	fmt.Println("Successful - output.txt")
}

func hasCycle4DFS(edges [][2]int) int {
	graph := make(map[int][]int)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	for s := range graph {
		visited := map[int]bool{s: true}
		if dfs(s, s, 0, graph, visited) {
			return 1
		}
	}
	return -1
}

func dfs(start, cur, depth int, graph map[int][]int, visited map[int]bool) bool {
	if depth == 3 {
		for _, nb := range graph[cur] {
			if nb == start {
				return true
			}
		}
		return false
	}
	for _, nb := range graph[cur] {
		if visited[nb] {
			continue
		}
		visited[nb] = true
		if dfs(start, nb, depth+1, graph, visited) {
			return true
		}
		visited[nb] = false
	}
	return false
}