package session

import (
	"database/sql"
	"orm/log"
	"strings"
)

type Sessions struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []any
}

func New(db *sql.DB) *Sessions {
	return &Sessions{db: db}
}

func (s *Sessions) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

func (s *Sessions) DB() *sql.DB {
	return s.db
}

func (s *Sessions) Raw(sql string, values ...any) *Sessions {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

func (s *Sessions) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

func (s *Sessions) QueryRow() (row *sql.Row) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

func (s *Sessions) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
