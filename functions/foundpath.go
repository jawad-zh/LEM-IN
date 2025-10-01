package lemin

import (
	"fmt"
	"os"
)

var allFoundPaths [][]string

func FindAllPaths() {
	var initialPaths [][]string
	for _, child := range colony.graph[colony.start] {
		if child == colony.end {
			newPath := []string{colony.start, colony.end}
			allFoundPaths = append(allFoundPaths, newPath)
			continue
		}
		newPath := append([]string{}, colony.start, child)
		initialPaths = append(initialPaths, newPath)
	}

	for index := range initialPaths {
		paths := bfsSearch(index, initialPaths)
		allFoundPaths = append(allFoundPaths, paths...)

	}

	if len(allFoundPaths) == 0 {
		fmt.Println("fff")
		os.Exit(1)
	}

	GrouppingPaths(allFoundPaths)
}

func bfsSearch(index int, initialPaths [][]string) [][]string {
	var count int = 0
	var visitedHistory  []string
	paths := [][]string{}
	queue := [][]string{initialPaths[index]}
	visited := make(map[string]struct{})
	relativeVisited := make(map[string][]string)
	visited[initialPaths[index][1]] = struct{}{}
	for len(queue) > 0 || len(initialPaths)-count> 0 {
	
		if len(queue) == 0 {
			if count == index && index < len(initialPaths)-1 {
				count++
			}
			queue = append(queue, initialPaths[count])
			visited[initialPaths[count][1]] = struct{}{}
			count++
			continue
		}
		current := queue[0]
		queue = queue[1:]
		path := current
		last := path[len(path)-1]
		
		
		for _, neighbor := range colony.graph[last] {
			if contains(path, neighbor) {
				continue
			}
			if neighbor == colony.end {
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				paths = append(paths, newPath)
				visitedHistory = append(visitedHistory, path...)
				for k := range visited {
					if !contains(visitedHistory, k) {
						delete(visited, k)
					}
				}
				for r := range relativeVisited {
					if !contains(visitedHistory, r) {
						delete(relativeVisited, r)
					}
				}
				queue = [][]string{}
				break
			}
			if len(colony.graph[last]) == 2 {
				if _,flag := visited[neighbor];flag {
					continue
				}
				if restartPath, flag := relativeVisited[neighbor]; flag {
					queue = append(queue, restartPath)
				}
				visited[neighbor] = struct{}{}
			} else {
					if _, flag := visited[neighbor]; flag {
						continue
					}
					if _, flag := relativeVisited[neighbor]; flag {
						continue
					}
					relativeVisited[neighbor] = path
			}
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)

		}

	}
	
	return paths
}

func contains(path []string, room string) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}
