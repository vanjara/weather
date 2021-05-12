package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


type Client struct{
	APIURL string
	APIKey string
}

type Conditions struct {
	Temp    float64
	Summary string
}

type owmAPIResponse struct {
	Main struct {
		Temp float64
	}
	Weather []struct {
		Main string
	}
}

func NewClient(APIKey string) (Client, error) {
	c := Client{
		APIURL: "https://api.openweathermap.org",
		APIKey: APIKey,
	}
	return c, nil
}

func ParseJSON(r io.Reader) (Conditions, error) {
	var result owmAPIResponse
	err := json.NewDecoder(r).Decode(&result)
	if err != nil {
		return Conditions{}, err
	}
	return Conditions{
		Temp:    result.Main.Temp,
		Summary: result.Weather[0].Main,
	}, nil
}

func Get(APIKey string) (Conditions, error) {
	URL := fmt.Sprintf("%s/data/2.5/weather?q=London,UK&appid=%s", APIURL, APIKey)
	resp, err := http.Get(URL)
	if err != nil {
		return Conditions{}, err
	}
	defer resp.Body.Close()
	w, err := ParseJSON(resp.Body)
	if err != nil {
		return Conditions{}, err
	}
	return w, nil
}