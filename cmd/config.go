package cmd

import "github.com/spf13/pflag"

var cfg Config

type Config struct {
	ApiToken string
	Bind     string
}

func (c *Config) Flags() *pflag.FlagSet {
	f := pflag.NewFlagSet("", pflag.PanicOnError)
	f.StringVar(&c.Bind, "bind", "0.0.0.0:6000", "host:port to bind")
	f.StringVar(&c.ApiToken, "api_token", "xxxx", "api token")
	return f
}
