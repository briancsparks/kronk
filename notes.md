# Notes

## Fixes to use `robotgo`

### Fix Building from Cygwin CLI

Essentially, need to have `go` installed with `cgo`, and the 64-bit C
compiler. Neither the go installer nor GoLand setup `cgo` (I thought GoLand did).

* Install TDM GCC64 - https://github.com/lowkey42/MagnumOpus/wiki/TDM-GCC-Mingw64-Installation
  * Put it into your `PATH` (`C:\TDM-GCC-64\bin`)
  * Make sure it is the first `gcc` in your `PATH`.
  * I had cygwin installed, which put its gcc at the very top of the path, so
    I was unsuccessful.

That much gets 64-bit `cgo` for you. However, robotgo also needed zlib installed.

* From the same TDM url is a link to get the 64-bit zlib libraries and headers.
  * ZIP/bin/*      ->  TDM/bin/
  * ZIP/lib/*      ->  TDM/lib/
  * ZIP/include/*  ->  TDM/include/

At this point, I was able to build against robotgo from the cygwin shell.

### Fix Building from GoLand

I got this error when trying to compile from GoLand.

```
# github.com/robotn/gohook
cc1.exe: sorry, unimplemented: 64-bit mode not compiled in
```

I could not figure out how to tell GoLand to use the TDM GCC compiler. The
go compiler settings are at File > Settings | Go > Build Tags & Vendoring.

* TBD


## Screen Saver (Blanking) on Pi

```shell
# To see the X settings:
xset q

# To change the blanking timeout (blank after an hour):
xset s 3600 3600
xset dpms 3600 3600 3600
```

## Putting Keys on Servers

```shell
# To see which keys:
aws s3 ls s3://briancsparks-files/keys/

# Get a specific key into the authorized_keys file
curl -L https://briancsparks-files.s3.amazonaws.com/keys/id_rsa.pub >> ${HOME}/.ssh/authorized_keys
```

## Dev Server

In order to use CLion remote, must setup the server for development. The only
thing that was difficultish was to install CMake 3.20, the version in the repos
was something earlier.

General dev-server setup (for C/C++):

```shell
#sudo apt-get install -y cmake gcc gdb tree curl git htop libssl-dev build-essential
sudo apt-get install -y gcc gdb tree curl git htop libssl-dev build-essential
```

Upgrading to CMake 3.20

```shell
#cmake -version
#sudo apt-get remove -y cmake
#sudo apt autoremove

curl -L 'https://github.com/Kitware/CMake/releases/download/v3.20.0/cmake-3.20.0.tar.gz' | tar -zxv
cd cmake-3.20.0/
./bootstrap
make
sudo make install
```








