// +build experimental

package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/spf13/cobra"
)

// webserverCmd represents the webserver command -- mostly from Mat Ryer YouTube
var webserverCmd = &cobra.Command{
	Use:   "webserver",
	Short: "A web server",
	Long: `A web server

for you.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webserver called")
	},
}

func init() {
	xCmd.AddCommand(webserverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webserverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webserverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

  bindFlags(webserverCmd)
}
