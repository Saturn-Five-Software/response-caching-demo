package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Setup the gin router
	router := setupRouter()
	// Listen and Serve on default address or address defined through config.json.
	router.Run("0.0.0.0:4443")
}

// setupRouter configures the gin router for the service-ds endpoints
func setupRouter() *gin.Engine {

	// Create a default router
	router := gin.Default()

	// Configure cors to only allow supported HTTP methods
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET"}
	config.AllowHeaders = []string{"Content-Type"}
	router.Use(cors.New(config))

	// Define routes here
	dogCtrl := DogController{}
	router.GET("/dogs", CacheCheck(DogCache), dogCtrl.Get)
	router.POST("/dogs", dogCtrl.Post)

	return router
}
