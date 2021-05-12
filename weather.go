package weather

import (
	"encoding/json"
	"fmt"
	"io"
)

type Conditions struct{
	Temp float64
	Summary string
}

type owmAPIResponse struct {
	List []struct {
		Main struct {
			Temp float64
		}
		Weather []struct {
			Main string
		}
	}
}

func ParseJSON(r io.Reader) (Conditions, error) {
	var result owmAPIResponse
	err := json.NewDecoder(r).Decode(&result)
	if err != nil {
		return Conditions{}, err
	}
	if len(result.List) == 0 {
		return Conditions{}, fmt.Errorf("bad API response: %+v", result)
	}
	return Conditions{
		Temp: result.List[0].Main.Temp,
		Summary: result.List[0].Weather[0].Main,
	}, nil
}