# Notes

## Bootstrap GO CLI

### Git

```shell
git init
gi go >> .gitignore
gi goland >> .gitignore
gi jetbrains+all >> .gitignore
git add .
git commit -am first
```

### Go

```shell
go mod init 'github.com/briancsparks/kronk'
```

### Cobra

```shell
cobra init --license MIT --viper --pkg-name kronk
cobra add hammer
```

Fix `go` files:

```go
package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

```

### TODO: Fix Viper / Cobra integration

See kronk main.go and cmd/flags.go

Fix `cmd/hammer.go`. Add this to the end of `init()`:

```go
bindFlags(hammerCmd)
```

Fix `cmd/root.go`:

Replace this:

```go
viper.AutomaticEnv() // read in environment variables that match
```

With this (replacing `KRONK`):

```go
viper.SetEnvPrefix("KRONK")
viper.AutomaticEnv() // read in environment variables that match

bindFlags(rootCmd)
```

Add this to the bottom of `cmd/root.go`, replacing `KRONK`:

```go
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
```

### Finish

```shell
go mod tidy
go list -m all
git add .
git commit -am "cobra init"
```

### Beyond Init

```shell
go get github.com/abc/def

cobra add -p hammerCmd nail
```

