package support

import (
	"database/sql"
	"time"
)

type SqlSupport struct {
}

func NewSqlSupport() *SqlSupport {
	return &SqlSupport{}
}

func (s *SqlSupport) NullTimeString(value sql.NullTime, layout string) string {
	if !value.Valid {
		return ""
	}

	return value.Time.Format(layout)
}

func (s *SqlSupport) NullTimeDateTimeString(value sql.NullTime) string {
	return s.NullTimeString(value, time.DateTime)
}

func (s *SqlSupport) NullTimeDateString(value sql.NullTime) string {
	return s.NullTimeString(value, time.DateOnly)
}
