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
	var roomStart int
	var roomEnd int
	var name []string
	var x []int
	var y []int
	var index int
	var Ant int
	links := make(map[string][]string)
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
		if data[i] == "##start" && i != len(data)-1 {
			Start = string(data[i+1][0])
			roomStart = i + 1
			for j := roomStart; j < len(data); j++ {
				if data[j] == "##start" {
					fmt.Println("Invalid Format ducblicated start")
					return
				}
			}
		}
		if data[i] == "##end" && i != len(data)-1 {
			End = string(data[i+1][0])
			roomEnd = i + 1
			for j := roomEnd; j < len(data); j++ {
				if data[j] == "##end" {
					fmt.Println("Invalid Format dublicated end")
					return
				}
			}
		}
		if i != len(data) {
			index = i + 1
		} else {
			index = i
		}
		for j := index; j < len(data); j++ {
			if !strings.HasPrefix(data[j], "#") && data[i] == data[j] {
				fmt.Println("invalid Format")
				return

			}
		}
		if strings.Contains(data[i], "-") {
			InStoredLinks := strings.Split(data[i], "-")

			if len(InStoredLinks) != 2 {
				fmt.Println("Ivalid format of links")
				return
			} else {
				links[InStoredLinks[0]] = append(links[InStoredLinks[0]], InStoredLinks[1])
				links[InStoredLinks[1]] = append(links[InStoredLinks[1]], InStoredLinks[0])
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
}
