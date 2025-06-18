package app

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func DisplayWeatherTable(data []WeatherData, cityname string) {
	if len(data) == 0 {
		println("Brak danych pogodowych dla miasta:", cityname)
		return
	}

	println("Pogoda dla miasta:", cityname)

	if len(data) == 1 {
		fmt.Printf("\n=== AKTUALNA POGODA ===\n")
		fmt.Printf("Czas: %s\n", data[0].Timestamp)
		fmt.Printf("Temperatura: %.1f°C\n", data[0].Temperature)
		fmt.Printf("Temp. odczuwalna: %.1f°C\n", data[0].ApparentTemperature)
		fmt.Printf("Wilgotność: %.0f%%\n", data[0].Humidity)
		fmt.Printf("Opady: %.1f mm\n", data[0].Precipitation)
		fmt.Printf("Prędkość wiatru: %.1f km/h\n", data[0].WindSpeed)
		fmt.Printf("Ciśnienie: %.1f hPa\n", data[0].Pressure)
		fmt.Printf("=======================\n\n")
		return
	}

	headers := []string{"Czas", "Temperatura (°C)", "Temp. odczuwalna (°C)", "Wilgotność (%)", "Opady (mm)", "Prędkość wiatru (km/h)", "Ciśnienie (hPa)"}
	var rows [][]string

	for _, row := range data {
		rows = append(rows, []string{
			row.Timestamp,
			formatFloat(row.Temperature),
			formatFloat(row.ApparentTemperature),
			formatFloat(row.Humidity),
			formatFloat(row.Precipitation),
			formatFloat(row.WindSpeed),
			formatFloat(row.Pressure),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(headers)
	table.Bulk(rows)
	table.Render()
}

func formatFloat(val float64) string {
	return fmt.Sprintf("%.2f", val)
}