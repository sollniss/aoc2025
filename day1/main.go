package main

import (
	"bufio"
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
	curr := 50
	zeroCount := 0
	for scanner.Scan() {
		b := scanner.Bytes()
		num := 0
		for i := 1; i < len(b); i++ {
			ch := b[i]
			ch -= '0'
			num = num*10 + int(ch)
		}

		switch b[0] {
		case 'R':
			curr += num
		case 'L':
			curr -= num
		}
		curr = mod(curr, 100)
		if curr == 0 {
			zeroCount++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return zeroCount, nil
}

func part2() (int, error) {
	file, err := f.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	curr := 50
	zeroCount := 0
	for scanner.Scan() {
		b := scanner.Bytes()
		num := 0
		for i := 1; i < len(b); i++ {
			ch := b[i]
			ch -= '0'
			num = num*10 + int(ch)
		}

		switch b[0] {
		case 'R':
			if curr+num >= 100 {
				zeroCount += (curr + num) / 100
			}
			curr += num
		case 'L':
			if curr > 0 && num >= curr {
				zeroCount++
			}
			if curr-num < 0 {
				zeroCount += -(curr - num) / 100
			}
			curr -= num
		}
		curr = mod(curr, 100)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return zeroCount, nil
}

func mod(a, b int) int {
	return (a%b + b) % b
}
