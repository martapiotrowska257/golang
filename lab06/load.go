package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"lab06/functions"
	"lab06/structures"
	"time"
)

func loadDataFromCSV(filePath string) ([]structures.Data, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("nie można otworzyć pliku %s: %w", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("nie można odczytać danych CSV: %w", err)
	}

	var dataEntries []structures.Data

	if len(rawCSVData) < 2 {
		return nil, fmt.Errorf("plik CSV jest pusty lub zawiera tylko nagłówek")
	}

	for i, record := range rawCSVData {
		if i == 0 {
			continue
		}

		if len(record) < 6 {
			fmt.Printf("Pominięto wiersz %d: nieprawidłowa liczba kolumn (%d)", i+1, len(record))
			continue
		}

		date, err := time.Parse("01/02/2006", record[0])
		if err != nil {
			fmt.Printf("Pominięto wiersz %d: błąd parsowania daty '%s': %v", i+1, record[0], err)
			continue
		}

		last, err := functions.ParseFloat64(record[1])
		if err != nil {
			fmt.Printf("Pominięto wiersz %d: błąd parsowania 'last' ('%s'): %v", i+1, record[1], err)
			continue
		}

		volume, err := functions.ParseInt32(record[2])
		if err != nil {
			fmt.Printf("Pominięto wiersz %d: błąd parsowania 'volume' ('%s'): %v", i+1, record[2], err)
			continue
		}

		open, err := functions.ParseFloat64(record[3])
		if err != nil {
			fmt.Printf("Pominięto wiersz %d: błąd parsowania 'open' ('%s'): %v", i+1, record[3], err)
			continue
		}

		high, err := functions.ParseFloat64(record[4])
		if err != nil {
			fmt.Printf("Pominięto wiersz %d: błąd parsowania 'high' ('%s'): %v", i+1, record[4], err)
			continue
		}

		low, err := functions.ParseFloat64(record[5])
		if err != nil {
			fmt.Printf("Pominięto wiersz %d: błąd parsowania 'low' ('%s'): %v", i+1, record[5], err)
			continue
		}

		dataEntry := structures.Data{
			Date:   date,
			Last:   last,
			Volume: volume,
			Open:   open,
			High:   high,
			Low:    low,
		}
		dataEntries = append(dataEntries, dataEntry)
	}

	if len(dataEntries) == 0 && len(rawCSVData) > 1 {
		return nil, fmt.Errorf("nie udało się sparsować żadnych poprawnych wierszy danych z pliku CSV")
	}
	if len(dataEntries) == 0 && len(rawCSVData) <= 1 {
		return nil, fmt.Errorf("plik CSV nie zawierał danych do przetworzenia (poza ewentualnym nagłówkiem)")
	}

	return dataEntries, nil
}