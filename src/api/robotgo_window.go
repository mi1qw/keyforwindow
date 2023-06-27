package api

import (
	"github.com/go-vgo/robotgo"
	"strings"
)

type WindowData struct {
}

func (w *WindowData) getTitle() string {
	winTitle := robotgo.GetTitle()
	split := strings.Split(winTitle, " – ")
	println(split)
	return split[1]
}
