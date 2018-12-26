package cmd

import (
  "log"
  "os"
  "os/exec"
  "github.com/garigari-kun/shutto/helper"
  "github.com/spf13/cobra"
  "github.com/skratchdot/open-golang/open"
)

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
      url := helper.Urlify(url_out)

      branch_out, err := exec.Command("git", "symbolic-ref", "--short", "HEAD").Output()
      if err != nil {
        log.Print(err)
        os.Exit(1)
      }
      branch_name := helper.FormatBrachName(branch_out)

      if branch_name == "master" {
        log.Print("It's master branch. No PR existed")
        os.Exit(1)
      } else {
        opening_url := url + "/pull/" + branch_name
        open.Run(opening_url)
      }

    },
  }

  return cobra
}
