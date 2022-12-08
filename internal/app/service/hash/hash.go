package hash

type Hash interface {
	Generate(string) string
}
