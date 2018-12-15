package cmd

import (
  "log"
  "os"
  "os/exec"
  "strings"
  "github.com/spf13/cobra"
  "github.com/skratchdot/open-golang/open"
)

func init() {
  RootCmd.AddCommand(prCmd())
}

func prCmd() *cobra.Command {
  cobra := &cobra.Command{
    Use: "pr",
    Short: "Open a current branch's PR",
    Run: func(cmd *cobra.Command, args []string) {
      url_out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
      if err != nil {
        log.Print(err)
        os.Exit(1)
      }

      branch_out, err := exec.Command("git", "symbolic-ref", "--short", "HEAD").Output()
      if err != nil {
        log.Print(err)
        os.Exit(1)
      }

      if strings.TrimSpace(string(branch_out)) == "master" {
        log.Print("It's master branch. No PR existed")
        os.Exit(1)
      } else {
        opening_url := string(url_out[:len(string(url_out)) - 5]) + "/pull/" + strings.TrimSpace(string(branch_out))
        open.Run(string(opening_url))
      }

    },
  }

  return cobra
}
