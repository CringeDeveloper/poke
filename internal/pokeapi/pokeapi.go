package pokeapi

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func (c *Client) GetExplore(name string) (LocationArea, error) {
	url := BaseUrl + "location-area/" + name

	var area LocationArea

	if val, ok := c.cache.Get(url); ok {
		dec := gob.NewDecoder(bytes.NewReader(val))
		err := dec.Decode(&area)

		if err != nil {
			return LocationArea{}, err
		}
		return area, nil
	} else {
		var myBuffer bytes.Buffer
		enc := gob.NewEncoder(&myBuffer)

		res, err := c.httpClient.Get(url)
		if err != nil {
			return LocationArea{}, err
		}

		defer res.Body.Close()

		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&area)
		if err != nil {
			return LocationArea{}, err
		}

		_ = enc.Encode(area)

		c.cache.Add(url, myBuffer.Bytes())
		return area, nil
	}
}

func (c *Client) GetMap(pageURL *string) (Locations, error) {
	url := BaseUrl + "location-area"

	if pageURL != nil {
		url = *pageURL
	}

	var loc Locations

	if val, ok := c.cache.Get(url); ok {
		dec := gob.NewDecoder(bytes.NewReader(val))
		err := dec.Decode(&loc)

		if err != nil {
			return Locations{}, err
		}
		return loc, nil
	} else {
		var myBuffer bytes.Buffer
		enc := gob.NewEncoder(&myBuffer)

		res, err := c.httpClient.Get(url)
		if err != nil {
			return Locations{}, err
		}

		defer res.Body.Close()

		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&loc)
		if err != nil {
			return Locations{}, err
		}

		_ = enc.Encode(loc)

		c.cache.Add(url, myBuffer.Bytes())
		return loc, nil
	}
}
