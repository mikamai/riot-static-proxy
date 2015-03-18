package riotAPI

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticRequestBaseURLInterpolatesValues(t *testing.T) {
	s := StaticRequest{Region: "a", Thing: "b", ID: "c"}
	expected := "https://global.api.pvp.net/api/lol/static-data/a/v1.2/b/c"
	assert.Equal(t, expected, s.baseURL())
}

func TestStaticRequestBaseURLWorksIfIDIsNil(t *testing.T) {
	s := StaticRequest{Region: "a", Thing: "b"}
	expected := "https://global.api.pvp.net/api/lol/static-data/a/v1.2/b/"
	assert.Equal(t, expected, s.baseURL())
}

func TestStaticRequestURLReturnsAWorkingURL(t *testing.T) {
	params := url.Values{}
	params.Set("k", "v")
	s := StaticRequest{Region: "a", Thing: "b", ID: "c", Params: params}
	obtained, _ := s.URL()
	assert.Equal(t,
		"https://global.api.pvp.net/api/lol/static-data/a/v1.2/b/c?api_key=foo&k=v",
		obtained.String())
}

func TestStaticRequestURLWorksWithMinimumArgs(t *testing.T) {
	s := StaticRequest{Region: "a", Thing: "b", Params: url.Values{}}
	obtained, _ := s.URL()
	assert.Equal(t,
		"https://global.api.pvp.net/api/lol/static-data/a/v1.2/b/?api_key=foo",
		obtained.String())
}

func TestStaticRequestExecute(t *testing.T) {
	t.Skip("PENDING")
}
