package function

import (
	"fmt"
	"strings"
)

type room struct {
	name []string
	x    []int
	y    []int
}

type graph struct {
	start string
	end   string
	links map[string][]string
}

func DataRange(data []string) graph {
	var test [][]string
	var Start string
	var End string
	var roomStart int
	var roomEnd int
	var rooms []string
	var split []string
	// var test []string
	// var name []string
	var x []int
	// var y[]int
	for i := 0; i < len(data); i++ {
		if data[i] == "##start" && i != len(data)-1 {
			Start = string(data[i+1][0])
			roomStart = i + 1
		}
		if data[i] == "##end" && i != len(data)-1 {
			End = string(data[i+1][0])
			roomEnd = i + 1
		}
	}
	for j := roomStart; j < roomEnd-1; j++ {
		rooms = append(rooms, data[j])
	}
	for i := 0; i < len(rooms); i++ {
		split = strings.Split(string(rooms[i]), " ")
		test = append(test, split)

	}
	// for i := 0; i < len(rooms); i++ {
	// 	name = append(name, string(rooms[i][0]))

	// 	for j:=0 ; j < len(rooms[i]) ; j++{
	// 		fmt.Println(rooms[j])
	// 	}

	// }
	for i:=0 ; i < len(test) ; i++{
		fmt.Println(test[i][1])
	}
	fmt.Println(test)
	fmt.Print(x)

	// fmt.Println(rooms)
	// fmt.Println(len(rooms))
	gra := graph{start: Start, end: End}
	// fmt.Printf("sart is :%s and the end is:%s",gra.start,gra.end)
	return gra
}
