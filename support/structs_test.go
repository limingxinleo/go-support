package support

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type DataExample struct {
	Id        uint64 `db:"id"`
	RequestId string `db:"request_id"`
	AdId      string `db:"ad_id"`
	AdName    string `db:"ad_name"`
}

func TestStructSupport_Fill(t *testing.T) {
	m := &DataExample{}
	err := NewStructSupport().Fill(m, map[string]string{"id": "123", "request_id": "1234"}, "db")

	assert.Nil(t, err)
	assert.Equal(t, m.Id, uint64(123))
	assert.Equal(t, m.RequestId, "1234")
}
