# TODO

## To Kronkify

Go Tech to put into Kronk

The list of things that are started:

* Webserver
  * built-in HTTP server
  * Mimimal minimum server -- `./kronk.exe webserver minimum`
  * Mat Ryer -- Typed in what was shown in video. Does not work.

Todo

* Projects
  * Pawns dont listen
  * Data Science

Technology

* Techs
  * Fyne
  * Ebiten SVG
  * Other Go plotting
  * io/fs
  * afero
  * TLS
  * Lets Encrypt
  * websockets
  * socket.io
  * Gorilla
  * Gin?
  * GraphQL
  * UDP Server
  * Lecture RSS feed
  * Git-go
  * Mongo
  * Redis
  * OpenCV

Categorized

* Automate UI / AI Game Driver / Screen Colors
  * RobotGo
* Automate Tasks (Kronkifications)
  * Daemonize
  * Files watcher
  * Win user acct APIs
* Cloud
  * AWS
  * GCP
* Go UI
  * Lorca
    * An FS or two (for Lorca generation)
* DebugViz
  * Lorca
* Games
  * Ebiten
  * Some 3D
  * Low-level keyboard for VIM-Platformer
* X-Join / Config
  * Otto
  * sgfs-like dir awareness
* Brash
  * An FS or two (in prep of brash)
  * TTY / PTY (in prep of brash)
  * sgfs-like dir awareness
  * Low-level keyboard for VIM-Platformer
* cgo
* ~~Dir Walk~~
* Win user acct APIs
  * https://github.com/iamacarpet/go-win64api

## General To Dos

* Remove cmd/runner.go
* Figure out how ex/runner.go can get to viper
* Re-enable ex/runner.go Verbose()-ability

## gitsstatus

* ~~Don't recurse AppData~~
* Be smart about AppData, and all the %AppWhatever% vars.

## tree

A tree utility

* Automagically know where not to recurse.
  * `node_modules`, `.git`, etc
  * Maybe in git submodules
* Put files horizontally in certain scenarios
  * Like when the dir has no sub-dirs, but plenty of files

## Other

### Pty / TTY

* `stty -a`
* https://pkg.go.dev/github.com/creack/pty
* https://dev.to/napicella/linux-terminals-tty-pty-and-shell-part-2-2cb2

### ImageMagick Fun

* Multiple PNGs to multi-page PDF

```shell
convert  page1.jpg page2.jpg page3.jpg -rotate 90 output.pdf
```

Force output type (jpeg):
previous_stage_in_pipeline | convert -filteroptions jpeg:- | next_stage_in_pipline

### Icon Maker

* Android-project-logo
* ICOs for tray
* favicon
* See https://legacy.imagemagick.org/Usage/advanced/

### Synergy-like

* Scroll the other computer.

### Brash

* See the dropbox cli `/cygdrive/d/data/projects/GolandProjects/from-github/dropbox/dbxcli`

### ????

* git status in prompt

### DNS

* Set A record

## The WSL2 Make-Right-er Tool

* Where did those scripts go?


