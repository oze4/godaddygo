package uri

// Gateway allows you to target specific API versions
type Gateway interface {
	// We do not validate the `Version(v string)` parameter!
	// We expect you to have already validated the
	// version string beforehand
	Version(v string) Version
}

type gateway struct {
	*cache
}

// We do not validate the `v string` parameter!
// We expect you to have already validated the
// version string beforehand
func (g *gateway) Version(v string) Version {
	g.path += "/" + v
	return &version{g.cache}
}