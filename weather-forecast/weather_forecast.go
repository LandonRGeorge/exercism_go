// Package weather does xyz...
package weather

// CurrentCondition represents xyz...
var CurrentCondition string
// CurrentLocation represents xyz...
var CurrentLocation string

// Forecast takes city and condition and returns a formatted string...
// More text here.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
