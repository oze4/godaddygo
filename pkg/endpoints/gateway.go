package endpoints

// Gateway allows you to select versions
type Gateway interface {
	V1() V1
}

// gateway implements Gateway
type gateway struct {
	connectionInterface
}

func (g *gateway) V1() V1 {
	return &v1{g.connectionInterface}
}
