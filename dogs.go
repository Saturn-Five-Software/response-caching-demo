package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

//DogController is a controller for getting dogs.
type DogController struct{}

// DogCache is the response cache for dog requests.
var DogCache = cache.New(5*time.Minute, 10*time.Minute)

// Get returns some doggos!
func (ctrl DogController) Get(c *gin.Context) {
	// Mimic some database delay
	time.Sleep(500 * time.Millisecond)

	c.JSON(http.StatusOK, gin.H{
		"dogs": []string{"Yorkie", "Hound", "Daschund", "Beagle"},
	})
}

// Post add some doggos!
func (ctrl DogController) Post(c *gin.Context) {
	// For sample app purposes we won't actually create any resources.
	// But we will clear the cache!
	DogCache.Flush()
	c.Status(201)
}
