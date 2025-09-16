package lemin

import (
	"fmt"
)

func JwFoundPathes(rooms map[string][]string, start string, end string) [][]string {
	var group [][]string
	queue := [][]string{{start}}
	for len(queue) > 0 {
		room := queue[0]
		queue = queue[1:]
		last := room[len(room)-1]
		if last == end {
			group = append(group, room)
			continue
		}

		for _, neighbor := range rooms[last] {
			if !visited(neighbor, room) {
				newPath := append([]string{}, room...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}

	}
	fmt.Println(queue)
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
