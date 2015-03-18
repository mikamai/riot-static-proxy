package riotAPI

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComposeURLParsesThePath(t *testing.T) {
	u, _ := composeURL("http://mikamai.com/a", url.Values{})
	assert.Equal(t, u.Host, "mikamai.com")
	assert.Equal(t, u.Path, "/a")
}

func TestAPIKeyToURLAddsParam(t *testing.T) {
	APIKey = "foo"
	u, _ := url.Parse("http://mikamai.com")
	addAPIKeyToURL(u)
	assert.Equal(t, u.RequestURI(), "/?api_key=foo")
}

func TestAPIKeyToURLNotAltersExistingParams(t *testing.T) {
	APIKey = "foo"
	u, _ := url.Parse("http://mikamai.com?a=b")
	addAPIKeyToURL(u)
	assert.Equal(t, u.RequestURI(), "/?a=b&api_key=foo")
}

func TestComposeURLReturnsErrorIfPathIsNotValid(t *testing.T) {
	_, err := composeURL("1/", url.Values{})
	assert.NotNil(t, err)
}

func TestComposeURLAddsParamsToReturn(t *testing.T) {
	values := url.Values{}
	values.Set("foo", "bar")
	url, _ := composeURL("http://mikamai.com", values)
	assert.Equal(t, url.Query().Get("foo"), "bar")
}

func TestComposeURLAddsAPIKeyToReturn(t *testing.T) {
	APIKey = "foo"
	u, _ := composeURL("http://mikamai.com", url.Values{})
	assert.Equal(t, u.RequestURI(), "/?api_key=foo")
}

func TestComposeURLMergesAPIKeyWithParams(t *testing.T) {
	APIKey = "foo"
	values := url.Values{}
	values.Set("a", "b")
	u, _ := composeURL("http://mikamai.com", values)
	assert.Equal(t, u.RequestURI(), "/?a=b&api_key=foo")
}

func TestCallAPI(t *testing.T) {
	t.Skip("PENDING")
}
