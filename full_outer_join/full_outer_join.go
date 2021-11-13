package full_outer_join

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func FullOuterJoin(f1Path, f2Path, resultPath string) {
	var lines1, lines2, result []string
	lines1 = getLines(f1Path)
	lines2 = getLines(f2Path)
	for _, line := range lines1 {
		if !contains(lines2, line){
			result = append(result, line)
		}
	}
	for _, line := range lines2 {
		if !contains(lines1, line){
			result = append(result, line)
		}
	}
	sort.Strings(result)
	writeToFile(resultPath, result)
}

func getLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func writeToFile(path string, content []string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	datawriter := bufio.NewWriter(file)
	defer datawriter.Flush()

	for i, data := range content {
		if i == len(content) - 1 {
			_, _ = datawriter.WriteString(data)
			break
		}
		_, _ = datawriter.WriteString(data + "\n")
	}
}
