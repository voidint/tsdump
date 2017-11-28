package txt

import (
	"fmt"
	"io"

	"github.com/bndr/gotabulate"
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
		dbT := v.dbTabulate(&items[i])
		dbT.SetWrapStrings(true)
		dbT.SetMaxCellSize(16)
		dbT.SetAlign("left")
		fmt.Fprintln(out, dbT.Render("grid"))

		for j := range items[i].Tables {
			tabT := v.tableTabulate(&items[i].Tables[j])
			tabT.SetWrapStrings(true)
			tabT.SetMaxCellSize(16)
			tabT.SetAlign("left")
			fmt.Fprintln(out, tabT.Render("grid"))
		}
	}

	return nil
}

func (v *TXTView) dbTabulate(db *model.DB) *gotabulate.Tabulate {
	headers := []string{"Database", "Character Set", "Collation"}
	row := []string{db.Name, db.CharSet, db.Collation}

	return gotabulate.Create([][]string{row}).SetHeaders(headers)
}

func (v *TXTView) tableTabulate(table *model.Table) *gotabulate.Tabulate {
	headers := []string{"Column", "Nullable", "Data Type", "Character Set", "Collation", "Comment"}
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
	return gotabulate.Create(rows).SetHeaders(headers)

}
