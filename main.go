package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	linesFlag := flag.Bool("l", false, "Count Lines")
	bytesFlag := flag.Bool("b", false, "Count Bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *linesFlag, *bytesFlag))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
