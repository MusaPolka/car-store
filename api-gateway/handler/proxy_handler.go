package handler

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ProxyToCars(c *gin.Context) {
	target, _ := url.Parse("http://inventory-service:8081")
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(c.Writer, c.Request)
}

func ProxyToOrders(c *gin.Context) {
	target, _ := url.Parse("http://order-service:8082")
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(c.Writer, c.Request)
}

func ProxyToUsers(c *gin.Context) {
	target, _ := url.Parse("http://order-service:8083")
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(c.Writer, c.Request)
}
