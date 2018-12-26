package cmd

import (
  "log"
  "os"
  "os/exec"
  "github.com/garigari-kun/shutto/helper"
  "github.com/spf13/cobra"
  "github.com/skratchdot/open-golang/open"
)

func RootCmd() *cobra.Command {
  cmd := &cobra.Command {
    Use: "shutto",
    Short: "Open a github repository page",
    Run: func(cmd *cobra.Command, args []string) {
     out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()

     if err != nil {
       log.Print(err)
       os.Exit(1)
     }

     url := helper.Urlify(out)
     open.Run(url)
    },
 }

 cmd.AddCommand(branchCmd())
 cmd.AddCommand(prCmd())
 return cmd
}

func Execute() {
  cmd := RootCmd()
  if err := cmd.Execute(); err != nil {
    log.Print(err)
    os.Exit(1)
  }
}
