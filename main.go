package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func normalize(s string, trim bool) string {
	if trim {
		s = strings.TrimSpace(s)
	}
	return strings.ToLower(s)
}

func main() {
	compareFile := flag.String("f", "", "comparison file (required)")
	outputFile := flag.String("o", "", "save output to file")
	quiet := flag.Bool("q", false, "silent mode, only works with -o")
	noTrim := flag.Bool("no-trim", false, "disable trimming whitespace before comparison")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: cat file.txt | dsame -f <compare_file> [options]\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *compareFile == "" {
		fmt.Fprintln(os.Stderr, "error: -f <compare_file> is required")
		flag.Usage()
		os.Exit(1)
	}

	// Load compare file
	cf, err := os.Open(*compareFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening compare file: %v\n", err)
		os.Exit(1)
	}
	defer cf.Close()

	seen := make(map[string]struct{})
	scanner := bufio.NewScanner(cf)
	for scanner.Scan() {
		key := normalize(scanner.Text(), !*noTrim)
		if key != "" {
			seen[key] = struct{}{}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading compare file: %v\n", err)
		os.Exit(1)
	}

	// Read stdin
	var inputLines []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		inputLines = append(inputLines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading stdin: %v\n", err)
		os.Exit(1)
	}

	// Filter lines
	var outputLines []string
	for _, line := range inputLines {
		key := normalize(line, !*noTrim)
		if key == "" {
			continue
		}
		if _, exists := seen[key]; !exists {
			outputLines = append(outputLines, line)
		}
	}

	// Save to file if -o
	if *outputFile != "" {
		of, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer of.Close()
		for _, line := range outputLines {
			fmt.Fprintln(of, line)
		}
	}

	// Print to stdout if not quiet
	if !*quiet {
		for _, line := range outputLines {
			fmt.Println(line)
		}
	}
}
