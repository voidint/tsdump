package model

import "errors"

var (
	// ErrDBNotFound 数据库不存在
	ErrDBNotFound = errors.New("db not found")
)

// DB 数据库实例
type DB struct {
	Name      string
	CharSet   string
	Collation string
	Tables    []Table
	Extra     map[string]string
}

// Table 表
type Table struct {
	DB        string
	Name      string
	Collation string
	Comment   string
	Columns   []Column
	Extra     map[string]string
}

// Column 列
type Column struct {
	DB        string
	Table     string
	Name      string
	Nullable  string
	DataType  string
	CharSet   string
	Collation string
	Comment   string
	Extra     map[string]string
}

type IRepo interface {
	GetDBs(cond *DB) ([]DB, error)
	GetTables(cond *Table) ([]Table, error)
	GetColumns(cond *Column) ([]Column, error)
}
