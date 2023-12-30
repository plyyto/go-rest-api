package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/teams", getTeams)

	router.Run("localhost:8080")
}

type team struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

var teams = []team{
	{Name: "Liverpool", Points: 42},
	{Name: "Arsenal", Points: 40},
	{Name: "Aston Villa", Points: 39},
	{Name: "Manchester City", Points: 37},
	{Name: "Tottenham", Points: 36},
}

func getTeams(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, teams)
}
