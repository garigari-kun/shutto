package cmd

import (
	"github.com/garigari-kun/shutto/helper"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

func branchCmd() *cobra.Command {
	cobra := &cobra.Command{
		Use:   "br",
		Short: "Open a current branch github page",
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
				open.Run(url)
			} else {
				opening_url := url + "/tree/" + branch_name
				open.Run(opening_url)
			}
		},
	}

	return cobra
}
