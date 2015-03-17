package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

func composeStaticDataAPI(params httprouter.Params, values url.Values) string {
	baseURL, err := url.Parse(
		fmt.Sprintf(
			"https://global.api.pvp.net/api/lol/static-data/%s/v1.2/%s/%s",
			params.ByName("region"), params.ByName("thing"), params.ByName("id"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	values.Set("api_key", os.Getenv("RIOT_API_KEY"))
	baseURL.RawQuery = values.Encode()
	return baseURL.String()
}

func staticDataAction(c *gin.Context) {
	c.Request.ParseForm()

	baseURL := composeStaticDataAPI(c.Params, c.Request.Form)
	c.String(http.StatusOK, baseURL)
}

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/:region/:thing", staticDataAction)
		v1.GET("/:region/:thing/:id", staticDataAction)
	}

	r.Run(":8080")
}
