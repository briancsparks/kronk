// +build experimental

package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "net/http"

  "github.com/spf13/cobra"
)

// matryerCmd represents the matryer command
var matryerCmd = &cobra.Command{
	Use:   "matryer",
  Short: "A Mat Ryer web server",
  Long: `A Mat Ryer web server

for you.`,

  Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("matryer called")
	},
}

func init() {
	webserverCmd.AddCommand(matryerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// matryerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// matryerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

  bindFlags(matryerCmd)
}

type server struct {
  // Fields like:   db *someDatabase

  // TODO: router := gin.New()        See: https://github.com/googollee/go-socket.io/blob/master/_examples/gin-gonic/main.go
  // TODO: or use echo                     https://github.com/labstack/echo
  //router *someRouter
}

func newServer() *server {
  s := &server{}
  //s.routes()
  return s
}

// This is all that is needed to make an http.Handler
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  //s.router.serveHTTP(w,r)
}

// Routes -- put in routes.go, usually
func (s *server) routes() {
  //s.router.Get("/api/", s.handleAPI())
  //s.router.Get("/about", s.handleAbout())
  //s.router.Get("/", s.handleIndex())
}

// Handlers hang off the server
func (s *server) handleSomething() http.HandlerFunc {
  //thing := prepareThing()
  return func(w http.ResponseWriter, r *http.Request) {
    // Use thing
  }
}

// Example
func (s *server) handleGreeting(format string) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, format, r.FormValue("name"))
  }
}

func (s *server) routesExample() {
  //s.router.Get("/en", s.handleGreeting("Hello %s"))
  //s.router.Get("/es", s.handleGreeting("Hola %s"))
}

// Middleware Example
func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    //if !currentUser(r).IsAdmin {
    http.NotFound(w,r)
    return
    //}
    h(w,r)
  }
}

func (s *server) routesExampleWithMw() {
  //s.router.Get("/admin", s.adminOnly(s.handleAdminIndex()))
}

