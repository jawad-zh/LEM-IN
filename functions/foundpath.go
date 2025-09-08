package lemin

func FoundPath(rooms map[string][]string, start string, end string) [][]string {
	var group [][]string
	var path []string
	visited := make(map[string]bool)
	queue := []string{start}
	visited[start] = true
	for len(queue) > 0 {
		room := queue[0]
		queue = queue[1:]
		path = append(path, room)
		if room == end {
			group = append(group, path)
			continue
		}
		for _, neighbor := range rooms[room] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return group
}
