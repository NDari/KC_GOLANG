package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	file, err := openit("somefile")
	handle(err)
	lines, err := readit(file)
	handle(err)
	fmt.Println(lines)
}

func openit(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return file, errors.Wrap(err, "Failed to open the file")
	}
	return file, nil
}

func readit(file *os.File) ([]string, error) {
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		return nil, errors.Wrap(err, "Encountered error while scanning file")
	}
	return lines, nil
}

func handle(err error) {
	if err != nil {
		fmt.Printf("%+v", err)
		fmt.Println()
		panic("panicing!!")
	}
}
