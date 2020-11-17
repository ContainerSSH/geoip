package dummy

import (
	"github.com/containerssh/geoip"
)

// New creates a dummy provider that always responds with "XX"
func New() (geoip.LookupProvider, error) {
	return &geoIPLookupProvider{}, nil
}
