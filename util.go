package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// abortWithMessage aborts with an error message.
func abortWithMessage(code int, err error, c *gin.Context) {
	c.String(code, http.StatusText(code))
	c.Error(err)
	c.Abort()
}
