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

	addedObservation := observations[1]
	if addedObservation.At != obsTime {
		t.Error("Added time doesnt match")
	}
}
