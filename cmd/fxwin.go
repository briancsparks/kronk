package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/spf13/cobra"
)

// fxwinCmd represents the fxwin command
var fxwinCmd = &cobra.Command{
	Use:   "fxwin",
	Short: "Kronk fix a window",
	Long: `Kronk fix a window

For you.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fxwin called")
	},
}

func init() {
	uiCmd.AddCommand(fxwinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fxwinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
  //fxwinCmd.Flags().Bool("all", false, "Show all repos (even those without changes.)")
  //viper.BindPFlag("all", fxwinCmd.Flags().Lookup("all"))
  //viper.BindEnv("all")

  bindFlags(fxwinCmd)
}
