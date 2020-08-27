package url

// Records is a struct builds out the `records` piece of GoDaddy's API
type Records struct {
    domain Domain
}

// func (r Records) Type() Type {}