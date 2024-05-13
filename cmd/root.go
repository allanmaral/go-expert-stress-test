package cmd

import (
	"fmt"
	"os"

	"github.com/allanmaral/go-expert-stree-test-challenge/internal/stresstest"
	"github.com/spf13/cobra"
)

var (
	url         string
	requests    int
	concurrency int
)

var rootCmd = &cobra.Command{
	Use:   "stress",
	Short: "Stress is a simple stress test CLI tools",
	Run: func(cmd *cobra.Command, args []string) {
		if requests < concurrency {
			fmt.Print("Error: number of requests must be greater then the number of concurrent requests\n\n")
			cmd.Help()
			os.Exit(1)
		}
		if url == "" {
			fmt.Print("Error: url is required\n\n")
			cmd.Help()
			os.Exit(1)
		}

		stressTest(url, requests, concurrency)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "url to run the stress test")
	rootCmd.PersistentFlags().IntVarP(&requests, "requests", "r", 10, "number of requests to run")
	rootCmd.PersistentFlags().IntVarP(&concurrency, "concurrency", "c", 1, "number of concurrent request")
	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("requests")
}

func stressTest(url string, requests int, concurrency int) {
	fmt.Printf("Stress testing %s with %d request(s) using %d concurrent requests\n", url, requests, concurrency)

	reporter := stresstest.NewReporter()
	tester := stresstest.NewTester(url, requests, concurrency)

	go tester.Run()
	report := reporter.Collect(tester.Results())

	fmt.Println(report.String())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
