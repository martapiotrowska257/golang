package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"lab05/models"
)

const ( 
	baseURL = "https://ckan2.multimediagdansk.pl"
	timeout = 5 * time.Second
)

type ZTMClient struct {	// struktura reprezentująca klienta ZTM
	client *http.Client
}

func NewZTMClient() *ZTMClient { // funkcja tworząca nowego klienta ZTM
	return &ZTMClient{	
		client: & http.Client{
			Timeout: timeout,
		},
	}
}

// fetchujemy dane z API ZTM
func (c *ZTMClient) fetchData(url string, target any) error {
	resp, err := c.client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data from %s: %w", url, err)
	}
	defer resp.Body.Close() 	 
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data from %s: %s", url, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)  
	if err!=nil{
		return fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, target) // używamy json.Unmarshal do parsowania odpowiedzi JSON
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}
	
	return nil
}

func (c *ZTMClient) getCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func (c *ZTMClient) LoadStops() (models.StopsResponse, error) {
	var stopsResponse models.StopsResponse
	url := fmt.Sprintf("%s/stops?date=%s", baseURL, c.getCurrentDate())
	err := c.fetchData(url, &stopsResponse)
	if err != nil {
		return models.StopsResponse{}, fmt.Errorf("failed to load stops: %w", err)
	}

	return stopsResponse, nil
}

func (c *ZTMClient) SearchStopsByName(stops []models.Stop, name string) []models.Stop {
	var results []models.Stop
	name = strings.ToLower(name)
	for _, stop := range stops {
		if strings.Contains(strings.ToLower(stop.StopName), name) {
			results = append(results, stop)
		}
	}

	return results
}

func (c *ZTMClient) LoadDepartures(stopId int) (models.DeparturesResponse, error) {
	var departuresResponse models.DeparturesResponse
	url := fmt.Sprintf("%s/departures?stopId=%d&date=%s", baseURL, stopId, c.getCurrentDate())
	err := c.fetchData(url, &departuresResponse)
	if err != nil {
		return models.DeparturesResponse{}, fmt.Errorf("failed to load departures: %w", err)
	}

	return departuresResponse, nil
}

func (c *ZTMClient) LoadStopTimes(routeId int) (models.StopTimesResponse, error) {
	var stopTimesResponse models.StopTimesResponse
	url := fmt.Sprintf("%s/stopTimes?date=%s&routeId=%d", baseURL, c.getCurrentDate(), routeId)
	err := c.fetchData(url, &stopTimesResponse)
	if err != nil {
		return models.StopTimesResponse{}, fmt.Errorf("failed to load stop times: %w", err)
	}

	return stopTimesResponse, nil
}

func (c *ZTMClient) LoadRoutes() (models.RoutesResponse, error) {
	var routesData models.RoutesResponse
	url := fmt.Sprintf("%s/routes?date=%s", baseURL, c.getCurrentDate())
	err := c.fetchData(url, &routesData)
	if err != nil {
		return models.RoutesResponse{}, fmt.Errorf("failed to load routes: %w", err)
	}

	return routesData, nil
}