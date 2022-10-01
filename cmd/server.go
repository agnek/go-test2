package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

var Cmd = &cobra.Command{
	Use:          "server",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Info().Interface("config", cfg).Msg("starting server")

		return nil
	},
}

func Execute() {
	err := Cmd.Execute()

	if err != nil {
		log.Err(err).Msg("process stopped with error. exit")
		os.Exit(1)
	}
}

func init() {
	Cmd.Flags().AddFlagSet(cfg.Flags())
}
