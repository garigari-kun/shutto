package main

import (
  "log"
  "os"
  "github.com/garigari-kun/shutto/cmd"
)

func main() {
  if err := cmd.RootCmd.Execute(); err != nil {
    log.Print(err)
    os.Exit(1)
  }
}
