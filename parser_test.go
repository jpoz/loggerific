package loggerific

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func BuildFileScanner(path string) (*bufio.Scanner, func() error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file), file.Close
}

func TestParserWithBasicLog(t *testing.T) {
	p := &Parser{}

	scanner, done := BuildFileScanner("test/fixtures/basic.txt")
	defer done()

	for scanner.Scan() {
		p.Parse(scanner.Text())
	}
}
