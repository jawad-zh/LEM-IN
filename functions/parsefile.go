package lemin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type room struct {
	roomName   string
	currentAnt int
	nextAnt    int
}

type Colony struct {
	antNumber  int
	start, end room
	graph      map[room][]room
}

var colony = Colony{
	graph: make(map[room][]room),
}

var fileContenue []string

func ParsFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		os.Exit(1)
	}
	fileBuffer := bufio.NewScanner(file)
	for fileBuffer.Scan() {
		line := fileBuffer.Text()
		if (line == "" || strings.HasPrefix(line, "#")) && !(line == "##start" || line == "##end") {
			continue
		}
		fileContenue = append(fileContenue, line)
	}
	for index, line := range fileContenue {
		if index == 0 {
			n, err := strconv.Atoi(line)
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
			if line == "##start" && index != len(fileContenue)-1 {
				if colony.start.roomName == "" {

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
						var start room
						start.roomName=startRoom[0]
						colony.start = start
					} else {
						fmt.Println("ERROR: invalid data format")
						os.Exit(1)
					}
				} else {
					fmt.Println("ERROR: invalid data format")
					os.Exit(1)
				}
			} else if line == "##end" && index != len(fileContenue)-1 {
				if colony.end.roomName == "" {

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
						var end room
						end.roomName=endRoom[0]
						colony.end = end
					} else {
						fmt.Println("ERROR: invalid data format")
						os.Exit(1)
					}
				} else {
					fmt.Println("ERROR: invalid data format")
					os.Exit(1)
				}
			} else if len(strings.Fields(line)) == 3 {
				roomInfos := strings.Fields(line)
				_, err := strconv.Atoi(roomInfos[1])
				if err != nil {
					fmt.Println("ERROR: invalid data format")
					os.Exit(1)
				}
				_, err = strconv.Atoi(roomInfos[2])
				if err != nil {
					fmt.Println("ERROR: invalid data format")
					os.Exit(1)
				}
				var newRoom room
				newRoom.roomName = roomInfos[0]
				if _, alreadyExist := colony.graph[newRoom]; alreadyExist {
					fmt.Println("ERROR: invalid data format")
					os.Exit(1)
				}
				colony.graph[newRoom] = nil
			} else if links := strings.Split(line, "-"); len(links) == 2 {
				var firstLinkChecker room
				firstLinkChecker.roomName = links[0]
				if _, validRoom := colony.graph[firstLinkChecker]; !validRoom {
					fmt.Println("ERROR: invalid data format")
					os.Exit(1)
				}
				var secondLinkChecker room
				secondLinkChecker.roomName = links[1]
				if _, validRoom2 := colony.graph[secondLinkChecker]; !validRoom2 {
					fmt.Println("ERROR: invalid data format")
					os.Exit(1)
				}
				for _, room := range colony.graph[firstLinkChecker] {
					if room == secondLinkChecker {
						fmt.Println("ERROR: invalid data format")
						os.Exit(1)
					}
				}
				colony.graph[firstLinkChecker] = append(colony.graph[firstLinkChecker], secondLinkChecker)
				for _, room := range colony.graph[secondLinkChecker] {
					if room == firstLinkChecker {
						fmt.Println("ERROR: invalid data format")
						os.Exit(1)
					}
				}
				colony.graph[secondLinkChecker] = append(colony.graph[secondLinkChecker],firstLinkChecker)

			} else {
				fmt.Println("ERROR: invalid data format")
				os.Exit(1)
			}
		}
	}
	FindAllPaths()
}
