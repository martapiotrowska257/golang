package app

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadThresholdConfig(path string) (ThresholdConfig, error) {
	var cfg struct {
		ExtremeWeatherThresholds ThresholdConfig `yaml:"extreme_weather_thresholds"`
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return ThresholdConfig{}, err
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return ThresholdConfig{}, err
	}
	return cfg.ExtremeWeatherThresholds, nil
}

func AnalyzeWeather(data []WeatherData, cfg ThresholdConfig) []Alert {
	var alerts []Alert
	for _, row := range data {
		if row.Temperature >= cfg.HighTemp {
			alerts = append(alerts, Alert{
				Type:  "Upał",
				Date:  row.Timestamp,
				Value: row.Temperature,
			})
		}
		if row.Temperature <= cfg.LowTemp {
			alerts = append(alerts, Alert{
				Type:  "Mróz",
				Date:  row.Timestamp,
				Value: row.Temperature,
			})
		}
		if row.WindSpeed >= cfg.HighWind {
			alerts = append(alerts, Alert{
				Type:  "Silny wiatr",
				Date:  row.Timestamp,
				Value: row.WindSpeed,
			})
		}
		if row.Precipitation >= cfg.HeavyRain {
			alerts = append(alerts, Alert{
				Type:  "Intensywne opady",
				Date:  row.Timestamp,
				Value: row.Precipitation,
			})
		}
	}
	return alerts
}

func DisplayAlerts(alerts []Alert) {
	if len(alerts) == 0 {
		fmt.Println("Brak zagrożeń pogodowych.")
		return
	}
	fmt.Println("Wykryto zagrożenia pogodowe:")
	for _, alert := range alerts {
		fmt.Printf("- [%s] %s: %.2f\n", alert.Date, alert.Type, alert.Value)
	}
}