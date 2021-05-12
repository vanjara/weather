package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"weather"
)

func main() {
	APIKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	URL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=London,UK&appid=%s", APIKey)
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	w, err := weather.ParseJSON(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("It's %s, and %.1fK\n", w.Summary, w.Temp)
}
