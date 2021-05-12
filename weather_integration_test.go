//+build integration

package weather_test

import (
	"os"
	"testing"
	"weather"
)

func TestWeatherIntegration(t *testing.T) {
	var w weather.Conditions
	APIKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if APIKey == "" {
		t.Fatal("Please set OPENWEATHERMAP_API_KEY to run this test")
	}
	w, err := weather.Get(APIKey)
	if err != nil {
		t.Fatal(err)
	}
	if w.Summary == "" {
		t.Errorf("empty summary: %+v", w)
	}
	if w.Temp == 0 {
		t.Errorf("zero temperature: %+v", w)
	}
}