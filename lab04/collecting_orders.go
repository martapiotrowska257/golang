package main

import "fmt"

func CollectResults(results <-chan ProcessResult, failedOrders chan<- Order, done chan<- bool) {
	var (
		totalOrders int
		successOrdersCount int
		failedOrdersCount int
	)

	for result := range results {
		totalOrders++
		if result.Success {
			successOrdersCount++
			fmt.Printf("✅ Order %d processed successfully in %s\n", result.OrderID, result.ProcessTime)
		} else {
			failedOrdersCount++
			fmt.Printf("❌ Order %d processed failed in %s\n", result.OrderID, result.ProcessTime)
			FindOrder(result.OrderID, result.CustomerName, failedOrders)
		}
	}

	successOrdersCount = successOrdersCount - failedOrdersCount
	totalOrders = totalOrders - failedOrdersCount
	
	successRate := float64(successOrdersCount) / float64(totalOrders) * 100
	failureRate := float64(failedOrdersCount) / float64(totalOrders) * 100

	fmt.Printf("\n Summary:\n")
	fmt.Printf("Total orders: %d\n", totalOrders)
	fmt.Printf("Successful orders: %d (%.2f%%)\n", successOrdersCount, successRate)
	fmt.Printf("Failed orders: %d (%.2f%%)\n", failedOrdersCount, failureRate)

	done <- true
}

func FindOrder(orderID int, customerName string, failedOrders chan<- Order) {
	order := Order{
		ID: orderID,
		CustomerName: customerName,
	}
	failedOrders <- order
	fmt.Printf("Order %d for %s added to the list of failed orders\n", order.ID, order.CustomerName)
}