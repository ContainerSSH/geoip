# Changelog

## 0.9.5: Dependency bump

This release bumps the testify ang geoip-golang libraries.

## 0.9.4: Added Validate

This release adds a `Validate()` method to the configuration structure.

## 0.9.3: Shorter config name

This change renames the `Configuration` struct to `Config` for brevity.

## 0.9.1: Better usage

This release moves the New() method to the geoipprovider package and the LookupProvider interface to the geoip package for easier usage.

## 0.9.0: Initial release

### Using this library

This library needs a configuration structure described in [config.go](config.go). This configuration structure can be passed to the `geoip.New()` method:

```go
provider, err := geoip.New(geoip.Configuration{
    Provider: "maxmind",
    GeoIP2File: "/path/to/maxmind/file.mmdb2",
})
if err != nil {
    // handle error
}
```

The GeoIP lookup can be performed using the `Lookup()` method:

```go
countryCode := provider.Lookup("127.0.0.1")
```

The `countryCode` field will contain the value of `XX` if the lookup failed.

### Implementing a lookup provider

A custom provider can be written by implementing the following interface:

```go
type LookupProvider interface {
	Lookup(remoteAddr net.IP) (countryCode string)
}
```

Once implemented you will need to add the necessary configuration options to [config.go](config.go) and add a factory
method to [factory.go](factory.go).