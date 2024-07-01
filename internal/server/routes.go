package server

import (
	geo "github.com/cjgiridhar/gin-geo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Use(geo.Default("github.com/cjgiridhar/gin-geo/db/GeoLite2-City.mmdb"))
	r.GET("/", s.HelloWorldHandler)
	r.GET("/api/hello", s.DisplayDetailsHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}
