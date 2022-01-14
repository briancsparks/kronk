// +build xcurr

package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"

	"github.com/spf13/cobra"
)

// imgandvidCmd represents the imgandvid command
var imgandvidCmd = &cobra.Command{
	Use:   "imgandvid",
	Short: "Image and Video",
	Long: `Image and Video

for you.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("imgandvid called")
	},
}

func init() {
	xcurrCmd.AddCommand(imgandvidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imgandvidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imgandvidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

  bindFlags(imgandvidCmd)
}
