package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

type Location struct {
	City string  `json:"city"`
	lat  float64 `json:"lat"`
	lon  float64 `json:"lon"`
}
type Weather struct {
	Current struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
	} `json:"current"`
}

func GetClientIP(c *gin.Context) string {
	// If your app is behind a proxy, you might want to check for 'X-Forwarded-For' header
	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1" // For testing purposes, replace IPv6 loopback with IPv4 loopback
	}
	return ip
}
func GetLocationFromIP(ip string) (*Location, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, berr := io.ReadAll(resp.Body)
	if berr != nil {
		return nil, berr
	}

	log.Printf("Response from API: %s\n", string(body)) // Log the raw response for debugging

	var location Location
	if err := json.Unmarshal(body, &location); err != nil {
		return nil, err
	}

	return &location, nil
}

func GetTemperature(ip string) (float64, error) {
	url := fmt.Sprintf("https://weatherapi-com.p.rapidapi.com/current.json?q=%s", ip)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("x-rapidapi-key", os.Getenv("RAPIDAPI_KEY"))
	req.Header.Add("x-rapidapi-host", "weatherapi-com.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	log.Printf("Response from Weather API: %s\n", string(body)) // Log the raw response for debugging

	var weather Weather
	if err := json.Unmarshal(body, &weather); err != nil {
		return 0, err
	}

	return weather.Current.TempC, nil
}
