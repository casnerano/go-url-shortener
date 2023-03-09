package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique_Generate(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Unique{}
			assert.Equalf(t, tt.want, u.Generate(tt.args.url), "Generate(%v)", tt.args.url)
		})
	}
}

func BenchmarkUnique_Generate(b *testing.B) {
	random := NewUnique()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		random.Generate("")
	}
}
