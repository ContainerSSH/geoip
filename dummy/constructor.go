package dummy

import (
	"github.com/containerssh/geoip/geoipprovider"
)

// New creates a dummy provider that always responds with "XX"
func New() (geoipprovider.LookupProvider, error) {
	return &geoIPLookupProvider{}, nil
}
