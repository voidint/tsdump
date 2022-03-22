package xlsx

import (
	"io"

	"github.com/voidint/tsdump/model"
	"github.com/voidint/tsdump/view"
	"gopkg.in/yaml.v2"
)

func init() {
	view.Register(Name, NewView())
}

const (
	// Name 视图名称
	Name = "yaml"
)

// View Markdown视图
type View struct {
}

// NewView 返回Markdown视图实例
func NewView() view.Viewer {
	return new(View)
}

type Column struct {
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
	Comment string `yaml:"comment"`
	Default string `yaml:"default"`
}

type Table struct {
	Name    string   `yaml:"name"`
	Columns []Column `yaml:"columns"`
}

type Database struct {
	Name   string   `yaml:"name"`
	Tables []*Table `yaml:"tables"`
}

// Do 将数据库元数据以JSON视图形式输出。
func (v *View) Do(items []model.DB, out io.Writer) error {
	ds := map[string]*Database{}
	for _, db := range items {
		d := ds[db.Name]
		if d == nil {
			d = &Database{
				Name: db.Name,
			}
		}

		found := false
		for _, table := range db.Tables {
			// find table
			var t *Table
			for _, tt := range d.Tables {
				if tt.Name == table.Name {
					t = tt
					found = true
					break
				}
			}
			if t == nil {
				t = &Table{
					Name: table.Name,
				}
			}

			var tableColumns []Column
			for _, c := range table.Columns {
				cc := Column{
					Name:    c.Name,
					Type:    c.DataType,
					Comment: c.Comment,
					Default: c.Default,
				}
				tableColumns = append(tableColumns, cc)
			}

			t.Columns = append(t.Columns, tableColumns...)

			if !found {
				d.Tables = append(d.Tables, t)
			}
		}

		ds[db.Name] = d
	}

	enc := yaml.NewEncoder(out)
	return enc.Encode(ds)
}
