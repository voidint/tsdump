package xlsx

import (
	"fmt"
	"io"
	"os"
	"unicode/utf8"

	"github.com/tealeg/xlsx/v3"
	"github.com/voidint/tsdump/model"
	"github.com/voidint/tsdump/view"
)

func init() {
	view.Register(Name, NewView())
}

const (
	// Name 视图名称
	Name = "xlsx"
)

// View Markdown视图
type View struct {
}

// NewView 返回Markdown视图实例
func NewView() view.Viewer {
	return new(View)
}

var (
	headerStyle *xlsx.Style
	rowStyle    *xlsx.Style
)

func init() {
	headerStyle = xlsx.NewStyle()
	headerStyle.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
	headerStyle.ApplyBorder = true
	headerStyle.Font = *xlsx.NewFont(12, "Verdana")
	headerStyle.ApplyFont = true

	rowStyle = xlsx.NewStyle()
	rowStyle.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
	rowStyle.ApplyBorder = true
	rowStyle.Font = *xlsx.NewFont(10, "Verdana")
	rowStyle.ApplyFont = true
}

var headerDefs = []struct {
	Title string
	Width float64
}{
	{Title: "Column", Width: 30},
	{Title: "Data Type", Width: 25},
	{Title: "Nullable", Width: 10},
	{Title: "Key", Width: 6},
	{Title: "Default", Width: 10},
	{Title: "Character Set", Width: 15},
	{Title: "Collation", Width: 20},
	{Title: "Comment", Width: 50},
}

// Do 将数据库元数据以Excel视图形式输出。
func (v *View) Do(items []model.DB, out io.Writer) error {
	if len(items) <= 0 {
		return nil
	}
	f := xlsx.NewFile()
	for _, table := range items[0].Tables {
		sheetName := table.Name

		runeLength := utf8.RuneCountInString(sheetName)
		if runeLength > 31 {
			fmt.Fprintf(os.Stderr, "The sheet name is too long, rename %q to %q\n", sheetName, sheetName[0:31])
			sheetName = sheetName[0:31]
		}

		sheet, err := f.AddSheet(sheetName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		firstCell := sheet.AddRow().AddCell()
		firstCell.SetString(fmt.Sprintf("Table: %s(%s)", table.Name, table.Comment))
		firstCell.Merge(7, 0)

		header := sheet.AddRow()
		for i, def := range headerDefs {
			col := xlsx.NewColForRange(i, i)
			col.SetWidth(def.Width)
			sheet.SetColParameters(col)

			cell := header.AddCell()
			cell.SetString(def.Title)
			cell.SetStyle(headerStyle)
		}

		for _, c := range table.Columns {
			row := sheet.AddRow()
			c0 := row.AddCell()
			c0.SetString(c.Name)
			c0.SetStyle(rowStyle)

			c1 := row.AddCell()
			c1.SetString(c.DataType)
			c1.SetStyle(rowStyle)

			c2 := row.AddCell()
			c2.SetString(c.Nullable)
			c2.SetStyle(rowStyle)

			c3 := row.AddCell()
			c3.SetString(c.Key)
			c3.SetStyle(rowStyle)

			c4 := row.AddCell()
			c4.SetString(c.Default)
			c4.SetStyle(rowStyle)

			c5 := row.AddCell()
			c5.SetString(c.CharSet)
			c5.SetStyle(rowStyle)

			c6 := row.AddCell()
			c6.SetString(c.Collation)
			c6.SetStyle(rowStyle)

			c7 := row.AddCell()
			c7.SetString(c.Comment)
			c7.SetStyle(rowStyle)
		}
	}
	defer func() {
		for i := range f.Sheets {
			f.Sheets[i].Close()
		}
	}()
	return f.Write(out)
}
