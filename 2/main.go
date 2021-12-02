package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, readErr := os.Open("./2/input")
	if readErr != nil {
		log.Fatalf("Failed to open input file %s", readErr)
		os.Exit(1)
	}

	answer, err := innerMain(input)
	if err != nil {
		log.Fatalf("Failed to get the answer %s", err)
	}

	log.Printf("Answer %d", answer)
}

func innerMain(input io.Reader) (int64, error) {
	sub := &Sub{depth: 0, position: 0}

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		pair := strings.Split(text, " ")
		log.Printf("pair %v", pair)
		if len(pair) < 2 {
			continue
		}
		command := pair[0]
		value, parseErr := strconv.ParseInt(pair[1], 10, 64)
		if parseErr != nil {
			return 0, fmt.Errorf("error parsing input %w", parseErr)
		}

		log.Printf("command %s", command)
		log.Printf("value %d", value)
		sub.in(command, value)
	}
	log.Printf("sub %+v", sub)
	return sub.depth * sub.position, nil
}

type Sub struct {
	depth    int64
	position int64
	aim      int64
}

func (s *Sub) in(command string, value int64) {
	switch command {
	case "forward":
		s.position += value
		s.depth += s.aim * value
	case "up":
		s.aim -= value
	case "down":
		s.aim += value
	}
}
