package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	orders := make(chan Order, 100)
	results := make(chan ProcessResult, 100)
	failedOrders := make(chan Order, 100)
	done := make(chan bool)

	go CollectResults(results, failedOrders, done)

	var wg sync.WaitGroup
	numberOfWorkers := 5
	for i:= 1; i <= numberOfWorkers; i++ {
		wg.Add(1)
		go Worker(i, orders, results, &wg)
	}

	go func() {
		for i := 1; i <= 5; i++ {
			order := GenerateRandomOrder(i)
			fmt.Printf("Generated order %d for %s for %.2f zł\n", order.ID, order.CustomerName, order.TotalAmount)
			orders <- order
			
			time.Sleep(time.Duration(rand.Intn(2) + 1) * time.Second)
		}
		
		close(orders)
	}()

	wg.Wait()
	
	println("\nAll orders have been processed.")
	println("Trying to process failed orders...")
	
	wg.Add(1)
	go RetryFailedOrders(failedOrders, results, &wg)
	
	close(failedOrders)
	wg.Wait()

	close(results)
	<-done

	fmt.Println("All GoRoutines finished their jobs.")
}