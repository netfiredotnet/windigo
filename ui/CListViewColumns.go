package ui

import (
	"fmt"
	"unsafe"

	"github.com/netfiredotnet/windigo/win"
	"github.com/netfiredotnet/windigo/win/co"
)

type _ListViewColumns struct {
	lv ListView
}

func (me *_ListViewColumns) new(ctrl ListView) {
	me.lv = ctrl
}

// Adds one or more columns with their widths.
// Widths will be adjusted to the current system DPI.
func (me *_ListViewColumns) Add(widths []int, titles ...string) {
	if len(titles) != len(widths) {
		panic(fmt.Sprintf("Unmatching titles (%d) and widths (%d).",
			len(titles), len(widths)))
	}

	lvc := win.LVCOLUMN{
		Mask: co.LVCF_TEXT | co.LVCF_WIDTH,
	}

	for i := 0; i < len(titles); i++ {
		colWidth := win.SIZE{Cx: int32(widths[i]), Cy: 0}
		_MultiplyDpi(nil, &colWidth)

		lvc.Cx = colWidth.Cx
		lvc.SetPszText(win.Str.ToNativeSlice(titles[i]))

		newIdx := int(
			me.lv.Hwnd().SendMessage(co.LVM_INSERTCOLUMN,
				0xffff, win.LPARAM(unsafe.Pointer(&lvc))),
		)
		if newIdx == -1 {
			panic(fmt.Sprintf("LVM_INSERTCOLUMN \"%s\" failed.", titles[i]))
		}
	}
}

// Retrieves the number of columns.
func (me *_ListViewColumns) Count() int {
	hHeader := win.HWND(me.lv.Hwnd().SendMessage(co.LVM_GETHEADER, 0, 0))
	if hHeader == 0 {
		panic("LVM_GETHEADER failed.")
	}

	count := int(hHeader.SendMessage(co.HDM_GETITEMCOUNT, 0, 0))
	if count == -1 {
		panic("HDM_GETITEMCOUNT failed.")
	}
	return count
}

// Retrieves all selected texts under this column.
func (me *_ListViewColumns) SelectedTexts(columnIndex int) []string {
	selItems := me.lv.Items().Selected()
	selTexts := make([]string, 0, len(selItems))

	for _, selItem := range selItems {
		selTexts = append(selTexts, selItem.Text(columnIndex))
	}
	return selTexts
}

// Sets the title of this column.
func (me *_ListViewColumns) SetTitle(columnIndex int, text string) {
	lvc := win.LVCOLUMN{
		ISubItem: int32(columnIndex),
		Mask:     co.LVCF_TEXT,
	}
	lvc.SetPszText(win.Str.ToNativeSlice(text))

	ret := me.lv.Hwnd().SendMessage(co.LVM_SETCOLUMN,
		win.WPARAM(columnIndex), win.LPARAM(unsafe.Pointer(&lvc)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETCOLUMN %d to \"%s\" failed.", columnIndex, text))
	}
}

// Sets the width of the column. Will be adjusted to the current system DPI.
func (me *_ListViewColumns) SetWidth(columnIndex, width int) {
	colWidth := win.SIZE{Cx: int32(width), Cy: 0}
	_MultiplyDpi(nil, &colWidth)

	ret := me.lv.Hwnd().SendMessage(co.LVM_SETCOLUMNWIDTH,
		win.WPARAM(columnIndex), win.LPARAM(colWidth.Cx))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETCOLUMNWIDTH %d to %d failed.", columnIndex, width))
	}
}

// Resizes the column to fill the remaining space.
func (me *_ListViewColumns) SetWidthToFill(columnIndex int) {
	numCols := me.Count()
	cxUsed := 0

	for i := 0; i < numCols; i++ {
		if i != columnIndex {
			cxUsed += me.Width(i) // retrieve cx of each column, but us
		}
	}

	rc := me.lv.Hwnd().GetClientRect() // list view client area
	me.lv.Hwnd().SendMessage(co.LVM_SETCOLUMNWIDTH,
		win.WPARAM(columnIndex), win.LPARAM(int(rc.Right)-cxUsed)) // fill available space
}

// Retrieves all texts under this column.
func (me *_ListViewColumns) Texts(columnIndex int) []string {
	items := me.lv.Items().All()
	texts := make([]string, 0, len(items))

	for _, item := range items {
		texts = append(texts, item.Text(columnIndex))
	}
	return texts
}

// Retrieves the title of the column at the given index.
func (me *_ListViewColumns) Title(columnIndex int) string {
	var titleBuf [128]uint16 // arbitrary

	lvc := win.LVCOLUMN{
		ISubItem: int32(columnIndex),
		Mask:     co.LVCF_TEXT,
	}
	lvc.SetPszText(titleBuf[:])

	ret := me.lv.Hwnd().SendMessage(co.LVM_GETCOLUMN,
		win.WPARAM(columnIndex), win.LPARAM(unsafe.Pointer(&lvc)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_GETCOLUMN %d failed.", lvc.ISubItem))
	}

	return win.Str.FromNativeSlice(titleBuf[:])
}

// Retrieves the width of the column.
func (me *_ListViewColumns) Width(columnIndex int) int {
	cx := int(
		me.lv.Hwnd().SendMessage(co.LVM_GETCOLUMNWIDTH,
			win.WPARAM(columnIndex), 0),
	)
	if cx == 0 {
		panic(fmt.Sprintf("LVM_GETCOLUMNWIDTH %d failed.", columnIndex))
	}
	return cx
}
