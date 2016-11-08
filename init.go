package gform

import (
	"github.com/Ribtoks/winapi4go"
)

func init() {
	gControllerRegistry = make(map[w32.HWND]Controller)
	gRegisteredClasses = make([]string, 0)

	var si w32.GdiplusStartupInput
	si.GdiplusVersion = 1
	w32.GdiplusStartup(&si, nil)
}
