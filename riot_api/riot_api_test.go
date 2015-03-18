package riotAPI

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComposeURLParsesThePath(t *testing.T) {
	u, _ := composeURL("http://mikamai.com/a", url.Values{})
	assert.Equal(t, "mikamai.com", u.Host)
	assert.Equal(t, "/a", u.Path)
}

func TestAPIKeyToURLAddsParam(t *testing.T) {
	APIKey = "foo"
	u, _ := url.Parse("http://mikamai.com")
	addAPIKeyToURL(u)
	assert.Equal(t, "/?api_key=foo", u.RequestURI())
}

func TestAPIKeyToURLNotAltersExistingParams(t *testing.T) {
	APIKey = "foo"
	u, _ := url.Parse("http://mikamai.com?a=b")
	addAPIKeyToURL(u)
	assert.Equal(t, "/?a=b&api_key=foo", u.RequestURI())
}

func TestComposeURLReturnsErrorIfPathIsNotValid(t *testing.T) {
	_, err := composeURL("1/", url.Values{})
	assert.NotNil(t, err)
}

func TestComposeURLAddsParamsToReturn(t *testing.T) {
	values := url.Values{}
	values.Set("foo", "bar")
	url, _ := composeURL("http://mikamai.com", values)
	assert.Equal(t, "bar", url.Query().Get("foo"))
}

func TestComposeURLAddsAPIKeyToReturn(t *testing.T) {
	APIKey = "foo"
	u, _ := composeURL("http://mikamai.com", url.Values{})
	assert.Equal(t, "/?api_key=foo", u.RequestURI())
}

func TestComposeURLMergesAPIKeyWithParams(t *testing.T) {
	APIKey = "foo"
	values := url.Values{}
	values.Set("a", "b")
	u, _ := composeURL("http://mikamai.com", values)
	assert.Equal(t, "/?a=b&api_key=foo", u.RequestURI())
}

func TestCallAPI(t *testing.T) {
	t.Skip("PENDING")
}
