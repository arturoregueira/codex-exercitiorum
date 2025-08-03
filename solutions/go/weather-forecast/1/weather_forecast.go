// Package weather, this Package weather implements a client for weather.com API at https://api.weather.com/v1/. Sometimes this API is called the "Weather Underground API".
package weather

// CurrentCondition is a string that contaims data on the current weather conditions.
var CurrentCondition string

// CurrentLocation is a string that contaims data on the current location.
var CurrentLocation string

// Forecast, this returs a string that infoms what are ther weather conditions in a perticular city.
func Forecast(city, condition string) string {

	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
