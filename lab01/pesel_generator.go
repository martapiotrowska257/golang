package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GeneratePESEL: geneuje numer PESEL
// Parametry:
// - birthDate: time.Time: reprezentacja daty urodzenia
// - płeć: znak "M" lub "K"
// Wyjscie:
//Tablica z cyframi numeru PESEL


func GenerujPESEL(birthDate time.Time, gender string) [11]int {

	// tablica zawierajaca kolejne cyfry numeru PESEL 
	var cyfryPESEL [11]int 

	// konwersja daty na dane skladowe 
	year := birthDate.Year()
	month := int(birthDate.Month())
	day := birthDate.Day()

	// --- Data urodzenia RRMMDD---

	// rok - RR
	cyfryPESEL[0] = (year % 100) / 10
	cyfryPESEL[1] = year % 10

	// miesiąc - MM
	switch {
	case year >= 1800 && year <= 1899:
		cyfryPESEL[2] = 8 + (month / 10)
	case year >= 1900 && year <= 1999:
		cyfryPESEL[2] = month / 10
	case year >= 2000 && year <= 2099:
		cyfryPESEL[2] = 2 + (month / 10)
	case year >= 2100 && year <= 2199:
		cyfryPESEL[2] = 4 + (month / 10)
	case year >= 2200 && year <= 2299:
		cyfryPESEL[2] = 6 + (month / 10)
	} 
	cyfryPESEL[3] = month % 10

	// dzień - DD
	cyfryPESEL[4] = day / 10
	cyfryPESEL[5] = day % 10

	// --- Liczba porządkowa - PPPP ---

	// losowy numer
	randomSerial := rand.Intn(900) + 100 // 3 cyfrowy losowy numer z zakresu 100-999
	cyfryPESEL[6] = randomSerial / 100
	cyfryPESEL[7] = (randomSerial % 100) / 10
	cyfryPESEL[8] = randomSerial % 10

	// płeć
	randomDigit := rand.Intn(5) * 2 // 0,2,4,6,8
	switch gender {
	case "M", "m":
		cyfryPESEL[9] = randomDigit + 1 // nieparzyste dla mężczyzn
	case "K", "k":
		cyfryPESEL[9] = randomDigit // parzyste dla kobiet
	default:
		cyfryPESEL[9] = -1
	}

	// --- Liczba kontrolna K ---
	var waga = [4]int{1, 3, 7, 9}
	counter := 0

	for i:= 0; i < len(cyfryPESEL) -1; i++ {
		wynik := (cyfryPESEL[i] * waga[i%4]) % 10
		counter += wynik
	}
	
	counter %= 10
	cyfryPESEL[10] = (10 - counter) % 10 // % 10, ponieważ jeśli counter = 0 to cyfryPESEL[10] = 10, a tak nie może być!

	return cyfryPESEL
}

// WeryfikujPESEL: weryfikuje poprawność numeru PESEL
// Parametry:
// - cyfryPESEL: Tablica z cyframi numeru PESEL
// Wyjscie:
//zmienna bool

func WeryfikujPESEL(cyfryPESEL [11]int) bool {
	
	var czyPESEL bool
	czyPESEL = true

	for i := 0; i < len(cyfryPESEL); i++ {
		if cyfryPESEL[i] < 0 || cyfryPESEL[i] > 9 {
			czyPESEL = false
			return czyPESEL
		}
	}
	
	var waga = [4]int{1, 3, 7, 9}
	counter := 0

	for i:= 0; i < len(cyfryPESEL) -1; i++ {
		wynik := (cyfryPESEL[i] * waga[i%4]) % 10
		counter += wynik
	}
	
	counter %= 10
	k := (10 - counter) % 10 // % 10, ponieważ jeśli counter = 0 to cyfryPESEL[10] = 10, a tak nie może być!	
	
	if (cyfryPESEL[10] != k) {
		czyPESEL = false
	}

	return czyPESEL
}

// Przykład użycia
func main() {
	// 
	birthDate := time.Date(1999, 8, 8, 0, 0, 0, 0, time.UTC)
	pesel := GenerujPESEL(birthDate, "K")
	
	fmt.Println("Wygenerowany PESEL:", pesel)

	fmt.Println("Czy numer PESEL jest poprawny:", WeryfikujPESEL(pesel))


}

