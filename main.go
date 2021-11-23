package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// player struct holds the player's name and nationality
type player struct {
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
}

// Test data
var players = []player{
	{Name: "Emile Smith-Rowe", Nationality: "English"},
	{Name: "Bakayo Saka", Nationality: "English"},
}

func main() {
	router := gin.Default()
	router.GET("/players", getPlayers)
	router.GET("/players/:name", getPlayerByName)

	router.POST("/players", postPlayers)

	router.Run("localhost:8080")
}

// postPlayers adds a new player to the list of players
func postPlayers(c *gin.Context) {
	var player player

	// Bind JSON to struct
	if err := c.BindJSON(&player); err != nil {
		return
	}

	// Add player to players slice
	players = append(players, player)
	c.IndentedJSON(http.StatusCreated, player)
}

// getPlayerByNationality responds with the list of players by nationailty
func getPlayerByNationality(c *gin.Context) {
	n := c.Params.ByName("nationality")
	for _, player := range players {
		if player.Nationality == n {
			c.IndentedJSON(http.StatusOK, player)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Nationality not found"})
}

// getPlayerByName responds with the player with the given name as JSON
func getPlayerByName(c *gin.Context) {
	name := c.Params.ByName("name")
	for _, player := range players {
		if player.Name == name {
			c.IndentedJSON(http.StatusOK, player)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Player not found"})
}

// getPlayers responds with the list of all players as JSON
func getPlayers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, players)
}
