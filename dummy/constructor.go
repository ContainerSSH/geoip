package dummy

import (
	"github.com/containerssh/geoip/geoipprovider"
)

func New() (geoipprovider.LookupProvider, error) {
	return &geoIPLookupProvider{}, nil
}
