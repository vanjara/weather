package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	//var result map[string]interface{}
	var result struct {
		Cod string
		//List []struct{ Main interface{} }
		List []struct {
			Main struct {
				Temp float64
			}
			Weather []struct {
				Main string
			}
		}
	}
	f, err := os.Open("testdata/weather.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(f).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.List[0].Main.Temp)
	fmt.Println(result.List[0].Weather[0].Main)
	// for key, value := range result {
	// 	fmt.Println(key, value)

	// }
}
