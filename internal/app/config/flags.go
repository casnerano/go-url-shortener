package config

import "flag"

// Application flags store
type Flags struct {
	ConfigName string
	Server     struct {
		TrustedSubnet string
		Addr          string
		EnableHTTPS   bool
		BaseURL       string
	}
	Storage struct {
		Path string
		DSN  string
	}
}

// ParseFlags parsing application flags.
func parseFlags(c *Config) *Flags {
	flags := &Flags{}
	flag.StringVar(&flags.Server.Addr, "a", c.Server.Addr, "Server addr")
	flag.BoolVar(&flags.Server.EnableHTTPS, "s", c.Server.EnableHTTPS, "Enable HTTPS")
	flag.StringVar(&flags.Server.BaseURL, "b", c.Server.BaseURL, "Base URL")
	flag.StringVar(&flags.Server.TrustedSubnet, "t", c.Server.TrustedSubnet, "Trusted Subnet")
	flag.StringVar(&flags.Storage.Path, "f", c.Storage.Path, "File storage path")
	flag.StringVar(&flags.Storage.DSN, "d", c.Storage.DSN, "Database connection DSN")

	flag.StringVar(&flags.ConfigName, "c", "", "Configuration file name (alias for -config)")
	flag.StringVar(&flags.ConfigName, "config", flags.ConfigName, "Configuration file name")

	flag.Parse()

	return flags
}
