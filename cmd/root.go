package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
	"github.com/spf13/cobra"
  "github.com/spf13/pflag"
  "os"
  "strings"

  homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var IsVerbose bool
var IsVverbose bool
var IsVvverbose bool
var IsVvvverbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kronk",

	Short: "The path that rocks",
	Long: `Hey, did you see that sky today -- Talk about blue!`,

  //Run: func(cmd *cobra.Command, args []string) {
  //  fmt.Printf("args: %v\n", args)
  //},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kronk.yaml)")

  rootCmd.PersistentFlags().BoolVarP(&IsVerbose, "verbose", "v", false, "Verbose output")
  viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

  rootCmd.PersistentFlags().BoolVarP(&IsVverbose, "vverbose", "", false, "Vverbose output")
  viper.BindPFlag("vverbose", rootCmd.PersistentFlags().Lookup("vverbose"))

  rootCmd.PersistentFlags().BoolVarP(&IsVvverbose, "vvverbose", "", false, "Vvverbose output")
  viper.BindPFlag("vvverbose", rootCmd.PersistentFlags().Lookup("vvverbose"))

  rootCmd.PersistentFlags().BoolVarP(&IsVvvverbose, "vvvverbose", "", false, "Vvvverbose output")
  viper.BindPFlag("vvvverbose", rootCmd.PersistentFlags().Lookup("vvvverbose"))

  // Kerbose
  rootCmd.PersistentFlags().Bool("kerbose", false, "Kerbose output")
  viper.BindPFlag("kerbose", rootCmd.PersistentFlags().Lookup("kerbose"))
  viper.BindEnv("kerbose")

  rootCmd.PersistentFlags().Bool("kkerbose", false, "Kkerbose output")
  viper.BindPFlag("kkerbose", rootCmd.PersistentFlags().Lookup("kkerbose"))
  viper.BindEnv("kkerbose")

  rootCmd.PersistentFlags().Bool("kkkerbose", false, "Kkkerbose output")
  viper.BindPFlag("kkkerbose", rootCmd.PersistentFlags().Lookup("kkkerbose"))
  viper.BindEnv("kkkerbose")

  rootCmd.PersistentFlags().Bool("kkkkerbose", false, "Kkkkerbose output")
  viper.BindPFlag("kkkkerbose", rootCmd.PersistentFlags().Lookup("kkkkerbose"))
  viper.BindEnv("kkkkerbose")



  // Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kronk" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kronk")
	}

  viper.SetEnvPrefix("KRONK")
	viper.AutomaticEnv() // read in environment variables that match

  bindFlags(rootCmd)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}


// Stolen from: https://github.com/carolynvs/stingoftheviper/blob/main/main.go
// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command /*, v *viper.Viper*/) {
  cmd.Flags().VisitAll(func(f *pflag.Flag) {

    // Environment variables can't have dashes in them, so bind them to their equivalent
    // keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
    if strings.Contains(f.Name, "-") {
      envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
      viper.BindEnv(f.Name, fmt.Sprintf("%s_%s", "KRONK", envVarSuffix))
    }

    // Apply the viper config value to the flag when the flag is not set and viper has a value
    if !f.Changed && viper.IsSet(f.Name) {
      val := viper.Get(f.Name)
      cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
    }
  })
}

func rootInitForSub() {
  if IsVvvverbose {
    IsVvverbose = true
    IsVverbose = true
    IsVerbose = true
  }

  if IsVvverbose {
    IsVverbose = true
    IsVerbose = true
  }

  if IsVverbose {
    IsVerbose = true
  }
}
