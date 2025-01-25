package support

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIdGenerator_Generate(t *testing.T) {
	id, err := NewIdGenerator().Generate()
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), id%10)

	ig := &IdGenerator{&DefaultIncrementer{Id: 9999}, 4}
	id, _ = ig.Generate()
	assert.Equal(t, uint64(0), id%10000)

	id, _ = ig.Generate()
	assert.Equal(t, uint64(1), id%10000)
}
