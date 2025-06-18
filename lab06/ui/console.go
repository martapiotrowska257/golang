package ui

import (
	"bufio"
	"fmt"
	"os"
	"lab06/functions"
	"lab06/structures"
	"strconv"
	"strings"
	"time"
)

func ShowMainMenu() int {
	fmt.Println("\n--- Stock Market Indicators ---")
	fmt.Println("1. Show data")
	fmt.Println("2. Calculate indicators (RSI, EMA, ATR)")
	fmt.Println("3. Exit")
	fmt.Print("Enter your choice: ")

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0
	}
	return choice
}

func ChooseData() int {
	filesList, err := functions.ListFiles("source")
	if err != nil {
		fmt.Println("Error listing files:", err)
		fmt.Println("Please try again.")
		return 0
	}

	fmt.Println("\n--- Choose Data ---")
	for i, file := range filesList {
		fmt.Printf("%d. %s\n", i+1, file)
	}
	fmt.Print("Enter your choice: ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || choice < 1 || choice > len(filesList) {
		fmt.Println("Invalid choice. Please try again.")
		return 0
	}
	return choice

}

func ShowData(data []structures.Data) {
	fmt.Println("\n--- Data Points ---")
	if len(data) == 0 {
		fmt.Println("No data points available.")
		return
	}

	for i, d := range data {
		fmt.Printf("%d. High: %.2f, Low: %.2f, Last: %.2f\n", i+1, d.High, d.Low, d.Last)
	}
	fmt.Println("Total data points:", len(data))
}

func CalculateChoice() int {
	fmt.Println("\n--- Calculate Indicators ---")
	fmt.Println("1. RSI")
	fmt.Println("2. EMA")
	fmt.Println("3. ATR")
	fmt.Println("4. Back to main menu")
	fmt.Print("Enter your choice: ")

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || choice < 1 || choice > 4 {
		return 0
	}
	return choice
}

func ChooseDataPeriod() (time.Time, time.Time) {
	fmt.Print("Enter start date (MM/DD/YYYY): ")
	startInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	startDate, err := time.Parse("01/02/2006", strings.TrimSpace(startInput))
	if err != nil {
		fmt.Println("Invalid date format. Please try again.")
		return time.Time{}, time.Time{}
	}

	fmt.Print("Enter end date (MM/DD/YYYY): ")
	endInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	endDate, err := time.Parse("01/02/2006", strings.TrimSpace(endInput))
	if err != nil {
		fmt.Println("Invalid date format. Please try again.")
		return time.Time{}, time.Time{}
	}

	return startDate, endDate
}
