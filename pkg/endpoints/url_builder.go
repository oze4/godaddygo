package endpoints

/**
 *
 * I felt the need to put this in it's own file to make it stand out more
 *
 * Otherwise, it was difficult to tell where this was coming from
 *
 */

import (
	"github.com/oze4/godaddygo/pkg/uri"
)

// URLBuilder wraps `uri.Builder`
func (s *session) URLBuilder() uri.Version {
	// s.IsProduction *has* to already be set by the time we see
	// it here
	// It does not matter if `s.apiVersion` is empty here or not
	// We let the API return an error, which should guide us
	p := s.IsProduction()
	v := s.apiVersion
	return uri.Builder(p).Version(v)
}
