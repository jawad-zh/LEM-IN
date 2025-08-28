package main

import (
	"fmt"
	"os"

	function "functions/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("[USAGE]: go run main.go <DataFileName.txt>")
		return
	}
	function.ReadFile(os.Args[1])
}
