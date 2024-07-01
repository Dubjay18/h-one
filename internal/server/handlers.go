package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) DisplayDetailsHandler(c *gin.Context) {
	visitorName := c.Query("visitor_name")
	ip := GetClientIP(c)
	ip = "8.8.8.8"
	location, err := GetLocationFromIP(ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	temperature, err := GetTemperature(ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client_ip": ip, "location": location.City, "greeting": fmt.Sprintf("Hello, %s! The temperature is %.2f degrees Celcius in %s", visitorName, temperature, location.City)})

}
