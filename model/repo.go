package model

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
	Table     string
	Nullable  string
	DataType  string
	CharSet   string
	Collation string
	Comment   string
	Extra     map[string]string
}

type Repo interface {
	GetDBs(cond *DB) ([]DB, error)
	GetTables(table *Table) ([]Table, error)
	GetColumns(col *Column) ([]Column, error)
}
