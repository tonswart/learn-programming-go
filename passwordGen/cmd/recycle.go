/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// recycleCmd represents the recycle command
var recycleCmd = &cobra.Command{
	Use:   "recycle",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(recycleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recycleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recycleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	var myJoke Joke

	url := "https://icanhazdadjoke.com"
	responseBytes := getJokeData(url)
	err := json.Unmarshal(responseBytes, &myJoke)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(myJoke.ID))
	fmt.Println(string(myJoke.Joke))
	fmt.Println(int(myJoke.Status))
}

func getJokeData(basAPI string) []byte {
	request, err := http.NewRequest(http.MethodGet, basAPI, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return responseBytes
}
