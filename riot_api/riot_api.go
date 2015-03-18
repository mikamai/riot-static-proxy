package riotAPI

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// APIKey ...
var APIKey string

func addAPIKeyToURL(u *url.URL) {
	q := u.Query()
	q.Set("api_key", APIKey)
	u.RawQuery = q.Encode()
}

func composeURL(path string, params url.Values) (*url.URL, error) {
	u, err := url.ParseRequestURI(path)
	if err != nil {
		return nil, err
	}
	u.RawQuery = params.Encode()
	addAPIKeyToURL(u)
	return u, nil
}

func callAPI(u *url.URL) (interface{}, error) {
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
