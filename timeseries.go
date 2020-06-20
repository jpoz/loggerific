package loggerific

import (
	"time"
)

type TimeseriesObservation struct {
	At    time.Time
	Count int64
}

type Timeseries struct {
	observations []*TimeseriesObservation
}

func NewTimeseries() *Timeseries {
	return &Timeseries{
		observations: []*TimeseriesObservation{},
	}
}

func (ts *Timeseries) Add(at time.Time) {
	obsLen := len(ts.observations)
	if obsLen == 0 {
		ts.observations = append(ts.observations, ts.newObservation(at))
		return
	}

	lastObservation := ts.observations[obsLen-1]
	if lastObservation.At.Before(at) {
		ts.observations = append(ts.observations, ts.newObservation(at))
		return
	}
	// TODO could use a way better algo here
	for idx, obs := range ts.observations {
		if obs.At.After(at) {
			ts.observations = append(ts.observations, nil)

			copy(ts.observations[(idx+1):], ts.observations[idx:])
			ts.observations[idx] = ts.newObservation(at)

			return
		} else if obs.At.Equal(at) {
			obs.Count += 1
		}
	}
}

func (ts *Timeseries) Observations() []*TimeseriesObservation {
	return ts.observations
}

func (ts *Timeseries) newObservation(at time.Time) *TimeseriesObservation {
	return &TimeseriesObservation{
		At:    at,
		Count: 1,
	}
}
