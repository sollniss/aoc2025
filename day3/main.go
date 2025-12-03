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

	res := 0
	for scanner.Scan() {
		b := scanner.Bytes()
		max10, max1 := -1, -1
		for i, ch := range b {
			n := int(ch - '0')
			if n > max10 && i+1 < len(b) {
				max10 = n
				max1 = -1
			} else if n > max1 {
				max1 = n
			}
		}
		res += max10*10 + max1
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

	res := 0
	for scanner.Scan() {
		b := scanner.Bytes()
		var max [12]int
		unset(&max, 0)
		for i, ch := range b {
			n := int(ch - '0')
			for j := range len(max) {
				if n > max[j] && i <= len(b)-(len(max)-j) {
					max[j] = n
					unset(&max, j+1)
					break
				}
			}
		}
		r := 0
		for _, n := range max {
			r = r*10 + n
		}
		res += r
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return res, nil
}

func unset(b *[12]int, i int) {
	for j := i; j < len(b); j++ {
		b[j] = -1
	}
}
