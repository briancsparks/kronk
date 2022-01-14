package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "github.com/spf13/viper"

  "github.com/spf13/cobra"
)

// notesCmd represents the notes command
var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "A collection of notes",
	Long: `A collection of notes

for you.`,

	Run: func(cmd *cobra.Command, args []string) {
    showNotes(false)
	},
}

func init() {
	rootCmd.AddCommand(notesCmd)

  notesCmd.Flags().BoolP("all", "a", false, "show all")
  viper.BindPFlag("all", notesCmd.Flags().Lookup("all"))

  notesCmd.Flags().Bool("vlc", false, "vlc notes")
  viper.BindPFlag("vlc", notesCmd.Flags().Lookup("vlc"))

  bindFlags(notesCmd)
}

func showNotes(forceAll bool) {
  count := 0

  // -------------------------------------------------------------------------------------------------------------------
  // vlc
  if forceAll || viper.GetBool("all") || viper.GetBool("vlc") {
    count += 1
    fmt.Println(`
===================
     VLC`)

    fmt.Println(`
-------------------
To capture frames:

  vlc.exe "C:\path\to\video-file.mp4" --video-filter=scene --scene-format=png --scene-prefix=[[image-file-prefix]] --scene-ratio=[[ratio-of-frames (10)]] --scene-path="C:\output\path\" --verbose=2 vlc://quit`)

  }

  
  // -------------------------------------------------------------------------------------------------------------------
  // -------------------------------------------------------------------------------------------------------------------
  if count == 0 {
    showNotes(true)
  } else {
    fmt.Println("")
  }
}
