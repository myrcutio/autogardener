package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/myrcutio/autogardener/pkg/cli"
)

var (
	autogardener = &cobra.Command{
		Use:   "autogardener",
		Short: "Auto Gardener Command Line Interface",
	}
)

func init() {
	autogardener.AddCommand(cli.Deploy)

	var scheduleFlag string
	var pawFlag float32
	cli.Register.Flags().StringVarP(&scheduleFlag, "schedule", "s", "daily", "defines how often to check water levels")
	cli.Register.Flags().Float32VarP(&pawFlag, "target-paw", "t", .3, "target moisture level to try to maintain")
	autogardener.AddCommand(cli.Register)
}

func main() {
	if err := autogardener.Execute(); err != nil {
		log.Fatal(err)
	}
}