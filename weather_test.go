package weather_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"weather"

	"github.com/google/go-cmp/cmp"
)

func TestParseJSONReturnsWeatherStructFromJSON(t *testing.T) {
	t.Parallel()
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
		Temp: 289.58,
		Summary: "Clouds",
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}


func TestGetSendsCorrectURL(t *testing.T) {
	t.Parallel()
	var w weather.Conditions
	s := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	APIKey := "dummy"
	c, err := weather.NewClient(APIKey)
	c.APIURL = s.URL
	w, err = c.Get()
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