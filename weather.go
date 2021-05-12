package weather

import (
	"encoding/json"
	"io"
)

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
