// +build robotgo

package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/go-vgo/robotgo"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

// fxwinCmd represents the fxwin command
var fxwinCmd = &cobra.Command{
	Use:   "fxwin",
	Short: "Kronk fix a window",
	Long: `Kronk fix a window

For you.`,

	Run: func(cmd *cobra.Command, args []string) {
    pname := viper.GetString("pname")
		fmt.Printf("fxwin called, looking for %s\n", pname)

    // PIDs
    fpid, err := robotgo.FindIds("Google")
    if err == nil {
     fmt.Println("pids...", fpid)
    }

    robotgo.ActiveName(pname)
    title := robotgo.GetTitle()
    fmt.Println("title@@@ ", title)
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
  fxwinCmd.Flags().String("pname", "Google", "Name of process to look for.")
  viper.BindPFlag("pname", fxwinCmd.Flags().Lookup("pname"))
  viper.BindEnv("pname")

  bindFlags(fxwinCmd)
}
