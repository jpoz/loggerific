package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/guptarohit/asciigraph"
	"github.com/rivo/tview"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func main() {
	app := tview.NewApplication()

	flex := tview.NewFlex()
	flex.Box.SetBackgroundColor(tcell.ColorDefault)
	flex.SetDirection(tview.FlexRow)

	chartView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	chartView.SetBackgroundColor(tcell.ColorDefault)
	flex.AddItem(chartView, 0, 1, false)

	logView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	logView.SetBackgroundColor(tcell.ColorDefault)
	flex.AddItem(logView, 0, 3, false)

	data := []float64{3, 4, 9, 6, 2, 4, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data, asciigraph.Height(100))

	io.WriteString(chartView, graph)

	go injestStdin(logView)

	if err := app.SetRoot(flex, true).SetFocus(logView).Run(); err != nil {
		panic(err)
	}
}

func injestStdin(logView *tview.TextView) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Fprintf(logView, "Input %s\n", scanner.Text())
	}

	if scanner.Err() != nil {
		// handle error.
	}
}
