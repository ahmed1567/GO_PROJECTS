package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type weatherResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		WindSpeed   float64 `json:"windspeed"`
		WeatherCode int     `json:"weathercode"`
	} `json:"current_weather"`
}

// weatherCodeText maps Open-Meteo's WMO weather codes to a short
// description — the API returns a number, not text, unlike weatherapi.com.
var weatherCodeText = map[int]string{
	0: "Clear sky",
	1: "Mainly clear", 2: "Partly cloudy", 3: "Overcast",
	45: "Fog", 48: "Depositing rime fog",
	51: "Light drizzle", 53: "Moderate drizzle", 55: "Dense drizzle",
	61: "Slight rain", 63: "Moderate rain", 65: "Heavy rain",
	71: "Slight snow", 73: "Moderate snow", 75: "Heavy snow",
	80: "Slight rain showers", 81: "Moderate rain showers", 82: "Violent rain showers",
	95: "Thunderstorm",
}

// fetchWeather calls Open-Meteo's free forecast API for a given
// coordinate pair — no key required.
func fetchWeather(lat, lon float64) (weatherResponse, error) {
	forecastURL := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true",
		lat, lon,
	)

	resp, err := http.Get(forecastURL)
	if err != nil {
		return weatherResponse{}, fmt.Errorf("failed to make weather request: %w", err)
	}
	defer resp.Body.Close() // MUST close the body!

	if resp.StatusCode != http.StatusOK {
		return weatherResponse{}, fmt.Errorf("weather API returned status %d", resp.StatusCode)
	}

	bodyBytes,err := io.ReadAll(resp.Body)
	if err != nil {
		return weatherResponse{}, fmt.Errorf("failed to read weather response: %w", err)
	}

	var result weatherResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return weatherResponse{}, fmt.Errorf("failed to parse weather JSON: %w", err)
	}

	return result, nil
}

func describeWeatherCode(code int) string {
	if text, ok := weatherCodeText[code]; ok {
		return text
	}
	return fmt.Sprintf("unknown (code %d)", code)
}
