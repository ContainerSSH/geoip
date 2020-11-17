package oschwald

import (
	"net"
)

func (g *geoIPLookupProvider) Lookup(remoteAddr net.IP) (countryCode string) {
	country, err := g.geo.Country(remoteAddr)
	if err != nil {
		return "XX"
	}
	return country.Country.IsoCode
}
