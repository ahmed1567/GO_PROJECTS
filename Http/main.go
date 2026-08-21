package main

import "fmt"

func main() {
	city := "Cairo"

	lat, lon, err := geocode(city)
	if err != nil {
		fmt.Println("Error finding location:", err)
		return
	}

	weather, err := fetchWeather(lat, lon)
	if err != nil {
		fmt.Println("Error fetching weather:", err)
		return
	}

	fmt.Printf("Location: %s (%.4f, %.4f)\n", city, lat, lon)
	fmt.Printf("Temperature: %.1f°C\n", weather.CurrentWeather.Temperature)
	fmt.Printf("Wind speed: %.1f km/h\n", weather.CurrentWeather.WindSpeed)
	fmt.Printf("Condition: %s\n", describeWeatherCode(weather.CurrentWeather.WeatherCode))

	fmt.Println("\n================READER=====================\n")
	readTest()
	fmt.Println("\n================WRITER=====================\n")
	writeTest()
	fmt.Println("\n=====================================\n")
}
