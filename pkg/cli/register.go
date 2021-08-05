package cli

import (
	"errors"
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	Register = &cobra.Command{
		Use:           "register [plant] --schedule [cron] --target-pwa [moistureBars]",
		Short:         "creates a new plant schedule with a target moisture level in bars",
		Args:          cobra.ExactArgs(1),
		SilenceErrors: false,
		SilenceUsage:  false,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			log.Info("creating new schedule")

			shCommand := exec.Command("echo","TODO: register new IoT thing ", args[0])
			stdout, err := shCommand.Output()

			if err != nil {
				fmt.Println(err.Error())
				return errors.New(fmt.Sprint("IoT thing failed provisioning"))
			}

			// Print the output
			fmt.Println(string(stdout))
			return nil
		},
	}
)
