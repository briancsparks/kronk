// +build xcurr

package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"

	"github.com/spf13/cobra"
)

// framesCmd represents the frames command
var framesCmd = &cobra.Command{
	Use:   "frames",
	Short: "Pull frames out of a video",
	Long: `Pull frames out of a video

So you can then evaluate the sub-frames. Make sure not to
have Dropbox running. Docs are at:
    https://wiki.videolan.org/VLC_command-line_help

    TBD`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("frames called")

    // See https://wiki.videolan.org/VLC_command-line_help
    // vlc.exe "C:\Users\sparksb\Desktop\camera-scans\cards-auto-fps.mp4" --video-filter=scene --scene-format=png --scene-prefix=cards-30fps --scene-ratio=10 --scene-path="C:\Users\sparksb\Desktop\camera-scans\cards-30fps-iii" --verbose=2 vlc://quit
    // vlc.exe "C:\path\to\video-file.mp4" --video-filter=scene --scene-format=png --scene-prefix=[[image-file-prefix]] --scene-ratio=[[ratio-of-frames (10)]] --scene-path="C:\output\path\" --verbose=2 vlc://quit
	},
}

func init() {
	imgandvidCmd.AddCommand(framesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// framesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// framesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

  bindFlags(framesCmd)
}
