package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
)

func parseSignaturesFile(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	sigs := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for lnum := 1; scanner.Scan(); lnum++ {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			return nil, fmt.Errorf("Error: Bad record at line: %d", lnum)
		}

		sigs[fields[1]] = fields[0]
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}

	return sigs, nil
}

func fileMD5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}

	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

type Result struct {
	path  string
	err   error
	match bool
}

func md5Worker(filename string, hash string, out chan *Result) {
	r := &Result{path: filename}

	s, err := fileMD5(filename)
	if err != nil {
		r.err = err
		out <- r
		return
	}

	r.match = (s == hash)
	out <- r
}

func main() {
	sigs, err := parseSignaturesFile("md5sum.txt")
	if err != nil {
		fmt.Println(err)
	}

	out := make(chan *Result)
	for file_name, hash := range sigs {
		go md5Worker(file_name, hash, out)
	}

	ok := true
	for range sigs {
		r := <-out
		switch {
		case r.err != nil:
			fmt.Printf("%s: error: %s\n", r.path, r.err)
			ok = false
		case !r.match:
			fmt.Printf("%s: Signature mismatch\n", r.path)
			ok = false
		}
	}

	if !ok {
		os.Exit(11)
	}
}
