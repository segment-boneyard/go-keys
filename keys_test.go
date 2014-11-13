package keys

import "github.com/bmizerany/assert"
import "testing"

func TestKeys(t *testing.T) {
	m := map[string]interface{}{
		"foo": true,
		"bar": true,
		"baz": true,
	}

	got := Sorted(m)
	exp := []string{"bar", "baz", "foo"}
	assert.Equal(t, exp, got)
}
