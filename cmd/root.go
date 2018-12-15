package cmd

import (
  "os"
  "os/exec"
  "github.com/spf13/cobra"
  "github.com/skratchdot/open-golang/open"
)

var RootCmd = &cobra.Command{
  Use: "shutto",
  Short: "Open a current branch github web page",
  Run: func(cmd *cobra.Command, args []string) {
    out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
    if err != nil {
      os.Exit(1)
    }
    open.Run(string(out))
  },
}
