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

var num_map = map[string]string{
	"zero":  "0",
	"0":     "0",
	"one":   "1",
	"1":     "1",
	"two":   "2",
	"2":     "2",
	"three": "3",
	"3":     "3",
	"four":  "4",
	"4":     "4",
	"five":  "5",
	"5":     "5",
	"six":   "6",
	"6":     "6",
	"seven": "7",
	"7":     "7",
	"eight": "8",
	"8":     "8",
	"nine":  "9",
	"9":     "9",
}

func find_numbers(line string) (string, string, error) {
	if line == "" || line == "\n" {
		return "", "", errors.New("Empty line")
	}

	found_numbers := []string{}
	window_size := 5
	for i := range line {
		window_end := i + window_size
		if window_end > len(line)-1 {
			window_end = len(line) - 1
		}
		sub_str := line[i:window_end]

		for key, value := range num_map {
			if strings.HasPrefix(sub_str, key) {
				found_numbers = append(found_numbers, value)
			}
		}
	}

	if len(found_numbers) == 0 {
		return "", "", errors.New("No numbers found")
	}

	first := found_numbers[0]
	last := found_numbers[len(found_numbers)-1]

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
