package errox

import "github.com/pkg/errors"

// IsAny returns a bool if it matches any of the target errors
// This helps consolidate code from
// errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.ErrClosedPipe)
// to errors.IsAny(err, io.EOF. io.ErrUnexpectedEOF, io.ErrClosedPipe)
func IsAny(err error, targets ...error) bool {
	for _, target := range targets {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

// GetBaseSentinelMessage returns the error message of the lowest found
// sentinel error.
func GetBaseSentinelMessage(err error) string {
	var re Error
	for errors.As(err, &re) {
		err = errors.Unwrap(re)
	}
	if re, ok := (re).(*RoxError); ok && re != nil {
		return re.message
	}
	return ""
}
