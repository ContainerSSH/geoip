package geoip

import (
	"fmt"

	"github.com/containerssh/geoip/dummy"
	"github.com/containerssh/geoip/geoipprovider"
	"github.com/containerssh/geoip/oschwald"
)

// New creates a new lookup provider based on the configuration.
//goland:noinspection GoUnusedExportedFunction
func New(config Configuration) (provider geoipprovider.LookupProvider, err error) {
	switch config.Provider {
	case DummyProvider:
		provider, err = dummy.New()
	case MaxMindProvider:
		provider, err = oschwald.New(config.GeoIP2File)
	default:
		return nil, fmt.Errorf("invalid provider: %s", config.Provider)
	}
	return provider, err
}
