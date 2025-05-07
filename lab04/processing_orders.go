package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ProcessingOrder(order Order) ProcessResult {
	proccesingTime := time.Duration(time.Duration(rand.Intn(2) + 1)) * time.Second
	time.Sleep(proccesingTime)

	sucess := rand.Float64() > 0.3
	var err error
	if !sucess {
		err = fmt.Errorf("error processing order %d", order.ID)
	}
	
	return ProcessResult{
		OrderID: order.ID,
		CustomerName: order.CustomerName,
		Success: sucess,
		ProcessTime: proccesingTime,
		Error: err,
	}
}

func Worker(id int, jobs <-chan Order, results chan<- ProcessResult, wg *sync.WaitGroup){
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing order %d for %s\n", id, job.ID, job.CustomerName)
		result := ProcessingOrder(job)
		results <- result
	}
}

func RetryFailedOrders(failedOrders <-chan Order, results chan<- ProcessResult, wg *sync.WaitGroup) {
	defer wg.Done()
	
	const maxRetries = 3
	for order := range failedOrders {
		for attempt := 1; attempt <= maxRetries; attempt++ {
			fmt.Printf("Retrying order %d for %s (attempt %d)\n", order.ID, order.CustomerName, attempt)
			result := ProcessingOrder(order)
			if result.Success {
				fmt.Printf("Order %d succeeded on attempt %d\n", order.ID, attempt)
				results <- result
				break
			}
					
			if attempt < maxRetries {
				time.Sleep(2 * time.Second)
			} else {
				fmt.Printf("Order %d failed after %d retries. Giving up.\n", order.ID, maxRetries)
				results <- result
			}
		}
	}
}