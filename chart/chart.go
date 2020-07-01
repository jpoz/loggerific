package chart

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// DataPoint //
type DataPoint struct {
	X, Y int
}

// BarChart //
type BarChart struct {
	*tview.Box
	dataPoints []DataPoint
	padding    []int
}

// NewBarChart //
func NewBarChart() *BarChart {
	return &BarChart{
		Box:        tview.NewBox().SetBackgroundColor(tcell.ColorDefault),
		dataPoints: []DataPoint{},
		padding:    []int{2, 2, 3, 5},
	}
}

// AddDataPoint will add a data point
func (c *BarChart) AddDataPoint(x, y int) {
	c.dataPoints = append(c.dataPoints, DataPoint{x, y})
}

// Draw where all the magic happens
func (c *BarChart) Draw(screen tcell.Screen) {
	c.Box.Draw(screen)
	c.drawAxis(screen)
	c.drawData(screen)
}

func (c *BarChart) drawData(screen tcell.Screen) {
	color := tcell.ColorIndianRed
	fieldStyle := tcell.StyleDefault.Background(color)

	_, y, width, height := c.GetInnerRect()
	chartBottom := y + height - c.padding[2] - 1
	paddingRight := c.padding[1]
	// paddingTop := c.padding[0]
	// paddingBottom := c.padding[2]

	// min, max := c.minMaxHeight()
	// dataRange := max - min
	// boxRange := height - paddingTop - paddingBottom

	for idx, dp := range c.dataPoints {
		for i := 0; i < dp.Y; i++ {
			screen.SetContent(
				width-paddingRight-idx-1,
				chartBottom-i,
				' ',
				nil,
				fieldStyle,
			)
		}
	}
}

func (c *BarChart) drawAxis(screen tcell.Screen) {
	x, y, width, height := c.GetInnerRect()

	color := tcell.ColorAntiqueWhite

	paddingTop := c.padding[0]
	paddingRight := c.padding[1]
	paddingBottom := c.padding[2]
	paddingLeft := c.padding[3]

	for i := paddingTop; i < height-paddingBottom; i++ {
		tview.Print(screen, "│", x+paddingLeft, (y + i), width, tview.AlignLeft, color)
	}

	for i := paddingLeft; i < width-paddingRight; i++ {
		tview.Print(screen, "⎼", x+i, (y + height - paddingBottom), width, tview.AlignLeft, color)
	}

	tview.Print(screen, "┼", x+paddingLeft, (y + height - paddingBottom), width, tview.AlignLeft, color)
}

func (c *BarChart) minMaxHeight() (int, int) {
	var max int = c.dataPoints[0].Y
	var min int = c.dataPoints[0].Y
	for _, dp := range c.dataPoints {
		if max < dp.Y {
			max = dp.Y
		}
		if min > dp.Y {
			min = dp.Y
		}
	}

	return min, max
}
