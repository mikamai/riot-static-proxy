package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/brandfolder/gin-gorelic"
	"github.com/gin-gonic/contrib/cache"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

func composeStaticDataAPI(params httprouter.Params, values url.Values) *url.URL {
	baseURL, err := url.Parse(
		fmt.Sprintf(
			"https://global.api.pvp.net/api/lol/static-data/%s/v1.2/%s/%s",
			params.ByName("region"), params.ByName("thing"), params.ByName("id"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	return baseURL
}

func addAPIKeyToURL(u *url.URL) {
	q := u.Query()
	q.Set("api_key", os.Getenv("RIOT_API_KEY"))
	u.RawQuery = q.Encode()
}

func callStaticDataAPI(u *url.URL) string {
	addAPIKeyToURL(u)
	log.Print("Calling " + u.String())
	res, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", data)
}

func staticDataAction(c *gin.Context) {
	c.Request.ParseForm()
	u := composeStaticDataAPI(c.Params, c.Request.Form)
	c.String(http.StatusOK, callStaticDataAPI(u))
}

//
// func prepareCacheStore() cache.CacheStore {
// 	host := os.Getenv("REDIS_URL")
// 	if host == "" {
// 		host = "localhost:6379"
// 	}
// 	return cache.NewRedisCache("192.121.111.111", "", time.Minute*15)
// }

func main() {
	r := gin.Default()
	gorelic.InitNewrelicAgent(os.Getenv("NEWRELIC_KEY"), "rgts static-proxy", true)
	r.Use(gorelic.Handler)
	store := cache.NewInMemoryStore(time.Second)
	v1 := r.Group("/v1")
	{
		v1.GET("/:region/:thing", cache.CachePage(store, time.Minute*15, staticDataAction))
		v1.GET("/:region/:thing/:id", cache.CachePage(store, time.Minute*15, staticDataAction))
	}

	r.Run(":8080")
}
