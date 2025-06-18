package app

import (
	"encoding/csv"
	"strconv"
	"strings"
	"fmt"
	"os"
)

func LoadCityCoordinates(filePath string) (map[string]City, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("nie można otworzyć pliku %s: %w", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Read()

	cities := make(map[string]City)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		name := strings.ToLower(record[0])
		lat, _ := strconv.ParseFloat(record[2], 64)
		lng, _ := strconv.ParseFloat(record[3], 64)

		cities[name] = City{
			Name:      record[0],
			Latitude:  lat,
			Longitude: lng,
		}
	}
	return cities, nil
}

func FindCityCoordinates(cityName string, cities map[string]City) (City, error) {
	city, exists := cities[strings.ToLower(cityName)]
	if !exists {
		return City{}, fmt.Errorf("miasto %s nie znalezione", cityName)
	}
	return city, nil
}