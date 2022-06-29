package noise

import (
	"testing"
)

func TestAquilaxGenerator(t *testing.T) {
	t.Parallel()

	generator := NewAquilaxGenerator(AquilaxGeneratorConfig{Seed: 0, N: 2, Alpha: 2, Beta: 2})

	t.Run("should generate different consecutive values", func(t *testing.T) {
		if generator.Get3D(0, 0, 0) == generator.Get3D(0.01, 0, 0) {
			t.Fatal("output noise values are the same for different inputs")
		}
	})

	t.Run("should generate expected value", func(t *testing.T) {
		got := generator.Get3D(0.5, 0.5, 0)
		want := 0.06357502654967262
		if got != want {
			t.Fatalf("unexpected value, want %#v but got %#v", want, got)
		}
	})
}
