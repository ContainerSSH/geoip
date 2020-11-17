package geoipfactory

import (
	"fmt"

	"github.com/containerssh/geoip"
	"github.com/containerssh/geoip/dummy"
	"github.com/containerssh/geoip/oschwald"
)

// New creates a new lookup provider based on the configuration.
//goland:noinspection GoUnusedExportedFunction
func New(config geoip.Configuration) (geoip.LookupProvider, error) {
	switch config.Provider {
	case geoip.DummyProvider:
		return dummy.New()
	case geoip.MaxMindProvider:
		return oschwald.New(config.GeoIP2File)
	default:
		return nil, fmt.Errorf("invalid provider: %s", config.Provider)
	}
}
