package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func SplitAndTrim(s string, sep string) []string {
	var parts []string
	for _, p := range strings.Split(s, sep) {
		parts = append(parts, strings.TrimSpace(p))
	}
	return parts
}

func Catch(err error) {
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}
}

func ReplaceAndGetInt(s string, label string) int {
	val, err := strconv.Atoi(strings.TrimSpace(strings.Replace(s, label, "", -1)))
	Catch(err)

	return val
}

func IsNumber(n byte) bool {
	return n >= '0' && n <= '9'
}

func ConvertToNumber(s string) int {
	i, err := strconv.Atoi(s)
	Catch(err)

	return i
}
