package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	APIURL     string
	APIKey     string
	HTTPClient *http.Client
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
		APIURL:     "https://api.openweathermap.org",
		APIKey:     APIKey,
		HTTPClient: http.DefaultClient,
	}
	return c, nil
}

func ParseJSON(r io.Reader) (Conditions, error) {
	var result owmAPIResponse
	err := json.NewDecoder(r).Decode(&result)
	//Decoder := json.NewDecoder(r) //
	//Decoder.Decode(&result) //Decode into the pointer
	if err != nil {
		return Conditions{}, err
	}
	if len(result.Weather) == 0 {
		return Conditions{}, fmt.Errorf("Invalid API response %+v", result)
	}
	return Conditions{
		Temp:    result.Main.Temp,
		Summary: result.Weather[0].Main,
	}, nil
}

func (c Client) Get(APIKey string) (Conditions, error) {
	URL := fmt.Sprintf("%s/data/2.5/weather?q=London,UK&appid=%s", c.APIURL, APIKey)
	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		return Conditions{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Conditions{}, fmt.Errorf("Unexpected http response status code %v", resp.StatusCode)
	}
	w, err := ParseJSON(resp.Body)
	if err != nil {
		return Conditions{}, err
	}
	return w, nil
}
