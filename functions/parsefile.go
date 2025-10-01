package lemin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Colony struct {
	antNumber  int
	start, end string
	room map[string][]int
	graph      map[string][]string
}

var colony = Colony{
	graph: make(map[string][]string),
	
}
var checker = Colony{
	room : make(map[string][]int),
}

var fileContenue []string

func ParseFile(fileName string) {
	var finishRoomes bool
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		os.Exit(1)
	}
	fileBuffer := bufio.NewScanner(file)
	for fileBuffer.Scan() {
		line := strings.TrimSpace(fileBuffer.Text())
		
		if (line == "" || strings.HasPrefix(line, "#")) && !(line == "##start" || line == "##end") {
			continue
		}
		fileContenue = append(fileContenue, line)
	}
	if len(fileContenue) == 0{
		fmt.Println("ERROR: invalid data format")
		os.Exit(1)
	}

	for index, line := range fileContenue {
		if index == 0 {
			n, err := strconv.Atoi(strings.TrimSpace(line))
			if err != nil {
				fmt.Println("ERROR: invalid data format")
				os.Exit(1)
			}
			if n <= 0 {
				fmt.Println("ERROR: invalid data format")
				os.Exit(1)
			}
			colony.antNumber = n
		} else {
			if strings.TrimSpace(line) == "##start" && index != len(fileContenue)-1 {
				if colony.start == "" {

					startRoom := strings.Fields(fileContenue[index+1])
					if len(startRoom) == 3 {
						_, err := strconv.Atoi(startRoom[1])
						if err != nil {
							fmt.Println("ERROR: invalid data format")
							os.Exit(1)
						}
						_, err = strconv.Atoi(startRoom[2])
						if err != nil {
							fmt.Println("ERROR: invalid data format")
							os.Exit(1)
						}
						colony.start = startRoom[0]
					} else {
						fmt.Println("ERROR: invalid data format")
						os.Exit(1)
					}
				} else {
					fmt.Println("ERROR: invalid data format")
					os.Exit(1)
				}
			} else if strings.TrimSpace(line) == "##end" && index != len(fileContenue)-1 {
				if colony.end == "" {

					endRoom := strings.Fields(fileContenue[index+1])
					if len(endRoom) == 3 {
						_, err := strconv.Atoi(endRoom[1])
						if err != nil {
							fmt.Println("ERROR: invalid data format")
							os.Exit(1)
						}
						_, err = strconv.Atoi(endRoom[2])
						if err != nil {
							fmt.Println("ERROR: invalid data format")
							os.Exit(1)
						}
						colony.end = endRoom[0]
					} else {
						fmt.Println("ERROR: invalid data format")
						os.Exit(1)
					}
				} else {
					fmt.Println("ERROR: invalid data format")
					os.Exit(1)
				}
			} else if len(strings.Fields(line)) == 3 {
				if finishRoomes{
					fmt.Println("ERROR: invalid data format")
					os.Exit(1)
				}
				room := strings.Fields(line)
				x, err := strconv.Atoi(room[1])
				if err != nil {
					fmt.Println("ERROR: ivalid data format")
					os.Exit(1)
				}
				checker.room[room[0]] = append(checker.room[room[0]],x )
				y, err := strconv.Atoi(room[2])
				if err != nil {
					fmt.Println("ERROR: invalid data format")
					os.Exit(1)
				}
				checker.room[room[0]] = append(checker.room[room[0]],y )
				if _, alreadyExist := colony.graph[room[0]]; alreadyExist {
					fmt.Println("ERROR: invalid data format")
					os.Exit(1)
				}
				colony.graph[room[0]] = nil
			} else if links := strings.Split(strings.TrimSpace(line), "-"); len(links) == 2 {
				finishRoomes = true
				if _, validRoom := colony.graph[links[0]]; !validRoom {
					fmt.Println("ERROR: invalid data format")
					os.Exit(1)
				}
				if _, validRoom2 := colony.graph[links[1]]; !validRoom2 {
					fmt.Println("ERROR: invalid data format")
					os.Exit(1)
				}
				for _, room := range colony.graph[links[0]] {
					if room == links[1] {
						fmt.Println("ERROR: invalid data format")
						os.Exit(1)
					}
				}
				colony.graph[links[0]] = append(colony.graph[links[0]], links[1])
				for _, room := range colony.graph[links[1]] {
					if room == links[0] {
						fmt.Println("ERROR: invalid data format")
						os.Exit(1)
					}
				}
				colony.graph[links[1]] = append(colony.graph[links[1]], links[0])

			} else {
				fmt.Println("ERROR: invalid data format")
				os.Exit(1)
			}
		}
		
	}
	if value,ok:=colony.graph[colony.start];ok{
		if len(value)== 0{
			fmt.Println("ERROR: invalid data format")
				os.Exit(1)
		}

	}
	if value,ok:=colony.graph[colony.end];ok{
		if len(value)== 0{
			fmt.Println("ERROR: invalid data format")
				os.Exit(1)
		}

	}

	if colony.start =="" || colony.end == ""{
		fmt.Println("ERROR: invalid aadfdfdfdfdffadata format")
				os.Exit(1)
	}

	for room1 := range checker.room {
		for room2 := range checker.room {
			if room1 != room2 {
				if checker.room[room1][0] == checker.room[room2][0] && checker.room[room1][1] == checker.room[room2][1]  {
					fmt.Println("ERROR: invalid iidata format")
				os.Exit(1)
				}
			}
		}
	}
	FindAllPaths()
}
