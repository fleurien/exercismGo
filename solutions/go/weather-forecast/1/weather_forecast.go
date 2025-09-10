// Package weather announces weather condition based on location.
package weather

// CurrentCondition takes a string value for the current weather conditions.
var CurrentCondition string
// CurrentLocation takes a string value for the current location.
var CurrentLocation string

// Forecast takes the city and condition to provide an announcement of the weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
