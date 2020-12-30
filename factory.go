package geoip

import (
	"fmt"

	"github.com/containerssh/geoip/dummy"
	"github.com/containerssh/geoip/geoipprovider"
	"github.com/containerssh/geoip/oschwald"
)

// New creates a new lookup provider based on the configuration.
func New(config Config) (geoipprovider.LookupProvider, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	switch config.Provider {
	case DummyProvider:
		return dummy.New()
	case MaxMindProvider:
		return oschwald.New(config.GeoIP2File)
	default:
		return nil, fmt.Errorf("invalid provider: %s", config.Provider)
	}
}
