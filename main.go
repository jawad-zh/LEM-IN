package main

import (
	"fmt"
	lemin "lemin/functions"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("[USAGE]: go run main.go <DataFileName.txt>")
		return
	}
	lemin.ParsFile(os.Args[1])
	
}
