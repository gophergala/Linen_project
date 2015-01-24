// Simple go package to read a word at a time from a file
// Used to following example to learn from: http://golang.org/pkg/bufio/#example_Scanner_words
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
/*
	// An artificial input source.
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
*/

	
	// A file input source.
	fin, err := os.Open("res/dict.txt")
	r := bufio.NewReader(fin)

	strFromReader, err := r.ReadString('\n')
	fmt.Printf("Line: %v (error %v)\n", strFromReader, err)
	
	// convert that line to a string
	scanner := bufio.NewScanner(strings.NewReader(strFromReader))

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)

}
