package txt

import (
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/voidint/tsdump/model"
	"github.com/voidint/tsdump/view"
)

type TXTView struct {
}

func NewView() view.Viewer {
	return new(TXTView)
}

func (v *TXTView) Do(items []model.DB, out io.Writer) error {
	for i := range items {
		v.renderDB(&items[i], out)
		fmt.Fprintln(out)
		for j := range items[i].Tables {
			fmt.Fprintf(out, "TABLE:\t%s\t%s\n",
				items[i].Tables[j].Name,
				items[i].Tables[j].Comment,
			)
			v.renderTable(&items[i].Tables[j], out)
			fmt.Fprintln(out)
		}
	}
	return nil
}

func (v *TXTView) renderDB(db *model.DB, out io.Writer) error {
	rows := [][]string{[]string{db.Name, db.CharSet, db.Collation}}

	t := tablewriter.NewWriter(out)
	t.SetHeader([]string{"Database", "Character Set", "Collation"})
	t.SetCenterSeparator("|")
	t.AppendBulk(rows)
	t.Render()
	return nil
}

func (v *TXTView) renderTable(table *model.Table, out io.Writer) {
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
	t.SetCenterSeparator("|")
	t.AppendBulk(rows)
	t.Render()
}
