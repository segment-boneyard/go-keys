package keys

import "github.com/bmizerany/assert"
import "testing"
import "sort"

func TestKeys(t *testing.T) {
	m := map[string]interface{}{
		"foo": true,
		"bar": true,
		"baz": true,
	}

	got := Keys(m)
  sort.Strings(got)

	exp := []string{"bar", "baz", "foo"}
	assert.Equal(t, exp, got)
}
