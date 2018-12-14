package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
  Use: "shutto",
  Short: "Open a current branch github web page",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hey this is shutto")
  },
}
