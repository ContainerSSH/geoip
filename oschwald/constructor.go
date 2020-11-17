package oschwald

import (
	backend "github.com/oschwald/geoip2-golang"

	"github.com/containerssh/geoip"
)

// New creates a new GeoIP lookup provider using Oschwald's backend.
func New(geoIpFile string) (geoip.LookupProvider, error) {
	geo, err := backend.Open(geoIpFile)
	if err != nil {
		return nil, err
	}
	return &geoIPLookupProvider{
		geo: geo,
	}, nil
}
