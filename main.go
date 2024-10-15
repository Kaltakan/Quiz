package main

import (
	"math"
	"net/http"

	"sync"

	"github.com/gin-gonic/gin"
)

type Questions struct {
	IdQuestion int       `json:"id"`
	Answer     int       `json:"answer"`
	Options    [4]string `json:"options"`
	Question   string    `json:"question"`
}

type Quiz struct {
	questions []Questions
	Mutex     sync.Mutex
	Results   []int
}

var quiz = Quiz{
	questions: []Questions{
		{IdQuestion: 1, Answer: 1, Question: "What's the color of Napoleon's white horse?", Options: [4]string{"White", "Red", "Blue", "Orange"}},
		{IdQuestion: 2, Answer: 3, Question: "What's the capital of Spain?", Options: [4]string{"Rome", "London", "Madrid", "Barcellona"}},
	},
	Results: []int{},
}

func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, quiz.questions)
}

// postAlbums adds an album from JSON received in the request body.
func handleSubmit(c *gin.Context) {
	var score int
	var answers map[int]int
	quiz.Mutex.Lock()
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&answers); err != nil {
		return
	}

	for id, answer := range answers {
		for _, q := range quiz.questions {
			if id == q.IdQuestion && answer == q.Answer {
				score++
			}
		}
	}

	quiz.Results = append(quiz.Results, score)
	quiz.Mutex.Unlock()

	c.IndentedJSON(http.StatusCreated, score)
}

func handleResults(c *gin.Context) {
	quiz.Mutex.Lock()
	defer quiz.Mutex.Unlock()

	if len(quiz.Results) > 0 {
		lastScore := quiz.Results[len(quiz.Results)-1]
		percentage := calculatePercentage(lastScore, quiz.Results)
		result := map[string]interface{}{
			"latest_score": lastScore,
			"percentage":   percentage,
		}

		c.IndentedJSON(http.StatusCreated, result)
	} else {
		c.IndentedJSON(http.StatusBadRequest, "message: Results not available")
	}

}

func calculatePercentage(score int, result []int) float64 {
	var count int

	for r := range quiz.Results {
		if score < r {
			count++
		}
	}

	return math.Round((float64(count) / float64(len(result))) * 100)
}

func main() {
	router := gin.Default()
	router.GET("/questions", getQuestions)
	router.POST("/submit", handleSubmit)
	router.GET("/results", handleResults)
	router.Run("localhost:8000")
}
