package lemin

import "fmt"

func Test(rooms map[string][]string, start, end string) [][]string {
	var room []string
	var group [][]string
	var group2 [][]string
	// var first bool
	var children []string
	children = rooms[start]
	// fmt.Println(rooms[start])
	// fmt.Println(len(rooms[start]))
	// fmt.Println(children)
	i := 0
	for i <= len(children)-1 {
		
		group = append(group,Bfs(room, children[i], rooms, end)...)
		group2 = append(group2, group...)
		
		// fmt.Println("i:",i)
		i++
		// fmt.Println("i++:",i)
	}
	fmt.Println(group2)
	fmt.Println(len(group2))
	// fmt.Println("children:",children)
	return group
}

func Bfs(room []string, child string, rooms map[string][]string, end string) [][]string {
	var queue [][]string
	var group [][]string
	visit := make(map[string]bool)
	queue = append(queue, []string{child})
	for len(queue) > 0 {
		room = queue[0]
		queue = queue[1:]
		last := room[len(room)-1]
		if last == end {
			group = append(group, room)
			continue
		}

		for _, neibors := range rooms[last] {
			if !visit[neibors] {
				visit[neibors] = true
				// if neibors == end {
				// 	newpath := append([]string{}, room...)
				// 	newpath = append(newpath, neibors)
				// 	group = append(queue, newpath)
				// 	break
				// }
				newpath := append([]string{}, room...)
				newpath = append(newpath, neibors)
				queue = append(queue, newpath)
			}
		}
	}
	return group
}
