package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	input, readErr := os.Open("./1/input")
	if readErr != nil {
		log.Fatalf("Failed to open input file %s", readErr)
		os.Exit(1)
	}

	count, err := innerMain(input)
	if err != nil {
		log.Fatalf("Error: %s", err)
		os.Exit(1)
	}

	log.Printf("There were %d depth increases", count)

	os.Exit(0)
}

func innerMain(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	depthSummaries := make([]*DepthSummary, 0)
	for scanner.Scan() {
		text := scanner.Text()
		depth, parseErr := strconv.ParseInt(text, 10, 32)
		if parseErr != nil {
			return 0, fmt.Errorf("error parsing line to int %w", parseErr)
		}

		depthSummary := newDepthSummary()
		depthSummaries = append(depthSummaries, depthSummary)

		for _, depthsum := range depthSummaries {
			if depthsum.Count() < 3 {
				depthsum.Append(depth)
			}
		}
	}

	increases := 0
	var previousSum int64
	for _, depthsum := range depthSummaries {
		log.Printf("depth summary %+v", depthsum)
		if previousSum != 0 && depthsum.Sum() < previousSum && depthsum.Count() == 3 {
			increases++
		}

		previousSum = depthsum.Sum()
	}

	return increases, nil
}

type DepthSummary struct {
	depths []int64
}

func newDepthSummary() *DepthSummary {
	return &DepthSummary{depths: make([]int64, 0)}
}

func (d *DepthSummary) Sum() int64 {
	var sum int64
	for _, depth := range d.depths {
		sum += depth
	}

	return sum
}

func (d *DepthSummary) Append(i int64) {
	d.depths = append(d.depths, i)
}

func (d *DepthSummary) Count() int {
	return len(d.depths)
}
