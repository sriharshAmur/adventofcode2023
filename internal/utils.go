package internal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFileContents(isTest bool) ([]string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current working directory: %s", err)
	}
	filename := "input.txt"
	if isTest {
		filename = "test.txt"
	}
	filePath := filepath.Join(cwd, filename)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s", err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		return nil, fmt.Errorf("error while scanning file: %s", err)
	}

	return lines, nil
}
