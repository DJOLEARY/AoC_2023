package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	content := read_file(filename)

	numbers := []int{}

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		first_num, last_num, err := find_numbers(line)

		if err != nil {
			continue
		}

		num_str := first_num + last_num
		num := strToInt(num_str)
		numbers = append(numbers, num)
	}

	acc := accumulate(numbers)

	println(acc)
}

func read_file(filename string) string {
	content, err := os.ReadFile(filename)

	check(err)

	return string(content)
}

func find_numbers(line string) (string, string, error) {
	if line == "" || line == "\n" {
		return "", "", errors.New("Empty line")
	}

	first := ""
	last := ""

	chars := strings.Split(line, "")
	for _, char := range chars {
		is_numeric := strings.Contains("0123456789", char)
		if !is_numeric {
			continue
		}

		if first == "" {
			first = char
		}

		last = char
	}

	return first, last, nil
}

func accumulate(numbers []int) int {
	acc := 0

	for _, num := range numbers {
		acc += num
	}

	return acc
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func strToInt(str string) int {
	num, err := strconv.ParseInt(str, 0, 0)

	check(err)

	return int(num)
}
