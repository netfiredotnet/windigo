/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"fmt"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// A single item row of a list view control.
type ListViewItem struct {
	owner *ListView
	index uint32
}

func (me *ListViewItem) Delete() {
	if me.index >= me.owner.ItemCount() { // index out of bounds: ignore
		return
	}
	ret := me.owner.sendLvmMessage(co.LVM_DELETEITEM,
		win.WPARAM(me.index), 0)
	if ret == 0 {
		panic(fmt.Sprintf("LVM_DELETEITEM failed, index %d.", me.index))
	}
}

func (me *ListViewItem) Focused() bool {
	return (me.State() & co.LVIS_FOCUSED) != 0
}

func (me *ListViewItem) Index() uint32 {
	return me.index
}

func (me *ListViewItem) Owner() *ListView {
	return me.owner
}

func (me *ListViewItem) Selected() bool {
	return (me.State() & co.LVIS_SELECTED) != 0
}

func (me *ListViewItem) SetFocus() *ListViewItem {
	return me.SetState(co.LVIS_FOCUSED, co.LVIS_FOCUSED)
}

func (me *ListViewItem) SetSelected(selected bool) *ListViewItem {
	state := co.LVIS(0)
	if selected {
		state = co.LVIS_SELECTED
	}
	return me.SetState(state, co.LVIS_SELECTED)
}

func (me *ListViewItem) SetState(
	state co.LVIS, stateMask co.LVIS) *ListViewItem {

	lvi := win.LVITEM{
		State:     state,
		StateMask: stateMask,
	}
	ret := me.owner.sendLvmMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed.")
	}
	return me
}

func (me *ListViewItem) SetText(text string) *ListViewItem {
	me.SubItem(0).SetText(text)
	return me
}

func (me *ListViewItem) State() co.LVIS {
	return co.LVIS(
		me.owner.sendLvmMessage(co.LVM_GETITEMSTATE,
			win.WPARAM(me.index), win.LPARAM(co.LVIS_FOCUSED)),
	)
}

func (me *ListViewItem) SubItem(index uint32) *ListViewSubItem {
	numCols := me.owner.ColumnCount()
	if index >= numCols {
		panic("Trying to retrieve sub item with index out of bounds.")
	}
	return &ListViewSubItem{
		item:  me,
		index: index,
	}
}

func (me *ListViewItem) Text() string {
	return me.SubItem(0).Text()
}

// Sends LVM_UPDATE for this item.
func (me *ListViewItem) Update() *ListViewItem {
	ret := me.owner.sendLvmMessage(co.LVM_UPDATE, win.WPARAM(me.index), 0)
	if ret == 0 {
		panic("LVM_UPDATE failed.")
	}
	return me
}

func (me *ListViewItem) Visible() bool {
	return me.owner.sendLvmMessage(co.LVM_ISITEMVISIBLE,
		win.WPARAM(me.index), 0) != 0
}
