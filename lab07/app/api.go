package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WeatherAPIResponse struct {
	Hourly struct {
		Timestamp        []string  `json:"time"`
		Temperature      []float64 `json:"temperature_2m"`
		ApparentTemp     []float64 `json:"apparent_temperature"`
		Humidity         []float64 `json:"relativehumidity_2m"`
		Precipitation    []float64 `json:"precipitation"`
		WindSpeed        []float64 `json:"windspeed_10m"`
		Pressure         []float64 `json:"surface_pressure"`
	} `json:"hourly"`
}

func FetchCurrentWeather(lat, lon float64) ([]WeatherData, error) {
	now := time.Now()

	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%.4f&longitude=%.4f&hourly=temperature_2m,apparent_temperature,relativehumidity_2m,precipitation,windspeed_10m,surface_pressure&start_date=%s&end_date=%s&timezone=UTC",
		lat, lon, now.Format("2006-01-02"), now.Format("2006-01-02"))

	return fetchWeather(url)
}

func FetchHistorical(lat, lon float64, date string) ([]WeatherData, error) {
	url := fmt.Sprintf(
		"https://archive-api.open-meteo.com/v1/archive?latitude=%.4f&longitude=%.4f&hourly=temperature_2m,apparent_temperature,relativehumidity_2m,precipitation,windspeed_10m,surface_pressure&start_date=%s&end_date=%s&timezone=UTC",
		lat, lon, date, date)

	return fetchWeather(url)
}

func FetchForecast(lat, lon float64, days int) ([]WeatherData, error) {
	if days < 1 {
		days = 1
	}
	if days > 16 {
		days = 16
	}
	start := time.Now()
	end := start.Add(time.Duration(days) * 24 * time.Hour)
	
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%.4f&longitude=%.4f&hourly=temperature_2m,apparent_temperature,relativehumidity_2m,precipitation,windspeed_10m,surface_pressure&start_date=%s&end_date=%s&timezone=UTC",
		lat, lon, start.Format("2006-01-02"), end.Format("2006-01-02"),
	)
	return fetchWeather(url)
}

func FetchCurrentWeatherFull(lat, lon float64) ([]WeatherData, error) {
	now := time.Now()
	yesterday := now.Add(-12 * time.Hour)
	tomorrow := now.Add(12 * time.Hour)

	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%.4f&longitude=%.4f&hourly=temperature_2m,apparent_temperature,relativehumidity_2m,precipitation,windspeed_10m,surface_pressure&start_date=%s&end_date=%s&timezone=UTC",
		lat, lon, yesterday.Format("2006-01-02"), tomorrow.Format("2006-01-02"))
	return fetchWeather(url)
}

func fetchWeather(url string) ([]WeatherData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error pobierania danych pogodowych: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error API pogodowego: %s (status: %d)", resp.Status, resp.StatusCode)
	}

	var apiResponse WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("error dekodowania odpowiedzi JSON: %w", err)
	}

	data := make([]WeatherData, len(apiResponse.Hourly.Timestamp))
	for i := range apiResponse.Hourly.Timestamp {
		temp := 0.0
		apparent := 0.0
		humidity := 0.0
		precipitation := 0.0
		windSpeed := 0.0
		pressure := 0.0

		if i < len(apiResponse.Hourly.Temperature) {
			temp = apiResponse.Hourly.Temperature[i]
		}
		if i < len(apiResponse.Hourly.ApparentTemp) {
			apparent = apiResponse.Hourly.ApparentTemp[i]
		}
		if i < len(apiResponse.Hourly.Humidity) {
			humidity = apiResponse.Hourly.Humidity[i]
		}
		if i < len(apiResponse.Hourly.Precipitation) {
			precipitation = apiResponse.Hourly.Precipitation[i]
		}
		if i < len(apiResponse.Hourly.WindSpeed) {
			windSpeed = apiResponse.Hourly.WindSpeed[i]
		}
		if i < len(apiResponse.Hourly.Pressure) {
			pressure = apiResponse.Hourly.Pressure[i]
		}

		data[i] = WeatherData{
			Timestamp:           apiResponse.Hourly.Timestamp[i],
			Temperature:         temp,
			ApparentTemperature: apparent,
			Humidity:            humidity,
			Precipitation:       precipitation,
			WindSpeed:           windSpeed,
			Pressure:            pressure,
		}
	}
	return data, nil
}