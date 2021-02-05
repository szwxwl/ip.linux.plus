package main

import (
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
    "github.com/ip2location/ip2location-go"
)

func main() {
	db, err := ip2location.OpenDB("/go/bin/IP2LOCATION-LITE-DB3.IPV6.BIN")
    if err != nil {
        return
    }
	router := gin.Default()
	router.LoadHTMLGlob("/go/templates/*")
	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) {
		ip := c.GetHeader('Cf-Connecting-Ip')
		if ip == "" {
			ip = c.ClientIP()
		}
		returnType := c.DefaultQuery("type", "")
		switch returnType {
			case "json": 
				c.JSON(http.StatusOK, gin.H{
					"ip": ip,
				})
				break
			case "jsonp":
				callback := c.DefaultQuery("callback", "callback")
				c.String(http.StatusOK, fmt.Sprintf("%s({\"ip\": \"%s\"});", callback, ip))
				break
			default: 
				c.String(http.StatusOK, ip)
		}
	})
	router.GET("/search/*ip", func(c *gin.Context) {
		ip := strings.TrimLeft(c.Param("ip"), "/")
		if ip == "" {
			ip = c.ClientIP()
		}
		results, err := db.Get_all(ip)
		if err != nil {
			fmt.Print(err)
			return
		}
		returnType := c.DefaultQuery("type", "json")
		if returnType == "jsonp" {
			callback := c.DefaultQuery("callback", "callback")
			resp := gin.H{
				"ip": ip,
				"country_short": results.Country_short,
				"country_long": results.Country_long,
				"region": results.Region,
				"city": results.City,
			}
			jsonString, err := json.Marshal(resp)
			if err != nil {
				c.String(http.StatusOK, fmt.Sprintf("failed!"))
				return
			}
			c.String(http.StatusOK, fmt.Sprintf("%s(%s);", callback, jsonString))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"ip": ip,
			"country_short": results.Country_short,
			"country_long": results.Country_long,
			"region": results.Region,
			"city": results.City,
		})
	})
	router.GET("/doc", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})
	router.Run(":80")
}
