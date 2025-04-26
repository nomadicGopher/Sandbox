// Package weather provides the weather conditions for a specified location.
package weather

var (
	// CurrentCondition holds the weather condition that associates with CurrentLocation. ie: cloudy, rainy, thunderstorms, sunny, snowing...
	CurrentCondition string
	// CurrentLocation holds the location that associates with CurrentCondition.
	CurrentLocation string
)

// Forecast returns a human readable string pertaining to a specified location's current weather conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
