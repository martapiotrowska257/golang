package app

import (
	"fmt"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func GenerateWeatherPlot(data []WeatherData, cityName, filename string) error {
	if len(data) == 0 {
		return fmt.Errorf("brak danych do wygenerowania wykresu")
	}

	p := plot.New()
	p.Title.Text = "Pogoda dla miasta: " + cityName 
	p.X.Label.Text = "Czas"
	p.Y.Label.Text = "Wartość"
	p.X.Tick.Marker = plot.TimeTicks{}

	tempPoints := make(plotter.XYs, len(data))
	windPoints := make(plotter.XYs, len(data))

	for i, row := range data {
		t, err := time.Parse("2006-01-02T15:04", row.Timestamp)
		if err != nil {
			t, err = time.Parse(time.RFC3339, row.Timestamp)
			if err != nil {
				return fmt.Errorf("nie udało się sparsować timestampu '%s': %w", row.Timestamp, err)
			}
		}

		tempPoints[i].X = float64(t.Unix())
		windPoints[i].X = float64(t.Unix())
		tempPoints[i].Y = row.Temperature
		windPoints[i].Y = row.WindSpeed
	}

	err := plotutil.AddLinePoints(p,
		"Temperatura (°C)", tempPoints,
		"Prędkość wiatru (km/h)", windPoints,
	)
	if err != nil {
		return err
	}

	if err := p.Save(12*vg.Inch, 6*vg.Inch, filename); err != nil {
		return err
	}
	return nil
}