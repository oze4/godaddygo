package godaddygo

// Gateway allows you to select versions
type Gateway interface {
	V1() V1
}

// gateway implements Gateway
type gateway struct {
	connectionBridge
}

func (g *gateway) V1() V1 {
	g.SetAPIVersion("v1")
	return &v1{g.connectionBridge}
}
