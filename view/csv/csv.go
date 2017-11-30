package csv

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/voidint/tsdump/model"
	"github.com/voidint/tsdump/view"
)

type CSVView struct {
}

func NewView() view.Viewer {
	return new(CSVView)
}

func (v *CSVView) Do(items []model.DB, out io.Writer) (err error) {
	for i := range items {
		if err = v.renderDB(&items[i], out); err != nil {
			return err
		}
		fmt.Fprintf(out, "\n\n\n")
		for j := range items[i].Tables {
			fmt.Fprintf(out, "TABLE: %s,%s\n",
				items[i].Tables[j].Name,
				items[i].Tables[j].Comment,
			)
			if err = v.renderTable(&items[i].Tables[j], out); err != nil {
				return err
			}
			fmt.Fprintf(out, "\n\n\n")
		}
	}
	return nil
}

func (v *CSVView) renderDB(db *model.DB, out io.Writer) error {
	rows := [][]string{
		[]string{"Database", "Character Set", "Collation"},
		[]string{db.Name, db.CharSet, db.Collation},
	}
	return csv.NewWriter(out).WriteAll(rows)
}

func (v *CSVView) renderTable(table *model.Table, out io.Writer) error {
	rows := make([][]string, 0, len(table.Columns)+1)
	rows = append(rows, []string{"Column", "Nullable", "Data Type", "Character Set", "Collation", "Comment"})
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
	return csv.NewWriter(out).WriteAll(rows)
}
