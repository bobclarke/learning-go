package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	dirPath := os.Args[1]
	dir, _ := os.Open(dirPath)
	files, _ := dir.Readdir(0)

	for fileIndex := range files {
		file := files[fileIndex]
		fileName := file.Name()
		santander123_processor(dirPath, fileName)
	}
}

func santander123_processor(dirPath string, fileName string) {

	filePath := fmt.Sprintf("%s/%s", dirPath, fileName)

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var properLines []string
	var properLine string

	for scanner.Scan() {
		l := scanner.Text()
		if strings.Contains(l, "Date:") || strings.Contains(l, "Description:") || strings.Contains(l, "Amount:") || strings.Contains(l, "Balance:") {
			value := strings.Split(l, ":")[1]
			re, _ := regexp.Compile(`[^\w]`)
			value = re.ReplaceAllString(value, "-")
			value = strings.Trim(value, "-")
			properLine += value + ", "
		} else {
			properLines = append(properLines, properLine)
			properLine = ""
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, line := range properLines {
		fmt.Println(line)
	}
}
