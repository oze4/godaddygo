package api

// Gateway implements Gateway
type Gateway struct {
	V1 v1
	currentRequest
}

// V1 returns the V1 section of the GoDaddy API
// func (g *gateway) V1() V1 {
// 	return &v1{g.currentRequest}
// }

// V2 returns the V2 section of the GoDaddy API
// func (g *gateway) V2() {
// 	panic("V2 not implemented yet!")
// }
