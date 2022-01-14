// +build experimental

package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "github.com/spf13/viper"

  "github.com/spf13/cobra"
)

// flagsCmd represents the flags command
var flagsCmd = &cobra.Command{
	Use:   "flags",
	Short: "Play with and test flags",
	Long: `Play with and test

flags.`,

	Run: func(cmd *cobra.Command, args []string) {
    s := viper.GetString("String")
    s2 := viper.GetString("TwoPartString")
		fmt.Printf("flags called, string: %s / %s\n", s, s2)

    //viper.WriteConfigAs("./config.yaml")
	},
}

func init() {

  xCmd.AddCommand(flagsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flagsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
  flagsCmd.Flags().String("string", "default-string", "string usage")
  viper.BindPFlag("String", flagsCmd.Flags().Lookup("string"))

  //viper.SetEnvPrefix("KRONKFLAGS")
  viper.BindEnv("String")

  flagsCmd.Flags().String("two-part-string", "default-two-part-string", "two-part-string usage")
  viper.BindPFlag("TwoPartString", flagsCmd.Flags().Lookup("two-part-string"))

  //viper.SetEnvPrefix("KRONKFLAGS")
  viper.BindEnv("TwoPartString")

  flagsCmd.Flags().BoolP("bool", "b", false, "bool")

  bindFlags(flagsCmd)
}
