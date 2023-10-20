// Package weather: provides tools to render forecasts.
package weather

// CurrentCondition: Stores the current weather condition, used in the
// previous call of Forecast.
var CurrentCondition string

// CurrentLocation: Stores the current location, used in the previous call of
// Forecast.
var CurrentLocation string

// Forecast: Render a forecast provided a city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
