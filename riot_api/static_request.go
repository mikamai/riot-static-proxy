package riotAPI

import (
	"fmt"
	"net/url"
)

// StaticRequest ...
type StaticRequest struct {
	Region string
	Thing  string
	ID     string
	Params url.Values
}

var staticRequestURLFormat = "https://global.api.pvp.net/api/lol/static-data/%s/v1.2/%s/%s"

func (r StaticRequest) baseURL() string {
	return fmt.Sprintf(
		staticRequestURLFormat,
		[]interface{}{r.Region, r.Thing, r.ID},
	)
}

// URL returns an url matching the request
func (r StaticRequest) URL() (*url.URL, error) {
	return composeURL(r.baseURL(), r.Params)
}

// Execute ...
func (r StaticRequest) Execute() (interface{}, error) {
	u, err := r.URL()
	if err != nil {
		return nil, err
	}
	return callAPI(u)
}
