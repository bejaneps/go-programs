package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func printRunesRange(min, max int) {
	letters := make([]rune, max-min)

	for i := 0; i+min < max; i++ {
		letters[i] = rune(i + min)
	}

	for i, val := range letters {
		if n, err := fmt.Printf("%d: %q\n", i+min, val); n == 0 || err != nil {
			log.Fatalf("%v\n", err)
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: ./program min max")
	}

	min, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	max, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	if max-min <= 0 {
		log.Fatalf("max < min.\n")
	} else {
		printRunesRange(min, max)
	}
}
