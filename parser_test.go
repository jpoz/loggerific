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
	p := NewParser()

	scanner, done := BuildFileScanner("test/fixtures/basic.txt")
	defer done()

	for scanner.Scan() {
		p.Parse(scanner.Bytes())
	}

	observations := p.Observations()

	actualLength := len(observations)
	if actualLength != 2 {
		t.Errorf("Observations length %d; want 2", actualLength)
	}

	firstObservation := observations[0]
	if firstObservation.Count != 10 {
		t.Errorf("firstObservation count %d; want 10", firstObservation.Count)
	}

	secondObservation := observations[1]
	if secondObservation.Count != 15 {
		t.Errorf("secondObservation count %d; want 10", secondObservation.Count)
	}
}
