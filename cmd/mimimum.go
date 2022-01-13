package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "github.com/spf13/viper"
  "log"
  "net/http"

  "github.com/spf13/cobra"
)

/*
 TODO:

 * Handle verbs (GET, POST, etc)
   * Get POST/PUT data
 * Parse JSON body
 * Parse Query params
 * Handle errors like 404, etc
 * Return JSON
 * Allow binding to IP address
 * Serve static files, dirs, embeds
 */

var port int

// mimimumCmd represents the mimimum command
var mimimumCmd = &cobra.Command{
  Use:   "minimum",
  Short: "A minimum web server",
  Long: `A minimum web server

for you.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mimimum called")

    // This handles EVERYTHING
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello!")
    })

    fmt.Printf("Starting server at port %d\n", port)
    if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
      log.Fatal(err)
    }
	},
}

func init() {
	webserverCmd.AddCommand(mimimumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mimimumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mimimumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
  mimimumCmd.Flags().IntVarP(&port, "port", "p", 8080, "The port to listen at")
  viper.BindPFlag("port", mimimumCmd.Flags().Lookup("port"))
  viper.BindEnv("port")

  bindFlags(mimimumCmd)
}
