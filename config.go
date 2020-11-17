package geoip

// Provider is a configuration option to select the GeoIP lookup provider.
type Provider string

const (
	// DummyProvider always returns the XX country code.
	DummyProvider Provider = "dummy"
	// MaxMindProvider looks up IP addresses via the MaxMind GeoIP database.
	MaxMindProvider Provider = "maxmind"
)

// Configuration is the structure configuring the GeoIP lookup process.
type Configuration struct {
	Provider   Provider `yaml:"provider" json:"provider" default:"dummy"`
	GeoIP2File string   `yaml:"maxmind-geoip2-file" json:"geoip2-file" default:"/var/lib/GeoIP/GeoIP2-Country.mmdb"`
}
