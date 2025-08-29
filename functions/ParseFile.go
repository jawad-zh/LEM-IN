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
	FirstCheck := true
	links := make(map[string][]string)
	check := make(map[string]string)
	for i := 0; i < len(data); i++ {
		if !strings.HasPrefix(data[i], "#") && !strings.Contains(data[i],"-")&& !strings.Contains(data[i]," "){
			Ant = Atoi(data[i])
			break
		}
	}
	if Ant <= 0 {
		fmt.Println("Invalid Number Of Ants")
		return
	}
	for i := 0; i < len(data); i++ {
		if !foundStart && data[i] == "##start" && i != len(data)-1 {
			Start = string(data[i+1][0])
			foundStart = true
			continue
		}
		if !foundEnd && data[i] == "##end" && i != len(data)-1 {
			End = string(data[i+1][0])
			foundEnd = true
			continue

		}
		if (data[i] == "##start" && foundStart) || (data[i] == "##end" && foundEnd) {
			fmt.Println("Invalid Format ducblicated Start or End")
			return
		}
		if strings.Contains(data[i], "-") {
			InStoredLinks := strings.Split(data[i], "-")

			if len(InStoredLinks) != 2 {
				fmt.Println("Ivalid format of links")
				return
			} else {
				if InStoredLinks[0] != "" && InStoredLinks[1] != "" && !strings.HasPrefix(InStoredLinks[0], "#") {

					links[InStoredLinks[0]] = append(links[InStoredLinks[0]], InStoredLinks[1])
					links[InStoredLinks[1]] = append(links[InStoredLinks[1]], InStoredLinks[0])
				} else if strings.HasPrefix(InStoredLinks[0], "#") {
					continue
				} else {
					fmt.Println("Ivalid format of links")
					return
				}
			}
		}
		InstoredRoom := strings.Fields(data[i])
		if len(InstoredRoom) == 3 && !strings.HasPrefix(InstoredRoom[0], "#") {
			name = append(name, string(InstoredRoom[0]))
			x = append(x, Atoi(string(InstoredRoom[1])))
			y = append(y, Atoi(string(InstoredRoom[2])))
		} else if len(InstoredRoom) != 3 && Atoi(data[i]) != Ant && !strings.HasPrefix(data[i], "#") && !strings.Contains(data[i], "-")  {
					fmt.Println("Invalid Format Of Rooms")
					return
		}
		nm := InstoredRoom[0]

		if !FirstCheck {
			if _, exists := check[nm]; exists && check[nm]!= string(Ant)&& strings.Contains(check[nm],"-"){
				fmt.Println("Invalid File Format (duplicate name):", nm)
				return
			}
		}
		if !strings.HasPrefix(InstoredRoom[0],"#"){

			check[InstoredRoom[0]] = string(InstoredRoom[0])
			FirstCheck = false
		}
	}
	if Start=="" || End == ""{
		fmt.Println("Invalid Format There Is No Start Or End")
		return
	}
	fmt.Print(links)
	gra := graph{ant: Ant, start: Start, end: End, links: links}
	ro := room{name: name, x: x, y: y}
	FoundPath(gra, ro)
}
