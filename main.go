package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// player struct holds the player's name and nationality
type player struct {
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Goals       goals  `json:"goals"`
}

type goals struct {
	Scored      int `json:"scored"`
	Headed      int `json:"headed"`
	Volleyed    int `json:"volleyed"`
	LeftFooted  int `json:"leftfooted"`
	RightFooted int `json:"rightfooted"`
	OwnGoals    int `json:"owngoals"`
}

// Test data
var players = []player{
	{
		Name:        "Emile Smith-Rowe",
		Nationality: "English",
		Goals: goals{
			Scored:   12,
			Headed:   1,
			Volleyed: 1,
		},
	},
	{Name: "Bakayo Saka", Nationality: "English"},
}

func main() {
	r := setupRouter()
	r.Run("localhost:8080")
}

// setupRouter sets up the router and returns it
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Top 100 players under 23",
		},
		)
	})

	router.GET("/players", getPlayers)
	router.GET("/players/:name", getPlayerByName)

	router.POST("/players", postPlayers)

	return router
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
