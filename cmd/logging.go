package cmd

import (
  "github.com/spf13/viper"
  "log"
)

func Verbose0(s string) {
  log.Printf("%s", s)
}

func Verbose(s string) {
  Kerbose(s)
}

func Vverbose(s string) {
  Kkerbose(s)
}

func Vvverbose(s string) {
  Kkkerbose(s)
}

func Vvvverbose(s string) {
  Kkkkerbose(s)
}

func Verprint(s string) {
  log.Printf("%s", s)
}




func Kerbose(s string) {
  if viper.GetBool("kerbose") || viper.GetBool("verbose") {
    log.Printf("%s", s)
  }
}

func Kkerbose(s string) {
  if viper.GetBool("kkerbose") || viper.GetBool("vverbose") {
    log.Printf("%s", s)
  }
}

func Kkkerbose(s string) {
  if viper.GetBool("kkkerbose") || viper.GetBool("vvverbose") {
    log.Printf("%s", s)
  }
}

func Kkkkerbose(s string) {
  if viper.GetBool("kkkkerbose") || viper.GetBool("vvvverbose") {
    log.Printf("%s", s)
  }
}


