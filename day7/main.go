package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
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

	var beams map[int]bool
	res := 0
	skip := false
	for scanner.Scan() {
		if skip {
			skip = false
			continue
		}
		b := scanner.Bytes()
		maxLen := len(b)
		if len(beams) == 0 {
			beams = make(map[int]bool, maxLen)
			start := bytes.IndexByte(b, 'S')
			beams[start] = true
			skip = true
			continue
		}

		idx := 0
		for {
			i := bytes.IndexByte(b[idx:], '^')
			if i == -1 {
				break
			}
			absIdx := idx + i

			if beams[absIdx] {
				res++
			}
			beams[absIdx] = false
			if i > 0 {
				beams[absIdx-1] = true
			}
			if i < maxLen-1 {
				beams[absIdx+1] = true
			}

			idx = absIdx + 1
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

	var beams map[int]int
	res := 0
	skip := false
	for scanner.Scan() {
		if skip {
			skip = false
			continue
		}
		b := scanner.Bytes()
		maxLen := len(b)
		if len(beams) == 0 {
			beams = make(map[int]int, maxLen)
			start := bytes.IndexByte(b, 'S')
			beams[start] = 1
			skip = true
			continue
		}

		idx := 0
		for {
			i := bytes.IndexByte(b[idx:], '^')
			if i == -1 {
				break
			}
			absIdx := idx + i

			incoming := beams[absIdx]
			beams[absIdx] = 0
			if i > 0 {
				beams[absIdx-1] += incoming
			}
			if i < maxLen-1 {
				beams[absIdx+1] += incoming
			}

			idx = absIdx + 1
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	for _, b := range beams {
		res += b
	}

	return res, nil
}
