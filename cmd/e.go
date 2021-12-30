package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"

	"github.com/spf13/cobra"
)

// eCmd represents the e command
var eCmd = &cobra.Command{
	Use:   "e",
	Short: "Execute a command",
	Long: `Execute any command.

And Kronk will Kronk-ify it.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("executing %v\n", args)
	},
}

func init() {
	rootCmd.AddCommand(eCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// eCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// eCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
