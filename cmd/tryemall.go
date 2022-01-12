// +build robotgo

package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "github.com/go-vgo/robotgo"
  "github.com/spf13/viper"

  "github.com/spf13/cobra"
)

// tryemallCmd represents the tryemall command
var tryemallCmd = &cobra.Command{
	Use:   "tryemall",
	Short: "Try all the robotgo stuff",
	Long: `Kronk try all the robotgo stuff

For you.`,

	Run: func(cmd *cobra.Command, args []string) {
    nameToFind := viper.GetString("pname")
		fmt.Printf("Try 'Em All\n")

    // Name -> PIDs
    fpid, err := robotgo.FindIds(nameToFind)
    if err == nil {
      fmt.Printf("%s -> %v\n", nameToFind, fpid)
    }

    // PID -> Name
    pidToFind := viper.GetInt("pid")
    foundName, err := robotgo.FindName(int32(pidToFind))
    if err == nil {
      fmt.Printf("Pid: %d -> %s\n", pidToFind, foundName)
    }

    // All names
    names, err := robotgo.FindNames()
    if err == nil {
      fmt.Printf("All names: %v\n", names)

//NamesLoop:
//      for _, name := range names {
//        foundPids, err := robotgo.FindIds(name)
//        if err != nil {
//          fmt.Printf("Error finding PIDs: %v\n", err)
//          continue NamesLoop
//        }
//
//        for _, pid := range foundPids {
//          foundPath, err := robotgo.FindPath(pid)
//          if err != nil {
//            fmt.Printf("Error finding path for pid %d -- %v\n", pid, err)
//            continue NamesLoop
//          }
//
//          fmt.Printf("Pid %d -> Path: %s\n", pid, foundPath)
//
//
//          foundName, err = robotgo.FindName(pid)
//          if err != nil {
//            fmt.Printf("Error finding name for pid %d -- %v\n", pid, err)
//            continue NamesLoop
//          }
//
//          fmt.Printf("Pid %d -> Name: %s\n", pid, foundName)
//        }
//      }
    }

    //mdata := robotgo.GetActive()


    //robotgo.ActiveName(nameToFind)
    //title := robotgo.GetTitle()
    //fmt.Println("title@@@ ", title)
	},
}

func init() {
	uiCmd.AddCommand(tryemallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tryemallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

  mkCliArgString(tryemallCmd, "pname", "Google", "Name of process to fetch its PID.")
  mkCliArgInt(tryemallCmd, "pid", 1000, "PID of process to find.")

  bindFlags(tryemallCmd)
}

func mkCliArgString(cmd *cobra.Command, name, deffault, usage string) {
  cmd.Flags().String(name, deffault, usage)
  viper.BindPFlag(name, cmd.Flags().Lookup(name))
  viper.BindEnv(name)
}

func mkCliArgBool(cmd *cobra.Command, name string, deffault bool, usage string) {
  cmd.Flags().Bool(name, deffault, usage)
  viper.BindPFlag(name, cmd.Flags().Lookup(name))
  viper.BindEnv(name)
}

func mkCliArgInt(cmd *cobra.Command, name string, deffault int, usage string) {
  cmd.Flags().Int(name, deffault, usage)
  viper.BindPFlag(name, cmd.Flags().Lookup(name))
  viper.BindEnv(name)
}


