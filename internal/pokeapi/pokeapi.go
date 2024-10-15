package pokeapi

import (
	"encoding/json"
	"net/http"
)

func GetLocations(url *string) (Locations, error) {
	res, err := http.Get(*url)

	if err != nil {
		return Locations{}, err
	}

	defer res.Body.Close()

	var loc Locations

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&loc)
	if err != nil {
		return Locations{}, err
	}

	return loc, nil
}
