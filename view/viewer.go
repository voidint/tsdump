package view

import (
	"io"

	"github.com/voidint/tsdump/model"
)

// Viewer 数据视图
type Viewer interface {
	Do(items []model.DB, out io.Writer) error
}
