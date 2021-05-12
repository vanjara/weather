package main

import (
	"fmt"
	"log"
	"os"
	"weather"
)

func main() {
	f, err := os.Open("testdata/weather.json")
	if err != nil {
		log.Fatal(err)
	}
	w, err := weather.ParseJSON(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("It's %s, and %.1fK\n", w.Summary, w.Temp)
}
