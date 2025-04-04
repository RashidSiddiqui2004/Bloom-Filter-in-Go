package main

import (
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func GetPlot(plotTitle string, xLabel string, yLabel string, xlist []int, ylist []float64, directory string, imageSaveLocation string) {
	if len(xlist) != len(ylist) {
		log.Fatalf("xlist and ylist must be the same length")
	}

	// Create a new plot
	p := plot.New()
	p.Title.Text = plotTitle
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel

	// Convert data to plotter.XYs
	points := make(plotter.XYs, len(xlist))
	for i := 0; i < len(xlist); i++ {
		points[i].X = float64(xlist[i])
		points[i].Y = float64(ylist[i])
	}

	line, err := plotter.NewLine(points)
	if err != nil {
		log.Fatalf("Error creating line plot: %v", err)
	}
	line.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red color

	p.Add(line)

	// Save the plot as a PNG file
	if err := p.Save(6*vg.Inch, 4*vg.Inch, directory+"/"+imageSaveLocation+".png"); err != nil {
		log.Fatalf("Error saving plot: %v", err)
	}

	log.Println("Plot saved as", imageSaveLocation+".png")
}
