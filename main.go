package main

import (
	"fmt"
	//"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type sinucaevent struct {
	ID        int    `json:"id"`
	UUID      string `json:"uuid"`
	PlayerA   string `json:"playera"`
	PlayerB   string `json:"playerb"`
	DateTime  string `json:"datetime"`
	StreamURL string `json:"streamurl"`
	CanBet    bool   `json:"canbet"`
	Location  string `json:"location"`
}

var sinucaevents = []sinucaevent{
	{ID: 0, UUID: "c32746f0-74d5-11ec-bd15-8d09a4545895", PlayerA: "Jogador 1", PlayerB: "Gugu Sem Medo", DateTime: "2012-04-23T18:25:43.511Z", StreamURL: "https://www.youtube.com/watch?v=xVMlw60nJ7Y", CanBet: true, Location: "Salvador-BA"},
	{ID: 1, UUID: "c32746f1-74d5-11ec-bd15-8d09a4545895", PlayerA: "Gugu Sem Medo", PlayerB: "Baianinho de Mauá", DateTime: "2021-04-23T18:25:43.511Z", StreamURL: "https://www.youtube.com/watch?v=xVMlw60nJ7Y", CanBet: true, Location: "São Paulo-SP"},
	{ID: 2, UUID: "c32746f2-74d5-11ec-bd15-8d09a4545895", PlayerA: "Luciano", PlayerB: "Adamastor", DateTime: "2021-12-10T03:30:00.000Z", StreamURL: "https://www.youtube.com/watch?v=xVMlw60nJ7Y", CanBet: true, Location: "Brasília-DF"},
}

func main() {
	router := gin.Default()
	router.GET("/sinucaevents", getSinucaEvents)
	router.GET("/sinucaevents/:id", getSinucaEventByID)
	router.GET("/sinucaevents/uuid/:uuid", getSinucaEventByUUID)
	router.POST("/sinucaevents", postSinucaEvents)

	// Start server with HTTPS TLS:
	//router.RunTLS("localhost:8443", "./data/server.crt", "./data/server.key")
	// Start server with HTTP:
	router.Run("localhost:8080")
}

// getSinucaEvents responds with the list of all sinucaEvents as JSON.
func getSinucaEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sinucaevents)
}

// postSinucaEvents adds an sinucaEvent from JSON received in the request body.
func postSinucaEvents(c *gin.Context) {
	var newSinucaEvent sinucaevent

	fmt.Println("Entrou no post:")
	fmt.Printf("JSON: %v\n", newSinucaEvent)

	// Call BindJSON to bind the received JSON to newSinucaEvent.
	if err := c.BindJSON(&newSinucaEvent); err != nil {
		return
	}

	// Add the new sinucaEvent to the slice.
	sinucaevents = append(sinucaevents, newSinucaEvent)
	c.IndentedJSON(http.StatusCreated, newSinucaEvent)
}

// getSinucaEventByID locates the sinucaEvent whose ID value matches the id
// parameter sent by the client, then returns that sinucaEvent as a response.
func getSinucaEventByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, a := range sinucaevents {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sinuca event not found"})
}

// getSinucaEventByUUID locates the sinucaEvent whose UUID value matches the id
// parameter sent by the client, then returns that sinucaEvent as a response.
func getSinucaEventByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	for _, a := range sinucaevents {
		if a.UUID == uuid {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sinuca event not found"})
}
