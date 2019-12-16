package bloom_filter

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBloom_Filter(t *testing.T) {
	f := filter{}
	f.set("HelloWord")

	assert.Equal(t, f.get("HelloWord"), true)
}