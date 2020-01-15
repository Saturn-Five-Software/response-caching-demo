package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

// bodyCacheWriter is used to cache responses in gin.
type bodyCacheWriter struct {
	gin.ResponseWriter
	cache      *cache.Cache
	requestURI string
}

// Write a JSON response to gin and cache the response.
func (w bodyCacheWriter) Write(b []byte) (int, error) {
	// Write the response to the cache only if a success code
	status := w.Status()
	if 200 <= status && status <= 299 {
		w.cache.Set(w.requestURI, b, cache.DefaultExpiration)
	}

	// Then write the response to gin
	return w.ResponseWriter.Write(b)
}

// CacheCheck sees if there are any cached responses and returns
// the cached response if one is available.
func CacheCheck(cache *cache.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the ignoreCache parameter
		ignoreCache := strings.ToLower(c.Query("ignoreCache")) == "true"

		// See if we have a cached response
		response, exists := cache.Get(c.Request.RequestURI)
		if !ignoreCache && exists {
			// If so, use it
			c.Data(200, "application/json", response.([]byte))
			c.Abort()
		} else {
			// If not, pass our cache writer to the next middleware
			bcw := &bodyCacheWriter{cache: cache, requestURI: c.Request.RequestURI, ResponseWriter: c.Writer}
			c.Writer = bcw
			c.Next()
		}
	}
}
