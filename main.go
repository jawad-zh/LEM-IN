package main

import (
	"bufio"
	"fmt"
	function "functions/functions"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("[USAGE]: go run main.go <DataFileName.txt>")
		return
	}
	var Data []string
	var NewData [][]string
	FileName := os.Args[1]
	file, err := os.Open(FileName)
	defer file.Close()
	if err != nil {
		fmt.Print("Error:", err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		
		Data = function.DataRange(line)
		NewData = append(NewData, Data)

	}

	fmt.Println(NewData)
	// fmt.Print(len(Data))
}
