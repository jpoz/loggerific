package loggerific

import (
	"testing"
	"time"
)

func TestTimeseries(t *testing.T) {
	ts := NewTimeseries()

	obsTime := time.Now()
	ts.Add(obsTime)
	ts.Add(obsTime)
	ts.Add(obsTime)

	observations := ts.Observations()

	addedObservation := observations[0]
	if addedObservation.At != obsTime {
		t.Error("Added time doesnt match")
	}
	if addedObservation.Count != 3 {
		t.Error("Count wasn't 3")
	}
}
