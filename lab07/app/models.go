package app

type City struct {
	Name      string
	Latitude  float64
	Longitude float64
}

type WeatherData struct {
	CityName            string
	Timestamp           string
	Temperature         float64
	ApparentTemperature float64
	Humidity            float64
	Precipitation       float64
	WindSpeed           float64
	Pressure            float64
}

type ThresholdConfig struct {
	HighTemp  float64 `yaml:"high_temperature_celsius"`
	LowTemp   float64 `yaml:"low_temperature_celsius"`
	HighWind  float64 `yaml:"strong_wind_speed_kmh"`
	HeavyRain float64 `yaml:"heavy_precipitation_mm_daily"`
}

type Alert struct {
	Type  string
	Date  string
	Value float64
}