package oschwald

import (
	backend "github.com/oschwald/geoip2-golang"
)

func New(geoIpFile string) (*GeoIPLookupProvider, error) {
	geo, err := backend.Open(geoIpFile)
	if err != nil {
		return nil, err
	}
	return &GeoIPLookupProvider{
		geo: geo,
	}, nil
}
