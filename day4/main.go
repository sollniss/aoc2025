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

	adj := make([][]byte, 0, 139)
	m := make([][]byte, 0, 139)
	for scanner.Scan() {
		b := []byte(scanner.Text())
		m = append(m, b)
	}
	// init adj matrix
	for _, row := range m {
		adj = append(adj, make([]byte, len(row)))
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	res := 0
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if m[y][x] != '@' {
				continue
			}
			if x+1 < len(m[y]) && m[y][x+1] == '@' {
				adj[y][x]++
				adj[y][x+1]++
			}
			if y+1 < len(m) && x-1 >= 0 && m[y+1][x-1] == '@' {
				adj[y][x]++
				adj[y+1][x-1]++
			}
			if y+1 < len(m) && m[y+1][x] == '@' {
				adj[y][x]++
				adj[y+1][x]++
			}
			if y+1 < len(m) && x+1 < len(m[y+1]) && m[y+1][x+1] == '@' {
				adj[y][x]++
				adj[y+1][x+1]++
			}

			if adj[y][x] < 4 {
				res++
			}
		}
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

	m := make([][]byte, 0, 139)
	for scanner.Scan() {
		b := []byte(scanner.Text())
		m = append(m, b)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	res := 0
	x, y := 0, 0
rows:
	for y < len(m) {
		for x < len(m[y]) {
			if m[y][x] != '@' {
				x++
				continue
			}
			var curr byte
			if x-1 >= 0 && m[y][x-1] == '@' {
				curr++
			}
			if y-1 >= 0 && x-1 >= 0 && m[y-1][x-1] == '@' {
				curr++
			}
			if y-1 >= 0 && m[y-1][x] == '@' {
				curr++
			}
			if y-1 >= 0 && x+1 < len(m[y-1]) && m[y-1][x+1] == '@' {
				curr++
			}

			if x+1 < len(m[y]) && m[y][x+1] == '@' {
				curr++
			}
			if y+1 < len(m) && x-1 >= 0 && m[y+1][x-1] == '@' {
				curr++
			}
			if y+1 < len(m) && m[y+1][x] == '@' {
				curr++
			}
			if y+1 < len(m) && x+1 < len(m[y+1]) && m[y+1][x+1] == '@' {
				curr++
			}

			if curr < 4 {
				res++
				m[y][x] = '.'

				x = max(0, x-1)
				y = max(0, y-1)

				continue rows
			}
			x++
		}
		x = 0
		y++
	}

	return res, nil
}
