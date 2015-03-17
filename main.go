package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func staticDataAction(c *gin.Context) {
	c.Request.ParseForm()
	count := strconv.Itoa(len(c.Params))
	count2 := strconv.Itoa(len(c.Request.Form))
	c.String(http.StatusOK, "pong-"+count+"-"+count2)
}

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/:thing", staticDataAction)
		v1.GET("/:thing/:id", staticDataAction)
	}

	r.Run(":8080")
}
