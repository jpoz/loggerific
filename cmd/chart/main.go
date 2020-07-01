package main

import (
	"time"

	"github.com/jpoz/loggerific/chart"
	"github.com/rivo/tview"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func main() {
	app := tview.NewApplication()

	chart := chart.NewBarChart()

	chart.AddDataPoint(1, 3)
	chart.AddDataPoint(2, 4)
	chart.AddDataPoint(3, 5)
	chart.AddDataPoint(4, 6)
	chart.AddDataPoint(5, 7)
	chart.AddDataPoint(6, 8)

	go doEvery(700*time.Millisecond, func(t time.Time) {
		app.QueueUpdate(func() {
			chart.AddDataPoint(1, 3)
			chart.AddDataPoint(2, 4)
			chart.AddDataPoint(3, 5)
			chart.AddDataPoint(4, 6)
			chart.AddDataPoint(5, 7)
			chart.AddDataPoint(6, 8)
		})
	})

	if err := app.SetRoot(chart, true).Run(); err != nil {
		panic(err)
	}
}
