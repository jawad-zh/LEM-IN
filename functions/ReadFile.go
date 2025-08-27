package function

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(FileName string) {
	var Data string
	var split []string
	file, err := os.Open(FileName)
	defer file.Close()
	if err != nil {
		fmt.Print("Error:", err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		Data += line + "\n"
	}
	split = strings.Split(Data,"\n")
	DataRange(split)
	fmt.Println(split)
	fmt.Print(len(split))
}
