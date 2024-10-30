package pokeapi

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"
)

func (c *Client) GetMap(pageURL *string) (Locations, error) {
	url := BaseUrl + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	var loc Locations

	if val, ok := c.cache.Get(url); ok {
		dec := gob.NewDecoder(bytes.NewReader(val))
		err := dec.Decode(&loc)

		if err != nil {
			log.Fatal("cannot decode")
		}
		return loc, nil
	} else {
		var myBuffer bytes.Buffer
		enc := gob.NewEncoder(&myBuffer)
		loc, _ = getLocations(url)

		_ = enc.Encode(loc)

		c.cache.Add(url, myBuffer.Bytes())
		return loc, nil
	}
}

func getLocations(url string) (Locations, error) {
	res, err := http.Get(url)

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
