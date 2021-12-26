package cmd

import (
  "bytes"
  "fmt"
  "log"
  "os/exec"
  "strings"
  "syscall"
  "time"
)

// TODO: Throttle the number of processes launched

func launch4Result(exename string , args []string) (chan string, error) {
  return launchForResult(exename, args, "", "")
}


func launchForResult(exename string, args []string, cwd string, deffault string) (chan string, error) {
  out := make(chan string)

  exepath, err := exec.LookPath(exename)
  if err != nil {
    return nil, err
  }
  //fmt.Println(exepath)

  go func() {
    defer close(out)

    start := time.Now()
    res := deffault
    cmd := exec.Command(exepath, args...)

    if len(cwd) > 0 {
      cmd.Dir = cwd
    }

    var stdout bytes.Buffer
    cmd.Stdout = &stdout

    //var stderr bytes.Buffer
    //cmd.Stderr = &stderr

    if err = cmd.Start(); err != nil {
      log.Panic(err)                          /* probably shouldn't exit */
    }

    if err = cmd.Wait(); err == nil {
      res = strings.TrimSpace(stdout.String())

    } else {
      // See if the cmd exited with code != 0
      if exitErr, isExitError := err.(*exec.ExitError); isExitError {
        if status, isWaitStatus := exitErr.Sys().(syscall.WaitStatus); isWaitStatus {
          Vvvverbose(fmt.Sprintf("  -- ExitFor(%s): %d\n", cmd.Dir, status.ExitStatus()))

          res = deffault
        }
      }
      //log.Panic(err)                          /* probably shouldn't exit */
    }

    Verbose(fmt.Sprintf("  ----- [%4d] launch: %s> %s, %v ==> %s\n", time.Since(start).Milliseconds(), cwd, exepath, args, res))
    out <- res
  }()

  return out, err
}



