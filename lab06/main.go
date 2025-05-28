package main

import "fmt"

func main() {

	dane, err := loadDataFromCSV("nflx.csv")

	if err != nil {
		fmt.Println("Błąd:", err)
		return
	}

	fmt.Println(dane[0])

	wynik := EMA(dane)

	fmt.Println(wynik)
}