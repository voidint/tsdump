package model

import "errors"

var (
	// ErrDBNotFound 数据库不存在
	ErrDBNotFound = errors.New("db not found")
)

// DB 数据库实例元信息
type DB struct {
	Name      string            `json:"name,omitempty"`
	CharSet   string            `json:"charset,omitempty"`
	Collation string            `json:"collation,omitempty"`
	Tables    []Table           `json:"tables,omitempty"`
	Extra     map[string]string `json:"extra,omitempty"`
}

// Table 表元信息
type Table struct {
	DB        string            `json:"-"`
	Name      string            `json:"name,omitempty"`
	Collation string            `json:"collation,omitempty"`
	Comment   string            `json:"comment,omitempty"`
	Columns   []Column          `json:"columns,omitempty"`
	Extra     map[string]string `json:"extra,omitempty"`
}

// Column 列元信息
type Column struct {
	DB        string            `json:"-"`
	Table     string            `json:"-"`
	Name      string            `json:"name,omitempty"`
	Default   string            `json:"default,omitempty"`
	Nullable  string            `json:"nullable,omitempty"`
	DataType  string            `json:"data_type,omitempty"`
	Key       string            `json:"key,omitempty"`
	CharSet   string            `json:"charset,omitempty"`
	Collation string            `json:"collation,omitempty"`
	Comment   string            `json:"comment,omitempty"`
	Extra     map[string]string `json:"extra,omitempty"`
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
