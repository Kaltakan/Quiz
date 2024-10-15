package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var submitAnswersCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit quiz answers",
	Run: func(cmd *cobra.Command, args []string) {
		answers := make(map[int]int)
		// Mock answers for demonstration. You can prompt for input here.
		answers[1] = 2
		answers[2] = 1

		body, err := json.Marshal(answers)
		if err != nil {
			fmt.Println("Error encoding answers:", err)
			return
		}

		resp, err := http.Post("http://localhost:8000/submit", "application/json", bytes.NewBuffer(body))
		if err != nil {
			fmt.Println("Error submitting answers:", err)
			return
		}
		defer resp.Body.Close()

		var result int
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			fmt.Println("Error decoding response:", err)
			return
		}

		fmt.Printf("You scored: %d\n", result)
	},
}

func init() {
	rootCmd.AddCommand(submitAnswersCmd)
}
