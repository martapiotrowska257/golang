package main

import (
	"fmt"
	"lab05/api"
	"lab05/ui"
	"os"
)

func main() {
	fmt.Println("Loading ZTM Gdańsk stops data...")
	client := api.NewZTMClient()	// tworzymy nowego klienta ZTM
	if client == nil {
		fmt.Println("Failed to create ZTM client.")
		os.Exit(1)
	}
	stopsData, err := client.LoadStops() // ładujemy dane przystanków
	if err != nil {
		fmt.Printf("Error loading stops: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Loaded %d stops successfully!\n", len(stopsData.Stops)) 
	console := ui.NewConsoleUI(client) // tworzymy nowy interfejs konsoli
	if console == nil {
		fmt.Println("Failed to create console UI.")
		os.Exit(1)
	}
	// przypisujemy dane przystanków do konsoli
	// aby mogła z nich korzystać w dalszej części programu
	console.Stops = stopsData.Stops

	for {
		choice := console.ShowMainMenu()
		switch choice {
			case 1:
				console.HandleStopSearch()
			case 2:
				console.HandleParallelMonitoring()
			case 3:
				console.HandleShowRoutes()
			case 4:
				fmt.Println("Exiting...")
				return
			default:
				fmt.Println("Invalid choice. Please try again.")
		}
	}
}