package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures the endpoints of the API Gateway.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Car service
	router.Any("/cars/*path", CarsProxyHandler)

	// Order service
	router.Any("/orders/*path", OrdersProxyHandler)

	// User service (fix: use UsersProxyHandler, not OrdersProxyHandler!)
	router.Any("/users/*path", UsersProxyHandler)

	// Optionally: health checks
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "gateway up"})
	})

	return router
}

// CarsProxyHandler forwards to the car service (inventory service)
func CarsProxyHandler(c *gin.Context) {
	target, err := url.Parse("http://localhost:8081")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid car service URL"})
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	origDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		origDirector(req)
		req.URL.Path = c.Param("path")
		req.Host = target.Host
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

// OrdersProxyHandler forwards to the order service
func OrdersProxyHandler(c *gin.Context) {
	target, err := url.Parse("http://localhost:8082")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid order service URL"})
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	origDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		origDirector(req)
		req.URL.Path = c.Param("path")
		req.Host = target.Host
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

// UsersProxyHandler forwards to the user service
func UsersProxyHandler(c *gin.Context) {
	target, err := url.Parse("http://localhost:8083")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user service URL"})
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	origDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		origDirector(req)
		req.URL.Path = c.Param("path")
		req.Host = target.Host
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
