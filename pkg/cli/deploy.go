package cli

import (
	"errors"
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	Deploy = &cobra.Command{
		Use:           "deploy --profile [aws-profile]",
		Short:         "deploys base infrastructure to the target profile",
		Args:          cobra.ExactArgs(1),
		SilenceErrors: false,
		SilenceUsage:  false,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			log.Info("deploying core infrastructure")

			shCommand := exec.Command("terraform","apply", "--auto-approve", "--profile", args[0])
			stdout, err := shCommand.Output()

			if err != nil {
				fmt.Println(err.Error())
				return errors.New(fmt.Sprint("terraform failed to deploy"))
			}

			// Print the output
			fmt.Println(string(stdout))
			return nil
		},
	}
)
