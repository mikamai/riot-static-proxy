package riotAPI

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// APIKey ...
var APIKey string

// StaticRequestParams ...
type StaticRequestParams struct {
	Region string
	Thing  string
	ID     string
	Params url.Values
}

func composeStaticRequestURL(params []interface{}) (*url.URL, error) {
	return url.Parse(
		fmt.Sprintf(
			"https://global.api.pvp.net/api/lol/static-data/%s/v1.2/%s/%s",
			params...,
		),
	)
}

func addAPIKeyToURL(u *url.URL) {
	q := u.Query()
	q.Set("api_key", os.Getenv("RIOT_API_KEY"))
	u.RawQuery = q.Encode()
}

// BuildStaticRequestURL returns an url matching the request
func BuildStaticRequestURL(params StaticRequestParams) (*url.URL, error) {
	u, err := composeStaticRequestURL(
		[]interface{}{params.Region, params.Thing, params.ID},
	)
	if err != nil {
		return u, err
	}
	u.RawQuery = params.Params.Encode()
	addAPIKeyToURL(u)
	return u, nil
}

// Call ...
func Call(u *url.URL) (interface{}, error) {
	log.Print("Calling " + u.String())
	res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%s", data), nil
}
