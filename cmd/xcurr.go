// +build xcurr

package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"

	"github.com/spf13/cobra"
)

// xcurrCmd represents the xcurr command
var xcurrCmd = &cobra.Command{
	Use:   "xcurr",
	Short: "The parent of all current experimentals",
	Long: `The parent of all current experimentals

4u.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("xcurr called")
	},
}

func init() {
	rootCmd.AddCommand(xcurrCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// xcurrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// xcurrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
