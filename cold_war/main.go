package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main() {
	fmt.Print("file 3.txt: ")
	solve("3.txt")
	fmt.Print("file 46.txt: ")
	solve("46.txt")		
	fmt.Print("file und.txt: ")
	solve("und.txt")
}

func dfs(v int, graph map[int][]int,visited map[int]bool){
	visited[v] = true
	for _, u := range graph[v]{
		if !visited[u]{
			dfs(u, graph, visited)
		}
	}

}

func solve(filename string){
	file, _ := os.Open(filename)
	defer file.Close()
	graph := make(map[int][]int)
	visited := make(map[int]bool)
	var n int
	var edge int

	scanner := bufio.NewScanner(file)
	flag:= false
	for scanner.Scan(){
		line := scanner.Text()
		if !flag{
			n, _=strconv.Atoi(strings.Trim(line,"\n"))
			flag = true
			continue
		}
		parts := strings.Fields(line)
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}


	for v:=1; v< n;v++{
		if !visited[v]{
			edge++
			dfs(v, graph, visited)
		}
	}
	

	fmt.Print(edge-1,"\n")

}
