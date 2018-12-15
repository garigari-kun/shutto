package cmd

import (
  "fmt"
  "os"
  "os/exec"
  "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
  Use: "shutto",
  Short: "Open a current branch github web page",
  Run: func(cmd *cobra.Command, args []string) {
    out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
    if err != nil {
      os.Exit(1)
    }
    fmt.Println(string(out))
  },
}
