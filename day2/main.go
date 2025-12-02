package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
//go:embed input_test.txt
var f embed.FS

func main() {
	res1, err := part1()
	if err != nil {
		panic(err)
	}
	fmt.Println(res1)

	res2, err := part2()
	if err != nil {
		panic(err)
	}
	fmt.Println(res2)
}

func part1() (int, error) {
	file, err := f.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(scanComma)

	res := 0
	for scanner.Scan() {
		b := scanner.Bytes()
		if b[0] == '0' { // Not sure if we are supposed to skip the whole range.
			continue
		}

		i := bytes.IndexByte(b, '-')
		from, err := strconv.Atoi(string(b[:i]))
		if err != nil {
			return 0, err
		}
		to, err := strconv.Atoi(string(b[i+1:]))
		if err != nil {
			return 0, err
		}
		for n := from; n <= to; n++ {
			str := strconv.Itoa(n)
			if str[0:len(str)/2] == str[len(str)/2:] {
				res += n
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return res, nil
}

func part2() (int, error) {
	file, err := f.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(scanComma)

	res := 0
	for scanner.Scan() {
		b := scanner.Bytes()
		if b[0] == '0' { // Not sure if we are supposed to skip the whole range.
			continue
		}

		i := bytes.IndexByte(b, '-')
		from, err := strconv.Atoi(string(b[:i]))
		if err != nil {
			return 0, err
		}
		to, err := strconv.Atoi(string(b[i+1:]))
		if err != nil {
			return 0, err
		}
		for n := from; n <= to; n++ {
			str := strconv.Itoa(n)
			for i := 1; i < len(str)/2+1; i++ {
				if strings.Repeat(str[0:i], len(str)/i) == str {
					res += n
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return res, nil
}

func scanComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	i := bytes.IndexByte(data, ',')
	if i > 0 {
		return i + 1, data[:i], nil
	}

	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}

	return 0, nil, nil
}
