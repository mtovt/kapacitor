package reporting

import (
	"github.com/influxdb/usage-client/v1"
)

type Config struct {
	Enabled bool   `toml:"enabled"`
	URL     string `toml:"url"`
}

func NewConfig() Config {
	return Config{
		Enabled: true,
		URL:     client.URL,
	}
}
