package models

type Stop struct { // struktura reprezentująca przystanek
	StopId        int     `json:"stopId"`
	StopCode      string  `json:"stopCode"`
	StopName      string  `json:"stopName"`
	StopShortName string  `json:"stopShortName"`
	StopDesc      string  `json:"stopDesc"`
	StopLat       float64 `json:"stopLat"` // szerokość geograficzna przystanku
	StopLon       float64 `json:"stopLon"` // długość geograficzna przystanku
	ZoneId        int     `json:"zoneId"`
}

type StopsResponse struct { // kontener na liste przystankow (struktura reprezentująca odpowiedź z serwera zawierającą dane przystanków)
	LastUpdate string `json:"lastUpdate"`
	Stops      []Stop `json:"stops"`
}

type Departure struct { // struktura reprezentująca odjazd z przystanku
	Id                     string `json:"id"`
	Delay                  int    `json:"delayInSeconds"`
	EstimatedTime          string `json:"estimatedTime"`
	HeadsignText           string `json:"headsign"`
	RouteId                int    `json:"routeId"`
	RouteShortName         string `json:"routeShortName"`
	ScheduledTripStartTime string `json:"scheduledTripStartTime"`
	TripId                 int    `json:"tripId"`
	StatusMessage          string `json:"status"`
	Theoretically          string `json:"theoreticalTime"`
	Timestamp              string `json:"timestamp"`
	Trip                   int    `json:"trip"`
	VehicleCode            int    `json:"vehicleCode"`
	VehicleId              int    `json:"vehicleId"`
	VehicleService         string `json:"vehicleService"`
}

type DeparturesResponse struct { // kontener na dane o odjazdach (struktura reprezentująca odpowiedź z serwera zawierającą dane odjazdów z przystanku)
	LastUpdate string      `json:"lastUpdate"`
	Departures []Departure `json:"departures"`
}

type StopTime struct { // struktura reprezentująca czas przyjazdu i odjazdu na przystanku
	RouteId       int    `json:"routeId"`
	TripId        int    `json:"tripId"`
	ArrivalTime   string `json:"arrivalTime"`
	DepartureTime string `json:"departureTime"`
	StopId        int    `json:"stopId"`
	StopSequence  int    `json:"stopSequence"`
}

type StopTimesResponse struct { // kontener na dane o czasach przystankow (struktura reprezentująca odpowiedź z serwera zawierającą czasy przyjazdów i odjazdów na przystanku)
	LastUpdate string     `json:"lastUpdate"`
	StopTimes  []StopTime `json:"stopTimes"`
}

type Route struct { // struktura reprezentująca trasę
	RouteId        int    `json:"routeId"`
	RouteShortName string `json:"routeShortName"`
	RouteLongName  string `json:"routeLongName"`
	RouteType      string `json:"routeType"`
	AgencyId       int    `json:"agencyId"`
}

type RoutesResponse struct { // kontener na dane o trasach (struktura reprezentująca odpowiedź z serwera zawierającą dane tras)
	LastUpdate string  `json:"lastUpdate"`
	Routes     []Route `json:"routes"`
}