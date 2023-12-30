package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type InputOptions struct {
	Path  string
	Split string
}

func GetInput(o InputOptions) ([]string, error) {
	file, err := os.Open(o.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		if o.Split != "" {
			t = strings.Split(t, o.Split)[1]
			t = strings.TrimSpace(t)
		}
		lines = append(lines, t)
	}

	return lines, scanner.Err()
}

func Catch(err error) {
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}
}
