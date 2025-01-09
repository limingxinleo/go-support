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

func (s *SqlSupport) NewNullTimeByString(value string, layout string) sql.NullTime {
	t, err := time.ParseInLocation(layout, value, time.Local)
	if err != nil {
		return sql.NullTime{}
	}

	return sql.NullTime{Time: t, Valid: true}
}

func (s *SqlSupport) NewNullTimeByDateString(value string) sql.NullTime {
	if len(value) < 10 {
		return sql.NullTime{}
	}

	return s.NewNullTimeByString(value[0:10], time.DateOnly)
}
