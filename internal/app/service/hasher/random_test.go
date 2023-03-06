package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRandom(t *testing.T) {
	tests := []struct {
		name     string
		min, max int
		wantErr  bool
	}{
		{"less than minimum", -1, 10, true},
		{"normal length", 10, 15, false},
		{"greater than the maximum", 100, 150, true},
		{"minimum is greater than maximum", 50, 10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewRandom(tt.min, tt.max)

			if !tt.wantErr {
				require.NoError(t, err)
			}
		})
	}
}

func TestRandom_Generate(t *testing.T) {
	tests := []struct {
		name     string
		min, max int
	}{
		{"length 1..2", 1, 2},
		{"length 3..5", 3, 5},
		{"length 5..5", 5, 5},
		{"length 5..10", 5, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			random, err := NewRandom(tt.min, tt.max)
			require.NoError(t, err)

			l := len(random.Generate(""))
			assert.Equal(t, true, l >= tt.min && l <= tt.max)
		})
	}
}

func BenchmarkRandom_Generate(b *testing.B) {
	random, _ := NewRandom(64, 64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		random.Generate("")
	}
}

func TestRandom_getRandomString(t *testing.T) {
	tests := []struct {
		name string
		len  int
	}{
		{"zero length", 0},
		{"less than minimum", 5},
		{"normal length", 15},
		{"greater than the maximum", 100},
	}

	random := Random{10, 20}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.len, len(random.getRandomString(tt.len)))
		})
	}
}
