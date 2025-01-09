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
