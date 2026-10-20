package main

import (
	"fmt"
)

func bfsLevelRecursive(adjMatrix [][]int, currentLevel []int, visited []bool, distances []int, distance int) {
	if len(currentLevel) == 0 {
		return
	}
	
	nextLevel := []int{}
	
	for _, vertex := range currentLevel {
		for neighbor := 0; neighbor < len(adjMatrix); neighbor++ {
			if adjMatrix[vertex][neighbor] == 1 && !visited[neighbor] {
				visited[neighbor] = true
				distances[neighbor] = distance + 1
				nextLevel = append(nextLevel, neighbor)
			}
		}
	}
	
	bfsLevelRecursive(adjMatrix, nextLevel, visited, distances, distance+1)
}

func main() {
	adjMatrix := [][]int{
		{0, 1, 1, 0, 0, 0}, // A
		{1, 0, 0, 1, 0, 0}, // B
		{1, 0, 0, 0, 1, 0}, // C
		{0, 1, 0, 0, 1, 1}, // D
		{0, 0, 1, 1, 0, 0}, // E
		{0, 0, 0, 1, 0, 0}, // F
	}
	
	N := len(adjMatrix)
	
	var startLetter string
	fmt.Print("Enter a node: (A-", string(rune('A'+N-1)), "): ")
	fmt.Scan(&startLetter)
	
	K := int(startLetter[0] - 'A')
	
	if K < 0 || K >= N {
		fmt.Println("Wrong node!")
		return
	}
	
	visited := make([]bool, N)
	distances := make([]int, N)
	for i := range distances {
		distances[i] = -1
	}
	
	visited[K] = true
	distances[K] = 0
	
	bfsLevelRecursive(adjMatrix, []int{K}, visited, distances, 0)
	
	fmt.Println("\nResult:")
	for vertex := 0; vertex < N; vertex++ {
		letter := string(rune('A' + vertex))
		if distances[vertex] != -1 {
			fmt.Printf("%s %d\n", letter, distances[vertex])
		} else {
			fmt.Printf("%s unattainable\n", letter)
		}
	}
}