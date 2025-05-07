package main

import "math/rand"

var itemPrices = map[string]float64{
    "Chleb":       4.50,
    "Ketchup":     7.99,
    "Musztarda":   5.49,
    "Kiełbasa":    12.99,
    "Polędwiczki": 24.99,
    "Camembert":   8.79,
    "Woda":        2.49,
    "Cola":        6.99,
    "Winogrona":   9.99,
}

func GenerateRandomOrder(id int) Order {
    customerNames := []string{"Marta", "Foka", "Lulu", "Franek", "Gustaw", "Agata"}
    items := []string{"Chleb", "Ketchup", "Musztarda", "Kiełbasa", "Polędwiczki", "Camembert", "Woda", "Cola", "Winogrona"}

    numItems := rand.Intn(3) + 1
    randomItems := make([]string, 0, numItems)
    
    totalAmount := 0.0
    
    for range numItems {
        selectedItem := items[rand.Intn(len(items))]
        randomItems = append(randomItems, selectedItem)
        
        totalAmount += itemPrices[selectedItem]
    }
    
    return Order{
        ID:           id,
        CustomerName: customerNames[rand.Intn(len(customerNames))],
        Items:        randomItems,
        TotalAmount:  totalAmount,
    }
}