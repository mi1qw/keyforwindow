package main

import "C"
import (
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"os/exec"
)

func main() {
	xMax, yMax := robotgo.GetScreenSize()
	println(xMax, yMax)

	VolumeMouse()

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
	println("Exit")
}

/*
XF86AudioLowerVolume
XF86AudioRaiseVolume
65299,
65297,

var raiseVolume = exec.Command("pactl", "set-sink-volume", "0", "+2%")
var lowerVolume = exec.Command("pactl", "set-sink-volume", "0", "+2%")
*/
func VolumeMouse() {
	robotgo.EventHook(hook.MouseWheel, []string{}, func(e hook.Event) {
		if e.X < 50 {
			if e.Rotation == -1 {
				var raiseVolume = exec.Command("pactl", "set-sink-volume", "0", "+2%")
				err := raiseVolume.Run()
				if err != nil {
					panic(err)
				}
			} else {
				var lowerVolume = exec.Command("pactl", "set-sink-volume", "0", "-2%")
				err := lowerVolume.Run()
				if err != nil {
					panic(err)
				}
			}
		}
	})
}
