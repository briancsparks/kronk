// +build robotgo

package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/spf13/cobra"
)

// uiCmd represents the ui command
var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Kronk automate UI",
	Long: `Kronk automate UI

For you.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ui called")
	},
}

func init() {
	rootCmd.AddCommand(uiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
  //uiCmd.Flags().Bool("all", false, "Show all repos (even those without changes.)")
  //viper.BindPFlag("all", uiCmd.Flags().Lookup("all"))
  //viper.BindEnv("all")

  bindFlags(uiCmd)
}
