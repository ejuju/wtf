package noise

import (
	"testing"
)

func TestOjracGenerator(t *testing.T) {
	t.Parallel()

	generator := NewOjracGenerator(0)

	t.Run("should generate different consecutive values", func(t *testing.T) {
		if generator.Get3D(0, 0, 0) == generator.Get3D(0.01, 0, 0) {
			t.Fatal("output noise values are the same for different inputs")
		}
	})

	t.Run("should generate expected value", func(t *testing.T) {
		got := generator.Get3D(0.5, 0.5, 0)
		want := 0.50171426145671
		if got != want {
			t.Fatalf("unexpected value, want %#v but got %#v", want, got)
		}
	})
}
