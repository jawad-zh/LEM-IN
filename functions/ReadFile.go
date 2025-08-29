package function

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(FileName string) {
	var Data []string
	
	file, err := os.Open(FileName)
	defer file.Close()
	if err != nil {
		fmt.Print("Error:", err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		
		Data = append(Data, scanner.Text())
	}
	
	DataRange(Data)
	// fmt.Println(split)
	// fmt.Print(len(split))
}
