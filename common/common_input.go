package common

import (
	"bufio"
	"os"
)

func ReadInputFile(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var buf = make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buf = append(buf, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return buf, nil
}
