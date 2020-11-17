package geoip

import (
	"fmt"

	"github.com/containerssh/geoip/dummy"
	"github.com/containerssh/geoip/geoipprovider"
	"github.com/containerssh/geoip/oschwald"
)

// New creates a new lookup provider based on the configuration.
//goland:noinspection GoUnusedExportedFunction
func New(config Configuration) (geoipprovider.LookupProvider, error) {
	switch config.Provider {
	case DummyProvider:
		return dummy.New()
	case MaxMindProvider:
		return oschwald.New(config.GeoIP2File)
	default:
		return nil, fmt.Errorf("invalid provider: %s", config.Provider)
	}
}
