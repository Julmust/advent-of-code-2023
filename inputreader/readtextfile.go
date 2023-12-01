package readinput

import (
	"bufio"
	"fmt"
	"os"
)

func ReadText(path string) []string {
	var data []string
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		val := fileScanner.Text()

		if err != nil {
			fmt.Println(err)
		}

		data = append(data, val)
	}

	return data
}
