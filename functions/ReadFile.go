package function

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(FileName string) {
	var Data []string
	file, err := os.Open(FileName)
	if err != nil {
		fmt.Print("Error:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		
		Data = append(Data, scanner.Text())
	}
	DataRange(Data)
}
