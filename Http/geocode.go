package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type geocodeResponse struct {
	Results []struct {
		Name      string  `json:"name"`
		Country   string  `json:"country"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"results"`
}

// geocode turns a city name into coordinates using Open-Meteo's free
// geocoding API — no key required.
func geocode(city string) (lat float64, lon float64, err error) {
	geocodeURL := "https://geocoding-api.open-meteo.com/v1/search?name=" + url.QueryEscape(city) + "&count=1"

	resp, err := http.Get(geocodeURL)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to make geocode request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("geocode API returned status %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read geocode response: %w", err)
	}

	var result geocodeResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return 0, 0, fmt.Errorf("failed to parse geocode JSON: %w", err)
	}

	if len(result.Results) == 0 {
		return 0, 0, fmt.Errorf("no location found for %q", city)
	}

	first := result.Results[0]
	return first.Latitude, first.Longitude, nil
}
