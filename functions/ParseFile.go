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
	ant   int
	start string
	end   string
	links map[string][]string
}

func DataRange(data []string) {
	// need to organiz
	var Start string
	var End string
	var name []string
	var x []int
	var y []int
	var Ant int
	var foundEnd bool
	var foundStart bool
	// var lineSearch string
	links := make(map[string][]string)
	// test:= make(map[int]string)
	// var count int
	for i := 0; i < len(data); i++ {
		if !strings.HasPrefix(data[i], "#") {
			Ant = Atoi(data[i])
			break
		}
	}
	if Ant <= 0 {
		fmt.Println("Invalid Number Of Ants")
		return
	}
	for i := 0; i < len(data); i++ {
		if !foundStart && data[i] == "##start" && i != len(data)-1{
				Start = string(data[i+1][0])
				foundStart = true
				continue
		}
		if !foundEnd && data[i]=="##end" && i != len(data)-1{
				End= string(data[i][0])
				foundEnd = true
				continue
			
		}
		if (data[i]== "##start"&& foundStart) || (data[i]=="##end"&& foundEnd){
			fmt.Println("Invalid Format ducblicated Start or End")
				return
		}
		//needed to optimze
		// if i != len(data) {
		// 	index = i + 1
		// } else {
		// 	index = i
		// }
		// for j := index; j < len(data); j++ {
		// 	if !strings.HasPrefix(data[j], "#") && data[i] == data[j] {
		// 		fmt.Println("invalid Format")
		// 		return

		// 	}
		// }
		if strings.Contains(data[i], "-") {
			InStoredLinks := strings.Split(data[i], "-")

			if len(InStoredLinks) != 2 {
				fmt.Println("Ivalid format of links")
				return
			} else {
				if InStoredLinks[0]!="" && InStoredLinks[1]!=""{

					links[InStoredLinks[0]] = append(links[InStoredLinks[0]], InStoredLinks[1])
					links[InStoredLinks[1]] = append(links[InStoredLinks[1]], InStoredLinks[0])
				}else{
					fmt.Println("Ivalid format of links")
				return
				}
			}
		}
		InstoredRoom := strings.Split(data[i], " ")
		if len(InstoredRoom) == 3 {
			name = append(name, string(InstoredRoom[0]))
			x = append(x, Atoi(string(InstoredRoom[1])))
			y = append(y, Atoi(string(InstoredRoom[2])))
		}
	}
		gra := graph{ant: Ant, start: Start, end: End, links: links}
		ro := room{name: name, x: x, y: y}
		Test(gra, ro)
		fmt.Print(links)
	

}