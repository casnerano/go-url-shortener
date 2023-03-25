package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique_Generate(t *testing.T) {
	t.Run("generate unique string length", func(t *testing.T) {
		u := NewUnique()
		uniqStr := u.Generate("example")
		assert.Equalf(t, 64, len(uniqStr), "Generate(%v)", uniqStr)
	})
}

func BenchmarkUnique_Generate(b *testing.B) {
	random := NewUnique()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		random.Generate("")
	}
}
