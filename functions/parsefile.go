package lemin

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParsFile(fileName string) {
	rooms := make(map[string][]string)
	var numberOfAnts int
	file, err := os.Open(fileName)
	var start, end string
	if err != nil {
		log.Fatalln(err)
	}
	var fileContenue []string
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
				log.Fatalln(err)
			}
			numberOfAnts = n
		} else {
			if line == "##start" && index != len(fileContenue)-1 {
				if start == "" {

					startRoom := strings.Fields(fileContenue[index+1])
					if len(startRoom) == 3 {
						_, err := strconv.Atoi(startRoom[1])
						if err != nil {
							log.Fatalln("x cordinnate is invalid nummber ")
						}
						_, err = strconv.Atoi(startRoom[2])
						if err != nil {
							log.Fatalln("y cordinnate is invalid nummber ")
						}
						start = startRoom[0]
					} else {
						log.Fatalln("invalid file format (start room)")
					}
				} else {
					log.Fatalln("start room already exist")
				}
			} else if line == "##end" && index != len(fileContenue)-1 {
				if end == "" {

					endRoom := strings.Fields(fileContenue[index+1])
					if len(endRoom) == 3 {
						_, err := strconv.Atoi(endRoom[1])
						if err != nil {
							log.Fatalln("x cordinnate is invalid nummber ")
						}
						_, err = strconv.Atoi(endRoom[2])
						if err != nil {
							log.Fatalln("y cordinnate is invalid nummber ")
						}
						end = endRoom[0]
					} else {
						log.Fatalln("invalid file format (start room)")
					}
				} else {
					log.Fatalln("end room already exist")
				}
			} else if len(strings.Fields(line)) == 3 {
				room := strings.Fields(line)
				_, err := strconv.Atoi(room[1])
				if err != nil {
					log.Fatalln("invalid room format")
				}
				_, err = strconv.Atoi(room[2])
				if err != nil {
					log.Fatalln("invalid room format")
				}
				if _, alreadyExist := rooms[room[0]]; alreadyExist {
					log.Fatalln("duplicated room")
				}
				rooms[room[0]] = nil
			} else if links := strings.Split(line, "-"); len(links) == 2 {
				if _, validRoom := rooms[links[0]]; !validRoom {
					log.Fatalln("invalid link")
				}
				if _, validRoom2 := rooms[links[1]]; !validRoom2 {
					log.Fatalln("invalid link")
				}
				for _, room := range rooms[links[0]] {
					if room == links[1] {
						log.Fatalln("duplicated link")
					}
				}
				rooms[links[0]] = append(rooms[links[0]], links[1])
				for _, room := range rooms[links[1]] {
					if room == links[0] {
						log.Fatalln("duplicated link")
					}
				}
				rooms[links[1]] = append(rooms[links[1]], links[0])

			} else {
				log.Fatalln("invalid file format")
			}
		}
	}
	FoundPath(rooms, start, end)
	fmt.Println("Ants:", numberOfAnts)
}
