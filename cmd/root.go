package cmd

import (
  "log"
  "os"
  "os/exec"
  "github.com/spf13/cobra"
  "github.com/skratchdot/open-golang/open"
)

var RootCmd = &cobra.Command{
  Use: "shutto",
  Short: "Open a github repository page",
  Run: func(cmd *cobra.Command, args []string) {
    out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
    if err != nil {
      log.Print(err)
      os.Exit(1)
    }
    url := string(out[:len(string(out)) - 5])
    open.Run(url)
  },
}

func Execute() {
  if err := RootCmd.Execute(); err != nil {
    log.Print(err)
    os.Exit(1)
  }
}
