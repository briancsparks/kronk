// +build experimental

package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  socketio "github.com/googollee/go-socket.io"
  "github.com/spf13/cobra"
  "log"
  "net/http"
)

// See this for clean shutdown: https://github.com/googollee/go-socket.io/blob/master/_examples/graceful-shutdown/main.go
// See this for redis:          https://github.com/googollee/go-socket.io/tree/master/_examples/redis-adapter

// socketioCmd represents the socketio command
var socketioCmd = &cobra.Command{
	Use:   "socketio",
	Short: "MVP Socket-IO server in Go",
	Long: `MVP Socket-IO server in Go

for you.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("socketio called")

    server := socketio.NewServer(nil)

    server.OnConnect("/", func(s socketio.Conn) error {
      s.SetContext("")
      fmt.Println("connected:", s.ID())
      return nil
    })

    server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
      fmt.Println("notice:", msg)
      s.Emit("reply", "have "+msg)
    })

    server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
      s.SetContext(msg)
      return "recv " + msg
    })

    server.OnEvent("/", "bye", func(s socketio.Conn) string {
      last := s.Context().(string)
      s.Emit("bye", last)
      s.Close()
      return last
    })

    server.OnError("/", func(s socketio.Conn, e error) {
      fmt.Println("meet error:", e)
    })

    server.OnDisconnect("/", func(s socketio.Conn, reason string) {
      fmt.Println("closed", reason)
    })

    go server.Serve()
    defer server.Close()

    http.Handle("/socket.io/", server)
    http.Handle("/", http.FileServer(http.Dir("./cmd/socketio-assets")))      // TODO: embed
    log.Println("Serving at localhost:8123...")
    log.Fatal(http.ListenAndServe(":8123", nil))
  },
}

func init() {
	xCmd.AddCommand(socketioCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// socketioCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// socketioCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

  //socketioCmd.Flags().IntVarP(&port, "port", "p", 8080, "The port to listen at")
  //viper.BindPFlag("port", socketioCmd.Flags().Lookup("port"))
  //viper.BindEnv("port")

  bindFlags(socketioCmd)
}
