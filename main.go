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

func staticRequestAction(c *gin.Context) {
	c.Request.ParseForm()
	request := riotAPI.StaticRequest{
		Region: c.Params.ByName("region"),
		Thing:  c.Params.ByName("thing"),
		ID:     c.Params.ByName("id"),
		Params: c.Request.Form,
	}
	data, err := request.Execute()
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
	}
	c.String(http.StatusOK, data.(string))
}

func prepareCacheStore() cache.CacheStore {
	host := os.Getenv("REDIS_URL")
	if host != "" {
		return cache.NewRedisCache(host, "", time.Second)
	}
	return cache.NewInMemoryStore(time.Second)
}

func setupRollbar(r *gin.Engine) {
	rollbar.Token = os.Getenv("ROLLBAR_KEY")
	if os.Getenv("GO_ENV") != "" {
		rollbar.Environment = os.Getenv("GO_ENV")
	}
}

func setupRiotAPI(r *gin.Engine) {
	riotAPI.APIKey = os.Getenv("RIOT_API_KEY")
}

func setupNewRelic(r *gin.Engine) {
	apiKey := os.Getenv("NEWRELIC_KEY")
	if apiKey != "" {
		gorelic.InitNewrelicAgent(apiKey, "rgts static-proxy", true)
		r.Use(gorelic.Handler)
	}
}

func main() {
	r := gin.Default()
	setupRollbar(r)
	setupRiotAPI(r)
	setupNewRelic(r)
	store := prepareCacheStore()

	cachedStaticRequestAction := cache.CachePage(
		store,
		time.Minute*15,
		staticRequestAction,
	)
	r.GET("/:region/:thing", cachedStaticRequestAction)
	r.GET("/:region/:thing/:id", cachedStaticRequestAction)

	r.Run(":8080")
}
