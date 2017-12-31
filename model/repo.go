package model

import "errors"

var (
	// ErrDBNotFound 数据库不存在
	ErrDBNotFound = errors.New("db not found")
)

// DB 数据库实例元信息
type DB struct {
	Name      string
	CharSet   string
	Collation string
	Tables    []Table
	Extra     map[string]string
}

// Table 表元信息
type Table struct {
	DB        string
	Name      string
	Collation string
	Comment   string
	Columns   []Column
	Extra     map[string]string
}

// Column 列元信息
type Column struct {
	DB        string
	Table     string
	Name      string
	Default   string
	Nullable  string
	DataType  string
	Key       string
	CharSet   string
	Collation string
	Comment   string
	Extra     map[string]string
}

// IRepo 数据库元信息查询接口
type IRepo interface {
	// GetDBs 查询数据库元信息
	GetDBs(cond *DB, lazy bool) ([]DB, error)
	// GetTables 查询表元信息
	GetTables(cond *Table) ([]Table, error)
	// GetColumns 查询列元信息
	GetColumns(cond *Column) ([]Column, error)
}
