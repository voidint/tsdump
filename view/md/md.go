package md

import (
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/voidint/tsdump/model"
	"github.com/voidint/tsdump/view"
)

type MarkdownView struct {
}

func NewView() view.Viewer {
	return new(MarkdownView)
}

func (v *MarkdownView) Do(items []model.DB, out io.Writer) error {
	for i := range items {
		v.renderDB(&items[i], out)
		fmt.Fprintln(out)
		for j := range items[i].Tables {
			fmt.Fprintf(out, "### `%s`\n%s\n\n",
				items[i].Tables[j].Name,
				items[i].Tables[j].Comment,
			)
			v.renderTable(&items[i].Tables[j], out)
			fmt.Fprintln(out)
		}
	}

	return nil
}

func (v *MarkdownView) renderDB(db *model.DB, out io.Writer) {
	rows := [][]string{[]string{db.Name, db.CharSet, db.Collation}}

	t := tablewriter.NewWriter(out)
	t.SetHeader([]string{"Database", "Character Set", "Collation"})
	t.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	t.SetCenterSeparator("|")
	t.AppendBulk(rows)
	t.Render()
}

func (v *MarkdownView) renderTable(table *model.Table, out io.Writer) {
	rows := make([][]string, 0, len(table.Columns))
	for i := range table.Columns {
		rows = append(rows, []string{
			table.Columns[i].Name,
			table.Columns[i].Nullable,
			table.Columns[i].DataType,
			table.Columns[i].CharSet,
			table.Columns[i].Collation,
			table.Columns[i].Comment,
		})
	}

	t := tablewriter.NewWriter(out)
	t.SetHeader([]string{"Column", "Nullable", "Data Type", "Character Set", "Collation", "Comment"})
	t.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	t.SetCenterSeparator("|")
	t.AppendBulk(rows)
	t.Render()
}
