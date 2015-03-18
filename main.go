package main

import (
	"net/http"
	"os"
	"riot-static-proxy/riot_api"
	"time"

	"github.com/brandfolder/gin-gorelic"
	"github.com/gin-gonic/contrib/cache"
	"github.com/gin-gonic/gin"
	"github.com/stvp/rollbar"
)

func staticDataAction(c *gin.Context) {
	c.Request.ParseForm()
	requestParams := riotAPI.StaticRequestParams{
		Region: c.Params.ByName("region"),
		Thing:  c.Params.ByName("thing"),
		ID:     c.Params.ByName("id"),
		Params: c.Request.Form,
	}
	u, err := riotAPI.BuildStaticRequestURL(requestParams)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
	}
	res, err := riotAPI.Call(u)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
	}
	c.String(http.StatusOK, res.(string))
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
	rollbar.Token = os.Getenv("ROLLBAR_KEY")
	riotAPI.APIKey = os.Getenv("RIOT_API_KEY")
	if os.Getenv("GO_ENV") != "" {
		rollbar.Environment = os.Getenv("GO_ENV")
	}

	r := gin.Default()
	gorelic.InitNewrelicAgent(os.Getenv("NEWRELIC_KEY"), "rgts static-proxy", true)
	r.Use(gorelic.Handler)
	store := cache.NewInMemoryStore(time.Second)

	r.GET("/:region/:thing", cache.CachePage(store, time.Minute*15, staticDataAction))
	r.GET("/:region/:thing/:id", cache.CachePage(store, time.Minute*15, staticDataAction))

	r.Run(":8080")
}
