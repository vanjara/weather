package weather_test

import (
	"os"
	"testing"
	"weather"

	"github.com/google/go-cmp/cmp"
)

func TestParseJSONReturnsWeatherStructFromJSON(t *testing.T) {
	f, err := os.Open("testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}
	var got weather.Conditions
	got, err = weather.ParseJSON(f)
	if err != nil {
		t.Fatal(err)
	}
	want := weather.Conditions{
		Temp: 286.0,
		Summary: "Clouds",
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}