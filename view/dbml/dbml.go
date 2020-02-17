package dbml

import (
	"fmt"
	"io"
	"strings"

	"github.com/voidint/tsdump/model"
	"github.com/voidint/tsdump/view"
)

func init() {
	view.Register(Name, NewView())
}

const (
	// Name 视图名称
	Name = "dbml"
)

// View DBML视图
type View struct {
}

// NewView 返回DBML视图实例
func NewView() view.Viewer {
	return new(View)
}

// Do 将数据库元数据以DBML视图形式输出。
func (v *View) Do(items []model.DB, out io.Writer) error {
	for i := range items {
		for j := range items[i].Tables {
			v.renderTable(&items[i].Tables[j], out)
			v.renderEnums(&items[i].Tables[j], out)
		}
	}

	return nil
}

func (v *View) renderTable(table *model.Table, out io.Writer) {
	fmt.Fprintf(out, "Table %s {\n", table.Name)
	for i := range table.Columns {
		v.renderColumn(&table.Columns[i], out)
	}
	fmt.Fprintf(out, "}\n\n")
}

func (v *View) renderColumn(col *model.Column, out io.Writer) {
	var ssettings string
	if settings := v.columnSettings(col); len(settings) > 0 {
		ssettings = fmt.Sprintf("[%s]", strings.Join(settings, ", "))
	}
	fmt.Fprintf(out, "\t%s %s %s\n", col.Name, v.columnType(col), ssettings)
}

func (v *View) columnType(col *model.Column) string {
	tpe := col.DataType
	if strings.HasPrefix(tpe, "enum") {
		return fmt.Sprintf("enum_%s_%s", col.Table, col.Name)
	}
	return strings.Fields(tpe)[0]
}

func (v *View) columnSettings(col *model.Column) (settings []string) {
	if col.Key == "PRI" {
		settings = append(settings, "pk")
	}
	if col.Key == "UNI" {
		settings = append(settings, "unique")
	}

	if len(col.Extra) > 0 {
		for i := range col.Extra {
			if col.Extra[i] == "auto_increment" {
				settings = append(settings, "increment")
				break
			}
		}
	}
	if col.Nullable == "YES" {
		settings = append(settings, "null")
	}
	if col.Nullable == "NO" {
		settings = append(settings, "not null")
	}
	if col.Comment != "" {
		settings = append(settings, fmt.Sprintf("note: %q", col.Comment))
	}
	return settings
}

// renderEnums 渲染表格中的枚举对象
func (v *View) renderEnums(table *model.Table, out io.Writer) {
	for i := range table.Columns {
		v.renderColumnEnum(&table.Columns[i], out)
	}
}

// renderColumnEnum 渲染列中的枚举对象
func (v *View) renderColumnEnum(col *model.Column, out io.Writer) {
	if col == nil || !strings.HasPrefix(col.DataType, "enum") {
		return
	}
	fmt.Fprintf(out, "enum enum_%s_%s {\n", col.Table, col.Name)
	tpe := strings.TrimSuffix(strings.TrimPrefix(col.DataType, "enum("), ")")
	for _, v := range strings.Split(tpe, ",") {
		v = strings.TrimSuffix(strings.TrimPrefix(v, "'"), "'")
		v = strings.ReplaceAll(v, "-", "_") // 枚举值中不能出现中划线
		fmt.Fprintf(out, "\t%s\n", v)
	}
	fmt.Fprintf(out, "}\n\n")
}
