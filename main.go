//This is a task for Go programmer vacancy at Geeks.Team
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

type Request struct { // Request from client
	Site       []string
	SearchText string
}

type Response struct { // Response to client
	FoundAtSite string
}

var req Request
var resp Response

func main() {
	// Set the router as the default one provided by Gin
	router = gin.Default()

	router.POST("/checktext", checkText)

	// Start serving the application
	router.Run()
}

// Handler to POST /checktext
func checkText(c *gin.Context) {

	// Decoding json into Request struct:
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}
	defer c.Request.Body.Close()

	resp.FoundAtSite = ""

	// Searching the requsted text on the websites
	for i := 0; i < len(req.Site); i++ {

		res, err := http.Get(req.Site[i])

		if err != nil {
			log.Fatalln("Get error: ", err)
		}

		body, err1 := ioutil.ReadAll(res.Body)
		if err1 != nil {
			log.Fatalln("Body error: ", err)
		}

		str := string(body)
		if strings.Contains(str, req.SearchText) {
			resp.FoundAtSite = req.Site[i]
			break
		}
	}

	// Analysis of search process:
	status := http.StatusOK
	if resp.FoundAtSite == "" {
		status = http.StatusNoContent
	}

	// Sending response:
	Render(c, status, gin.H{
		"response": resp})
}

// Response function
func Render(c *gin.Context, status int, data gin.H) {
	c.JSON(
		status,
		data["response"])
}
