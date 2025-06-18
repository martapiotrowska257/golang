package ui

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"lab05/api"
	"lab05/models"
)

type ConsoleUI struct {
  client *api.ZTMClient
  Stops  []models.Stop
  reader *bufio.Reader
}

func NewConsoleUI(client *api.ZTMClient) *ConsoleUI {
  return &ConsoleUI{
      client: client,
      reader: bufio.NewReader(os.Stdin),
  }
}

func (ui *ConsoleUI) ShowMainMenu() int {
  fmt.Println("\n--- ZTM Gdańsk Monitor ---")
  fmt.Println("1. Search for a stop and show departures")
  fmt.Println("2. Monitor two routes in parallel")
  fmt.Println("3. Show all routes")
  fmt.Println("4. Exit")
  fmt.Print("Enter your choice: ")

  input, _ := ui.reader.ReadString('\n')
  choice, err := strconv.Atoi(strings.TrimSpace(input))
  if err != nil { return 0 }
  return choice
}

func (ui *ConsoleUI) HandleStopChoose() (int, models.Stop, map[int]string) {
  fmt.Println("\n--- Stop Search ---")
  fmt.Print("Enter stop name: ")
  input, _ := ui.reader.ReadString('\n')
  stopName := strings.TrimSpace(input)
  foundStops := ui.client.SearchStopsByName(ui.Stops, stopName)

  if len(foundStops) == 0 {
    fmt.Println("No stops found with that name.")
    return 0, models.Stop{}, nil
  }

  fmt.Println("Found stops:")
  for i, stop := range foundStops {
    fmt.Printf("%d. %s (ID: %d)\n", i+1, stop.StopName, stop.StopId)
  }

  fmt.Print("Select a stop number to see departures (or any other to go back to menu): ")
  input, _ = ui.reader.ReadString('\n')
  stopNumber, err := strconv.Atoi(strings.TrimSpace(input))
  if err != nil || stopNumber < 1 || stopNumber > len(foundStops) { return 0, models.Stop{}, nil }

  selectedStop := foundStops[stopNumber-1]
  fmt.Printf("Loading departures for stop %s (ID: %d)...\n", selectedStop.StopName, selectedStop.StopId)

  departuresData, err := ui.client.LoadDepartures(selectedStop.StopId)
  if err != nil {
    fmt.Printf("Error loading departures: %v\n", err)
    return 0, models.Stop{}, nil
  }

  if len(departuresData.Departures) == 0 {
    fmt.Println("No departures found for this stop.")
    return 0, models.Stop{}, nil
  }
  LastUpdateTime, _ := time.Parse(time.RFC3339, departuresData.LastUpdate)
  LastUpdateTime = LastUpdateTime.Local()
  fmt.Printf("--- Departures from %s (Last update: %s) ---\n", selectedStop.StopName, LastUpdateTime)
  
  displayedRoutes := make(map[int]string)
  for _, dep := range departuresData.Departures {
    estimatedTime, _ := time.Parse(time.RFC3339, dep.EstimatedTime)
    estimatedTime = estimatedTime.Local()
    fmt.Printf("Line %d -> %s: Estimated: %s (Delay: %d sec)\n",
      dep.RouteId, dep.HeadsignText, estimatedTime.Format("15:04:05"), dep.Delay)
    if _, exists := displayedRoutes[dep.RouteId]; !exists {
      displayedRoutes[dep.RouteId] = dep.HeadsignText
    }  
  }
  return stopNumber, selectedStop, displayedRoutes
}

func (ui *ConsoleUI) HandleStopSearch() {
  stopNumber, selectedStop, displayedRoutes := ui.HandleStopChoose()
  if stopNumber == 0 { return }
  fmt.Print("\nEnter a line number (RouteID) to monitor its upcoming stops (or press Enter to skip): ")
  input, _ := ui.reader.ReadString('\n')
  routeIdStr := strings.TrimSpace(input)
  if routeIdStr == "" { return }

  routeId, err := strconv.Atoi(routeIdStr)
  if err != nil {
    fmt.Println("Invalid route number.")
    return
  }

  if _, exists := displayedRoutes[routeId]; !exists {
    fmt.Println("Route not found in departures.")
    return
  }

  ui.HandleRouteMonitoring(routeId, selectedStop.StopId, nil)
}

func (ui *ConsoleUI) HandleRouteMonitoring(routeId int, stopId int, wg *sync.WaitGroup) {
  if wg != nil {
    defer wg.Done()
  }

  startStopName := "Unknown"
  for _, s := range ui.Stops {
    if s.StopId == stopId {
      startStopName = s.StopName
      break
    }
  }
  
  fmt.Printf("\n--- [Monitor Start] Route %d from %s (ID: %d) ---\n", routeId, startStopName, stopId)
  stopTimesData, err := ui.client.LoadStopTimes(routeId)
  if err != nil {
    fmt.Printf("  [Route %d] Error loading stop times: %v\n", routeId, err)
    return
  }
  if len(stopTimesData.StopTimes) == 0 {
    fmt.Printf(" [Route %d] No stop times found.\n", routeId)
    return
  }

  stopSequenceMap := make(map[int]models.StopTime)
  var sequenceOrder []int
  for _, stopTime := range stopTimesData.StopTimes {
    if _, exists := stopSequenceMap[stopTime.StopSequence]; !exists {

      stopSequenceMap[stopTime.StopSequence] = stopTime
      sequenceOrder = append(sequenceOrder, stopTime.StopSequence)
    }
  }
  sort.Ints(sequenceOrder)

  startSequence := -1
  for _, seq := range sequenceOrder {
    if stopSequenceMap[seq].StopId == stopId {
      startSequence = seq
      break
    }
  }

  if startSequence == -1 {
    fmt.Printf("  [Route %d] Start stop %s (ID: %d) not found in typical sequence.\n", routeId, startStopName, stopId)
    return
  }

  fmt.Printf("  [Route %d] Sequence starting after stop: %s (Seq: %d)\n", routeId, startStopName, startSequence)
  foundNextStops := false
  for _, seq := range sequenceOrder {
    if seq <= startSequence {
      continue
    }
    foundNextStops = true
    nextStopInfo := stopSequenceMap[seq]

    nextStopName := "Unknown Stop Name"
    for _, s := range ui.Stops {
      if s.StopId == nextStopInfo.StopId {
        nextStopName = s.StopName
        break
      }
    }

    fmt.Printf("\n  [Route %d] Checking next stop: %s (ID: %d, Seq: %d)\n",
      routeId, nextStopName, nextStopInfo.StopId, nextStopInfo.StopSequence)

    departuresData, err := ui.client.LoadDepartures(nextStopInfo.StopId)
    if err != nil {
      fmt.Printf("  [Route %d] Error loading departures for %s: %v\n", routeId, nextStopName, err)
      continue
    }

    foundDepartures := false
    for _, dep := range departuresData.Departures {
      if dep.RouteId == routeId {
        estimatedTime, _ := time.Parse(time.RFC3339, dep.EstimatedTime)
        estimatedTime = estimatedTime.Local()
        fmt.Printf("    > [Route %d] -> %s: Estimated arrival: %s (Delay: %d sec)\n",
            routeId, dep.HeadsignText, estimatedTime.Format("15:04:05"), dep.Delay)
        foundDepartures = true
        break
      }
    }
      
    if !foundDepartures {
      fmt.Printf("  > [Route %d] No current real-time departure info found at %s.\n", routeId, nextStopName)
    }
  }
  if !foundNextStops {
    fmt.Printf("  [Route %d] No further stops found in sequence after %s.\n", routeId, startStopName)
  }
  fmt.Printf("--- [Monitor End] Route %d ---\n", routeId)
}

func (ui *ConsoleUI) HandleParallelMonitoring() {
  fmt.Println("\n --- Parallel Route Monitoring ---")
  stopNumber, selectedStop, displayedRoutes := ui.HandleStopChoose()
  if stopNumber == 0 { return }
  
  var routeId1 int
  err := error(nil)
  for {
    fmt.Print("Enter first route number to monitor: ")
    input, _ := ui.reader.ReadString('\n')
    routeIdStr := strings.TrimSpace(input)
    if routeIdStr == "" { return }

    routeId1, err = strconv.Atoi(routeIdStr)
    if err != nil {
      fmt.Println("Invalid route number.")
      continue
    }

    if _, exists := displayedRoutes[routeId1]; !exists {
      fmt.Println("Route not found in departures.")
      continue
    }
    break
  }

  var routeId2 int
  for {
    fmt.Print("Enter second route number to monitor: ")
    input, _ := ui.reader.ReadString('\n')
    routeIdStr := strings.TrimSpace(input)
    if routeIdStr == "" { return }

    routeId2, err = strconv.Atoi(routeIdStr)
    if err != nil {
      fmt.Println("Invalid route number.")
      continue
    }

    if _, exists := displayedRoutes[routeId2]; !exists {
      fmt.Println("Route not found in departures.")
      continue
    }
    break
  }

  fmt.Printf("\nStarting parallel monitoring for Route %d and Route %d from stop %s (ID: %d)...\n",
    routeId1, routeId2, selectedStop.StopName, selectedStop.StopId)

  var wg sync.WaitGroup
  wg.Add(2) 

  go ui.HandleRouteMonitoring(routeId1, selectedStop.StopId, &wg)
  go ui.HandleRouteMonitoring(routeId2, selectedStop.StopId, &wg)

  wg.Wait()

  fmt.Println("\n--- Parallel Monitoring Complete ---")
}

func (ui *ConsoleUI) HandleShowRoutes() {
  fmt.Println("Loading all routes...")
  routesData, err := ui.client.LoadRoutes()
  if err != nil {
    fmt.Printf("Error loading routes: %v\n", err)
    return
  }
  fmt.Printf("--- All Routes (Last update: %s) ---\n", routesData.LastUpdate)
  for _, route := range routesData.Routes {
    fmt.Printf("Route %s (%s) - ID: %d\n", route.RouteShortName, route.RouteLongName, route.RouteId)
  }
}