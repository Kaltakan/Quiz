package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type Questions struct {
	IdQuestion int       `json:"id"`
	Answer     int       `json:"answer"`
	Options    [4]string `json:"options"`
	Question   string    `json:"question"`
}

var getQuestionsCmd = &cobra.Command{
	Use:   "questions",
	Short: "Fetch quiz questions",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:8000/questions")
		if err != nil {
			fmt.Println("Error fetching questions:", err)
			return
		}
		defer resp.Body.Close()

		var questions []Questions
		if err := json.NewDecoder(resp.Body).Decode(&questions); err != nil {
			fmt.Println("Error decoding response:", err)
			return
		}

		for _, q := range questions {
			fmt.Printf("Q%d: %s\n", q.IdQuestion, q.Question)
			for i, opt := range q.Options {
				fmt.Printf("  %d. %s\n", i, opt)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getQuestionsCmd)
}
