package loggerific

import (
	"log"
	"regexp"
	"time"
)

var djangoTimestampRegexp = regexp.MustCompile(`\[(\d\d/\w\w\w/\d\d\d\d\s\d\d:\d\d:\d\d).*\]`)

const djangoTimeFormat = "02/Jan/2006 15:04:05"

type Parser struct {
	TS *Timeseries
}

func NewParser() *Parser {
	return &Parser{
		TS: NewTimeseries(),
	}
}

func (p *Parser) Parse(input []byte) {
	matches := djangoTimestampRegexp.FindSubmatch(input)
	if len(matches) == 2 {
		t, err := time.Parse(djangoTimeFormat, string(matches[1]))
		if err != nil {
			log.Println("Error parsing time", string(matches[1]))
			return
		}

		p.TS.Add(t)
	}
}

func (p *Parser) Observations() []*TimeseriesObservation {
	return p.TS.Observations()
}
