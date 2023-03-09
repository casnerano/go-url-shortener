// Package for shortening link services.
package hasher

// Defines an interface for generating a short link from the original.
type Hash interface {
	Generate(string) string
}
