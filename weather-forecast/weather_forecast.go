// Package weather describes where the file is located.
package weather

// CurrentCondition describes current condition.
var CurrentCondition string
// CurrentLocation describes current location.
var CurrentLocation string

// Forecast displays the current location and condition in the following format: <location> - current weather condition: <condition>.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
