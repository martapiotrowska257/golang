package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"weatherTracker/app"
)

func main() {
	aktualnaCmd := flag.NewFlagSet("aktualna", flag.ExitOnError)
	prognozaCmd := flag.NewFlagSet("prognoza", flag.ExitOnError)
	historiaCmd := flag.NewFlagSet("historia", flag.ExitOnError)

	prognozaDni := prognozaCmd.Int("dni", 1, "Liczba dni prognozy")
	historiaData := historiaCmd.String("data", "", "Data historyczna (RRRR-MM-DD)")

	if len(os.Args) < 2 {
		fmt.Println("Użyj: pogoda <komenda> <miasto> [opcje]")
		fmt.Println(`Dostępne komendy:
  aktualna <miasto>
  prognoza <miasto> [--dni N]
  historia <miasto> --data RRRR-MM-DD`)
		return
	}

	command := os.Args[1]
	if command != "aktualna" && command != "prognoza" && command != "historia" {
		fmt.Println("Nieznana komenda:", command)
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Podaj nazwę miasta!")
		return
	}
	
	cityName := os.Args[2]
	citiesPath := filepath.Join("data", "worldcities.csv")
	cities, err := app.LoadCityCoordinates(citiesPath)
	if err != nil {
		fmt.Println("Błąd ładowania miast:", err)
		return
	}
	city, err := app.FindCityCoordinates(cityName, cities)
	if err != nil {
		fmt.Println("Błąd miasta:", err)
		return
	}

	cfgPath := filepath.Join("config.yaml")
	cfg, _ := app.LoadThresholdConfig(cfgPath)
	switch command {	case "aktualna":
		aktualnaCmd.Parse(os.Args[3:])
		
		currentData, err := app.FetchCurrentWeather(city.Latitude, city.Longitude)
		if err != nil {
			fmt.Println("Błąd pobierania pogody:", err)
			return
		}

		fullData, err := app.FetchCurrentWeatherFull(city.Latitude, city.Longitude)
		if err != nil {
			fmt.Println("Błąd pobierania pełnych danych pogody:", err)
			return
		}

		if len(currentData) > 0 {
			latestPoint := currentData[len(currentData)-1]
			app.DisplayWeatherTable([]app.WeatherData{latestPoint}, city.Name)
		}
		
		alerts := app.AnalyzeWeather(fullData, cfg)
		app.DisplayAlerts(alerts)
		plotFile := fmt.Sprintf("wykres_%s.png", city.Name)
		err = app.GenerateWeatherPlot(fullData, city.Name, plotFile)
		if err != nil {
			fmt.Println("Błąd generowania wykresu:", err)
		} else {
			fmt.Println("Wykres zapisany jako:", plotFile)
		}
	case "prognoza":
		prognozaCmd.Parse(os.Args[3:])
		dni := *prognozaDni
		if dni < 1 || dni > 16 {
			fmt.Println("Liczba dni prognozy musi być w zakresie 1-16")
			return
		}
		data, err := app.FetchForecast(city.Latitude, city.Longitude, dni)
		if err != nil {
			fmt.Println("Błąd pobierania prognozy:", err)
			return
		}

		app.DisplayWeatherTable(data, city.Name)
		alerts := app.AnalyzeWeather(data, cfg)
		app.DisplayAlerts(alerts)
		plotFile := fmt.Sprintf("wykres_prognoza_%s.png", city.Name)
		err = app.GenerateWeatherPlot(data, city.Name, plotFile)
		if err != nil {
			fmt.Println("Błąd generowania wykresu:", err)
		} else {
			fmt.Println("Wykres zapisany jako:", plotFile)
		}
	case "historia":
		historiaCmd.Parse(os.Args[3:])
		if *historiaData == "" {
			fmt.Println("Podaj datę za pomocą --data RRRR-MM-DD")
			return
		}
		data, err := app.FetchHistorical(city.Latitude, city.Longitude, *historiaData)
		if err != nil {
			fmt.Println("Błąd pobierania danych historycznych:", err)
			return
		}
		app.DisplayWeatherTable(data, city.Name)
		alerts := app.AnalyzeWeather(data, cfg)
		app.DisplayAlerts(alerts)
		plotFile := fmt.Sprintf("wykres_historia_%s.png", city.Name)
		err = app.GenerateWeatherPlot(data, city.Name, plotFile)
		if err != nil {
			fmt.Println("Błąd generowania wykresu:", err)
		} else {
			fmt.Println("Wykres zapisany jako:", plotFile)
		}
	}
}