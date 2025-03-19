package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// struktury

type Location struct {
	Lat float64
	Long float64
}

type CovidData struct {
	Code int
	State string
	City string
	Date string
	Total_Death int
	Total_Confirmed int
	Location Location

}

// funkcja main do odczytania pliku csv z danymi

func main() { 
	file, err := os.Open("dane.csv") 
	if err != nil { 
		log.Fatal("Error while reading the file", err) 
	} 
	defer file.Close() 

	reader := csv.NewReader(file)
	reader.Comma = ';'
	records, err := reader.ReadAll() 

	if err != nil { 
		log.Fatal("Error reading records:", err) 
	} 

	var data []CovidData	// slice struktur
	
	for i, r := range records {
		if i == 0 {		// pomijamy pierwszy wiersz (nagłówek z nazwami kolumn) z pliku
			continue
		}
		code, _ := strconv.Atoi(r[0])
		total_death, _ := strconv.Atoi(r[4])
		total_confirmed, _ := strconv.Atoi(r[5])
		

		// obsługa lokalizacji (Lat, Long w jednej kolumnie)
		var lat, long float64
		locationParts := strings.Split(r[6], ", ")
		if len(locationParts) == 2 {
			lat, _ = strconv.ParseFloat(locationParts[0], 64)
			long, _ = strconv.ParseFloat(locationParts[1], 64)
		}

		data = append(data, CovidData{code, r[1], r[2], r[3], total_death, total_confirmed, Location{lat, long}})
	}

	// sortowanie slice na dwa różne sposoby

	// 1. sortowanie po liczbie potwierdzonych przypadków rosnąco
	sort.Slice(data, func(i, j int) bool { return data[i].Total_Confirmed < data[j].Total_Confirmed })
	fmt.Println("Top 5 regionów z najmniejszą liczbą potwierdzonych przypadków choroby:")
	for _, d := range data[:5] {
		fmt.Println(d)
	}

	// 2. sortowanie po liczbie zgonów malejąco
	sort.Slice(data, func(i, j int) bool { return data[i].Total_Death > data[j].Total_Death })
	fmt.Println("Top 5 regionów z największa liczbą śmierci wskutek choroby:")
	for _, d := range data[:5] {
		fmt.Println(d)
	}

	// przedstawienie wybranej statystyki
	casesByDate := make(map[string]int)
	for _, d := range data {
		casesByDate[d.Date] += d.Total_Confirmed
	}

	maxCases := 0
	maxDate := ""
	
	for date, cases := range casesByDate {
		if cases > maxCases {
			maxCases = cases
			maxDate = date
		}
	}
	fmt.Printf("\nDzień z największą liczbą potwierdzonych przypadków choroby: %s (%d cases)\n", maxDate, maxCases)
}
