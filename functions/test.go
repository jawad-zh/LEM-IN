package lemin

import "fmt"

func Test(rooms map[string][]string, start, end string) [][]string {
	var room []string
	var group [][]string
	var group2 [][]string
	var children []string
	children = rooms[start]
	i := 0
	for i <= len(children)-1 {
		
		group = append(group,Bfs(room, children[i], rooms, end)...)
		group2 = append(group2, group...)
		i++
	}
	group2 = ValidGroup(group2,children)
	fmt.Println(group2)
	return group2
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
				newpath := append([]string{}, room...)
				newpath = append(newpath, neibors)
				queue = append(queue, newpath)
			}
		}
	}
	return group
}
func ValidGroup(group [][]string,children []string)[][]string{
	var Result [][]string
	for i:=0 ; i < len(group); i++{
		for j:=0 ; j < len(group[i]) ; j++{
			if !child(children,group[i][j]){
				for k:=len(group)-1 ; k > 0; k--{
					for h:= len(group[k])-1; h > 0 ; h--{
						if group[i][j] == group[k][h]{
							for n:=0 ;  n < len(group) ; n++{
								Num:= (len(group))/2
								if n != k{
									fmt.Println(group[n:Num])
									// Result= append(Result, group[n:Num]...)
								}
							}
						}
					}
				} 
			}
		}
	}
	return Result
}
func child(children []string,s string)bool{
	for i:=0 ; i < len(children) ; i++{
		if s == children[i]{
			return true 
		}
	}
	return false
}
