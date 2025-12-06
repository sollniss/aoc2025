package main

import (
	"bufio"
	"embed"
	"errors"
	"fmt"
	"io"
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

func part1() (uint64, error) {
	file, err := f.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	rows := make([][]uint64, 0)
	currNo := uint64(0)
	currRow := make([]uint64, 0)

	currCol := 0
	res := uint64(0)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return 0, err
		}
		switch b {
		case ' ':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if currNo > 0 {
				currNo *= 10
			}
			currNo += uint64(b - '0')
			// A space following a number indicates a new column.
			n, err := reader.Peek(1)
			if err != nil {
				return 0, err
			}
			if n[0] == ' ' {
				currRow = append(currRow, currNo)
				currNo = 0
			}
		case '\n':
			rows = append(rows, append(currRow, currNo))
			currNo = 0
			currRow = make([]uint64, 0, len(currRow))
		case '+':
			currRes := rows[0][currCol]
			for i := 1; i < len(rows); i++ {
				currRes += rows[i][currCol]
			}
			res += currRes
			currCol++
		case '*':
			currRes := rows[0][currCol]
			for i := 1; i < len(rows); i++ {
				currRes *= rows[i][currCol]
			}
			res += currRes
			currCol++
		}
	}
	return res, nil
}

func part2() (uint64, error) {
	file, err := f.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	data := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b := []byte(scanner.Text())
		data = append(data, b)
	}

	res := uint64(0)

	currNos := make([]uint64, 0)
	var op func(a uint64, b uint64) uint64

	x := 0
outer:
	for {
		curr := uint64(0)
		allSpaces := true
		for y := 0; y < len(data); y++ {
			if x >= len(data[y]) {
				break outer
			}

			switch data[y][x] {
			case ' ', '\n':
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				allSpaces = false
				curr = curr*10 + uint64(data[y][x]-'0')
			case '+':
				allSpaces = false
				op = func(a uint64, b uint64) uint64 {
					return a + b
				}
			case '*':
				allSpaces = false
				op = func(a uint64, b uint64) uint64 {
					return a * b
				}
			}
		}
		if allSpaces {
			r := currNos[0]
			for i := 1; i < len(currNos); i++ {
				r = op(r, currNos[i])
			}
			res += r
			currNos = currNos[:0]
		} else {
			currNos = append(currNos, curr)
		}
		x++
	}

	// Add final problem.
	r := currNos[0]
	for i := 1; i < len(currNos); i++ {
		r = op(r, currNos[i])
	}
	res += r

	return res, nil
}
