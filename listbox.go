package gform

import (
	"syscall"
	"unsafe"

	"github.com/jx3fans/w32"
)

type ListBox struct {
	W32Control

	//	onEndLabelEdit EventManager
	onDBLClick EventManager
	onClick    EventManager
}

func NewListBox(parent Controller) *ListBox {
	lv := new(ListBox)
	lv.init(parent)

	lv.SetFont(DefaultFont)
	lv.SetSize(100, 100)

	return lv
}

func AttachListBox(parent Controller, id int32) *ListBox {
	lv := new(ListBox)
	lv.attach(parent, id)
	RegMsgHandler(lv)
	//w32.SendMessage(lv.Handle(), w32.LVM_SETUNICODEFORMAT, w32.TRUE, 0)
	return lv
}

func (this *ListBox) init(parent Controller) {
	this.W32Control.init("ListBox", parent, 0, w32.WS_CHILD|w32.WS_VISIBLE|w32.WS_BORDER|w32.LVS_REPORT|w32.LVS_EDITLABELS)
	RegMsgHandler(this)
}

// Changes the state of an item in a list-view control. Refer SETSEL message.
func (this *ListBox) setItemState(i int, state, mask uint) {
	var item w32.LVITEM
	item.State, item.StateMask = uint32(state), uint32(mask)

	w32.SendMessage(this.hwnd, w32.LB_SETSEL, uintptr(i), uintptr(unsafe.Pointer(&item)))
}

func (this *ListBox) EnableSingleSelect(enable bool) {
	ToggleStyle(this.hwnd, enable, w32.LVS_SINGLESEL)
}

func (this *ListBox) EnableSortHeader(enable bool) {
	ToggleStyle(this.hwnd, enable, w32.LVS_NOSORTHEADER)
}

func (this *ListBox) EnableSortAscending(enable bool) {
	ToggleStyle(this.hwnd, enable, w32.LVS_SORTASCENDING)
}

func (this *ListBox) EnableEditLabels(enable bool) {
	ToggleStyle(this.hwnd, enable, w32.LVS_EDITLABELS)
}

func (this *ListBox) EnableFullRowSelect(enable bool) {
	if enable {
		w32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, 0, w32.LVS_EX_FULLROWSELECT)
	} else {
		w32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, w32.LVS_EX_FULLROWSELECT, 0)
	}
}

func (this *ListBox) EnableDoubleBuffer(enable bool) {
	if enable {
		w32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, 0, w32.LVS_EX_DOUBLEBUFFER)
	} else {
		w32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, w32.LVS_EX_DOUBLEBUFFER, 0)
	}
}

func (this *ListBox) EnableHotTrack(enable bool) {
	if enable {
		w32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, 0, w32.LVS_EX_TRACKSELECT)
	} else {
		w32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, w32.LVS_EX_TRACKSELECT, 0)
	}
}

func (this *ListBox) SetItemCount(count int) bool {
	return w32.SendMessage(this.hwnd, w32.LB_SETCOUNT, uintptr(count), 0) != 0
}

func (this *ListBox) ItemCount() int {
	return int(w32.SendMessage(this.hwnd, w32.LB_GETCOUNT, 0, 0))
}

func (this *ListBox) AddItem(text ...string) {
	if len(text) > 0 {
		var li w32.LVITEM
		li.Mask = w32.LVIF_TEXT
		li.PszText = syscall.StringToUTF16Ptr(text[0])
		li.IItem = int32(this.ItemCount())

		this.InsertLvItem(&li)

		// for i := 1; i < len(text); i++ {
		// 	li.PszText = syscall.StringToUTF16Ptr(text[i])
		// 	li.ISubItem = int32(i)

		// 	this.SetLvItem(&li)
		// }
	}
}

func (this *ListBox) InsertLvItem(lvItem *w32.LVITEM) {
	w32.SendMessage(this.hwnd, w32.LB_INSERTSTRING, 0, uintptr(unsafe.Pointer(lvItem)))
}

func (this *ListBox) SetLvItem(lvItem *w32.LVITEM) {
	w32.SendMessage(this.hwnd, w32.LB_ADDSTRING, 0, uintptr(unsafe.Pointer(lvItem)))
}

func (this *ListBox) DeleteAllItems() bool {
	return w32.SendMessage(this.hwnd, w32.LB_DELETESTRING, 0, 0) == w32.TRUE
}

func (this *ListBox) Item(item *w32.LVITEM) bool {
	return w32.SendMessage(this.hwnd, w32.LB_GETSELITEMS, 0, uintptr(unsafe.Pointer(item))) == w32.TRUE
}

func (this *ListBox) ItemAtIndex(i int) *w32.LVITEM {
	var item w32.LVITEM
	item.Mask = w32.LVIF_PARAM | w32.LVIF_TEXT
	item.IItem = int32(i)

	this.Item(&item)
	return &item
}

// mask is used to set the LVITEM.Mask for ListBox.GetItem which indicates which attributes you'd like to receive
// of LVITEM.
func (this *ListBox) SelectedItems(mask uint) []*w32.LVITEM {
	items := make([]*w32.LVITEM, 0)

	var i int = -1
	for {
		if i = int(w32.SendMessage(this.hwnd, w32.LB_GETSELITEMS, uintptr(i), uintptr(w32.LVNI_SELECTED))); i == -1 {
			break
		}

		var item w32.LVITEM
		item.Mask = uint32(mask)
		item.IItem = int32(i)
		if this.Item(&item) {
			items = append(items, &item)
		}
	}
	return items
}

func (this *ListBox) SelectedCount() uint {
	return uint(w32.SendMessage(this.hwnd, w32.LB_GETCOUNT, 0, 0))
}

// Set i to -1 to select all items.
func (this *ListBox) SetSelectedItem(i int) {
	this.setItemState(i, w32.LVIS_SELECTED, w32.LVIS_SELECTED)
}

// Event publishers
// func (this *ListBox) OnEndLabelEdit() *EventManager {
// 	return &this.onEndLabelEdit
// }

func (this *ListBox) OnDBLClick() *EventManager {
	return &this.onDBLClick
}

func (this *ListBox) OnClick() *EventManager {
	return &this.onClick
}

// Message processer
func (this *ListBox) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case w32.WM_NOTIFY:
		nm := (*w32.NMHDR)(unsafe.Pointer(lparam))
		switch int(nm.Code) {
		case w32.LVN_BEGINLABELEDITW:
			// println("Begin label edit")
		// case w32.LVN_ENDLABELEDITW:
		// 	nmdi := (*w32.NMLVDISPINFO)(unsafe.Pointer(lparam))
		// 	if nmdi.Item.PszText != nil {
		// 		this.onEndLabelEdit.Fire(NewEventArg(this, &LVEndLabelEditEventData{Item: &nmdi.Item}))
		// 		return w32.TRUE
		// 	}
		case w32.NM_DBLCLK:
			nmItem := (*w32.NMITEMACTIVATE)(unsafe.Pointer(lparam))
			this.onDBLClick.Fire(NewEventArg(this, &LVDBLClickEventData{NmItem: nmItem}))
		case w32.NM_CLICK:
			nmItem := (*w32.NMITEMACTIVATE)(unsafe.Pointer(lparam))
			this.onClick.Fire(NewEventArg(this, &LVDBLClickEventData{NmItem: nmItem}))
		}
	}

	return this.W32Control.WndProc(msg, wparam, lparam)
}
