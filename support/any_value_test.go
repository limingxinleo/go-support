package support

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnyValueFactory_NewByUint64(t *testing.T) {
	v := NewAnyValueFactory().NewByUint64(uint64(9999))

	assert.Equal(t, uint64(9999), v.Uint64())
	assert.Equal(t, "9999", v.String())
}
