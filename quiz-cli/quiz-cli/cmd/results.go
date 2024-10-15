package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var resultsCmd = &cobra.Command{
	Use:   "results",
	Short: "Get your results and compare with others",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:8000/results")
		if err != nil {
			fmt.Println("Error fetching results:", err)
			return
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			fmt.Println(json.NewDecoder(resp.Body))
			fmt.Println("Error decoding response:", err)
			return
		}

		fmt.Printf("Your latest score: %d\n", int(result["latest_score"].(float64)))
		fmt.Printf("You were better than %.2f%% of all quizzers\n", result["percentage"].(float64))
	},
}

func init() {
	rootCmd.AddCommand(resultsCmd)
}
