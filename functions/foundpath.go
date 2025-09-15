package lemin

import "fmt"

func FoundPath(rooms map[string][]string, start string, end string) [][]string {
	var group [][]string
	queue1 := [][]string{{end}}
	queue := [][]string{{start}}

	for len(queue) > 0 && len(queue1) > 0 {
		room := queue[0]
		room1 := queue1[0]
		queue = queue[1:]
		queue1 = queue1[1:]

		last := room[len(room)-1]
		last1 := room1[len(room1)-1]

		if last == last1 {
			reversed := reversePath(room1[:len(room1)-1])
			fullPath := append(room, reversed...)
			group = append(group, fullPath)
			continue
		}

		for _, neighbor := range rooms[last] {
			if !visited(neighbor, room) {
				newPath := append([]string{}, room...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
		for _, neighbor := range rooms[last1] {
			if !visited(neighbor, room1) {
				newPath := append([]string{}, room1...)
				newPath = append(newPath, neighbor)
				queue1 = append(queue1, newPath)
			}
		}
	}
	fmt.Println(group)
	return group
}

func visited(neighbor string, room []string) bool {
	for _, r := range room {
		if neighbor == r {
			return true
		}
	}
	return false
}

func reversePath(path []string) []string {
	reversed := make([]string, len(path))
	for i := range path {
		reversed[len(path)-1-i] = path[i]
	}
	return reversed
}
