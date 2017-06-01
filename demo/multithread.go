package main

import (
	"strconv"
	"time"

	"github.com/jx3fans/gform"
)

var (
	pb  *gform.ProgressBar
	btn *gform.PushButton
	lb  *gform.Label
)

func onclick(arg *gform.EventArg) {
	go setProgress()
}

func setProgress() {
	btn.SetEnabled(false)
	for i := 0; i < 100; i++ {
		pb.SetValue(uint32(i))
		lb.SetCaption("Done: " + strconv.Itoa(i) + "%")
		time.Sleep(50 * 1E6)
	}
	btn.SetEnabled(true)
	pb.SetValue(0)
}

func main() {
	gform.Init()

	mw := gform.NewForm(nil)
	mw.SetSize(360, 170)
	mw.SetCaption("Progress bar")
	mw.EnableMaxButton(false)
	mw.EnableSizable(false)
	mw.OnClose().Bind(func(arg *gform.EventArg) {
		gform.Exit()
	})

	lb = gform.NewLabel(mw)
	lb.SetPos(21, 10)
	lb.SetSize(300, 25)
	lb.SetCaption("Installing...")

	pb = gform.NewProgressBar(mw)
	pb.SetPos(20, 35)
	pb.SetSize(300, 25)

	btn = gform.NewPushButton(mw)
	btn.SetPos(220, 80)
	btn.SetCaption("Run")
	btn.OnLBUp().Bind(onclick)

	mw.Show()
	mw.Center()

	gform.RunMainLoop()
}
