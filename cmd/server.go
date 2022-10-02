package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go-test2/internal/app"
	"os"
)

var Cmd = &cobra.Command{
	Use:          "server",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Info().Interface("config", app.Cfg).Msg("starting server")

		b := app.NewBank()
		return app.StartServer(b, app.Cfg)
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
	Cmd.Flags().AddFlagSet(app.Cfg.Flags())
}
