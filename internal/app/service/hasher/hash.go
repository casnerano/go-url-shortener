package hasher

type Hash interface {
	Generate(string) string
}
