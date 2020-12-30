package geoip

import (
	"fmt"
	"os"
)

// Provider is a configuration option to select the GeoIP lookup provider.
type Provider string

const (
	// DummyProvider always returns the XX country code.
	DummyProvider Provider = "dummy"
	// MaxMindProvider looks up IP addresses via the MaxMind GeoIP database.
	MaxMindProvider Provider = "maxmind"
)

func (p Provider) Validate() error {
	switch p {
	case DummyProvider:
	case MaxMindProvider:
	default:
		return fmt.Errorf("invalid GeoIP provider: %s", p)
	}
	return nil
}

// Config is the structure configuring the GeoIP lookup process.
type Config struct {
	Provider   Provider `yaml:"provider" json:"provider" default:"dummy"`
	GeoIP2File string   `yaml:"maxmind-geoip2-file" json:"maxmind-geoip2-file" default:"/var/lib/GeoIP/GeoIP2-Country.mmdb"`
}

// Validate checks the configuration.
func (config Config) Validate() error {
	if err := config.Provider.Validate(); err != nil {
		return err
	}
	if config.Provider == MaxMindProvider {
		stat, err := os.Stat(config.GeoIP2File)
		if err != nil {
			return fmt.Errorf("invalid MaxMind GeoIP2 file: %s (%w)", config.GeoIP2File, err)
		}
		if stat.IsDir() {
			return fmt.Errorf("invalid MaxMind GeoIP2 file: %s (is a directory)", config.GeoIP2File)
		}
	}
	return nil
}
