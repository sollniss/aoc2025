package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"slices"
	"strconv"
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

	type Range struct {
		From uint64
		To   uint64
	}
	ranges := make([]Range, 0)

	for scanner.Scan() {
		b := scanner.Bytes()
		if len(b) == 0 {
			break
		}

		i := bytes.IndexByte(b, '-')
		from, err := strconv.ParseUint(string(b[:i]), 10, 64)
		if err != nil {
			return 0, err
		}
		to, err := strconv.ParseUint(string(b[i+1:]), 10, 64)
		if err != nil {
			return 0, err
		}
		ranges = append(ranges, Range{
			From: from,
			To:   to,
		})
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	ids := make([]uint64, 0)
	for scanner.Scan() {
		b := scanner.Bytes()
		id, err := strconv.ParseUint(string(b), 10, 64)
		if err != nil {
			return 0, err
		}
		ids = append(ids, id)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		if a.From < b.From {
			return -1
		} else {
			return 1
		}
	})

	res := 0
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.From && id <= r.To {
				res++
				break
			}
			if id < r.From {
				break
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

	type Range struct {
		From uint64
		To   uint64
	}
	ranges := make([]Range, 0)

	for scanner.Scan() {
		b := scanner.Bytes()
		if len(b) == 0 {
			break
		}

		i := bytes.IndexByte(b, '-')
		from, err := strconv.ParseUint(string(b[:i]), 10, 64)
		if err != nil {
			return 0, err
		}
		to, err := strconv.ParseUint(string(b[i+1:]), 10, 64)
		if err != nil {
			return 0, err
		}
		ranges = append(ranges, Range{
			From: from,
			To:   to,
		})
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	slices.SortFunc(ranges, func(a, b Range) int {
		if a.From < b.From {
			return -1
		} else {
			return 1
		}
	})

	res := 0
	var curr uint64
	for _, r := range ranges {
		if curr > r.To {
			continue
		}
		from := max(curr, r.From)
		if from <= r.To {
			res += int(r.To-from) + 1
		}

		curr = r.To + 1
	}

	return res, nil
}
