package support

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNullTimeString(t *testing.T) {
	ts, _ := time.Parse(time.DateOnly, "2025-01-01")
	assert.Equal(t, "2025-01-01", NewSqlSupport().NullTimeDateString(sql.NullTime{Time: ts, Valid: true}))
	assert.Equal(t, "", NewSqlSupport().NullTimeDateString(sql.NullTime{Valid: false}))
}

func TestNewNullTimeByDateString(t *testing.T) {
	nt := NewSqlSupport().NewNullTimeByDateString("2025-01-01")
	assert.Equal(t, "2025-01-01", nt.Time.Format(time.DateOnly))
	assert.Equal(t, int64(1735660800), nt.Time.Unix())
	assert.True(t, nt.Valid)

	nt = NewSqlSupport().NewNullTimeByDateString("2025-01-01 08:00:00")
	assert.True(t, nt.Valid)

	nt = NewSqlSupport().NewNullTimeByDateString("asdf")
	assert.False(t, nt.Valid)
}
