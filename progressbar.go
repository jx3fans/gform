package gform

import (
	"github.com/jx3fans/w32"
)

type ProgressBar struct {
	W32Control
}

func AttachProgressBar(parent Controller, id int32) *ProgressBar {
	prb := new(ProgressBar)
	prb.attach(parent, id)
	RegMsgHandler(prb)

	return prb
}

func AttachProgressBar(parent Controller, id int32) *ProgressBar {
	prb := new(ProgressBar)
	prb.attach(parent, id)
	RegMsgHandler(prb)

	return prb
}

func NewProgressBar(parent Controller) *ProgressBar {
	pb := new(ProgressBar)
	pb.init(parent)

	pb.SetSize(200, 25)

	return pb
}

func (this *ProgressBar) init(parent Controller) {
	this.W32Control.init(w32.PROGRESS_CLASS, parent, 0, w32.WS_CHILD|w32.WS_VISIBLE)
	RegMsgHandler(this)
}

func (this *ProgressBar) Value() uint {
	ret := w32.SendMessage(this.hwnd, w32.PBM_GETPOS, 0, 0)
	return uint(ret)
}

func (this *ProgressBar) SetValue(v uint32) {
	w32.SendMessage(this.hwnd, w32.PBM_SETPOS, uintptr(v), 0)
}

func (this *ProgressBar) Range() (min, max uint) {
	min = uint(w32.SendMessage(this.hwnd, w32.PBM_GETRANGE, uintptr(w32.BoolToBOOL(true)), 0))
	max = uint(w32.SendMessage(this.hwnd, w32.PBM_GETRANGE, uintptr(w32.BoolToBOOL(false)), 0))
	return
}

func (this *ProgressBar) SetRange(min, max uint32) {
	w32.SendMessage(this.hwnd, w32.PBM_SETRANGE32, uintptr(min), uintptr(max))
}
