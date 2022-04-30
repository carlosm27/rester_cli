/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/carlosm27/rester_cli/httpMethods"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rester_cli",
	Short: "A REST client cli",
	Long:  `A REST Client cli that support the methods GET, POST, DELETE and POST`,

	Run: func(cmd *cobra.Command, args []string) {
		httpMethod := ""
		prompt := &survey.Select{
			Message: "Choose a HTTP Method:",
			Options: []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		}
		survey.AskOne(prompt, &httpMethod)

		switch httpMethod {
		case "GET":
			url := ""
			prompt := &survey.Input{
				Message: "Enter URL:",
			}
			survey.AskOne(prompt, &url)

			get, err := httpMethods.Get(url)

			if err != nil {
				log.Println(err)
			}
			fmt.Println(get)

		case "POST":
			url := ""
			prompt := &survey.Input{
				Message: "Enter URL:",
			}
			survey.AskOne(prompt, &url)
			body := ""
			prompt = &survey.Input{
				Message: "Enter Body:",
			}
			survey.AskOne(prompt, &body)

			post, err := httpMethods.Post(url, body)
			if err != nil {
				log.Println(err)
			}

			fmt.Println(post)

		case "PATCH":
			url := ""
			prompt := &survey.Input{
				Message: "Enter URL:",
			}
			survey.AskOne(prompt, &url)
			body := ""
			prompt = &survey.Input{
				Message: "Enter Body:",
			}
			survey.AskOne(prompt, &body)

			patch, err := httpMethods.Patch(url, body)
			if err != nil {
				log.Println(err)
			}

			fmt.Println(patch)

		case "PUT":
			url := ""
			prompt := &survey.Input{
				Message: "Enter URL:",
			}
			survey.AskOne(prompt, &url)
			body := ""
			prompt = &survey.Input{
				Message: "Enter Body:",
			}
			survey.AskOne(prompt, &body)

			put, err := httpMethods.Put(url, body)
			if err != nil {
				log.Println(err)
			}

			fmt.Println(put)

		case "DELETE":
			url := ""
			prompt := &survey.Input{
				Message: "Enter URL:",
			}
			survey.AskOne(prompt, &url)

			delete, err := httpMethods.Delete(url)
			if err != nil {
				log.Println(err)
			}

			fmt.Println(delete)
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rester_cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
