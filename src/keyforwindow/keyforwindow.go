package main

import "C"
import (
	hook "github.com/robotn/gohook"
	"os/exec"
)

//var raiseVolume = exec.Command("pactl", "set-sink-volume", "0", "+2%")
//var lowerVolume = exec.Command("pactl", "set-sink-volume", "0", "+2%")

func main() {

	VolumeMouse()
	//for n := 1; ; n++ {
	//	//time.Sleep(250 * time.Millisecond)
	//	//println(n, robotgo.GetPid())

	//println(robotgo.GetHandle())
	//}

	s := hook.Start()
	<-hook.Process(s)
}

func VolumeMouse() {
	hook.Register(hook.MouseWheel, []string{}, func(e hook.Event) {
		if e.X < 50 {
			//fmt.Printf("mouse left @ %v - %v\n", e.X, e.Y)
			if e.Rotation == -1 {
				//robotgo.KeyTap("XF86AudioLowerVolume")
				var raiseVolume = exec.Command("pactl", "set-sink-volume", "0", "+2%")
				err := raiseVolume.Run()
				if err != nil {
					panic(err)
				}
			} else {
				//robotgo.KeyTap("XF86AudioRaiseVolume")
				var lowerVolume = exec.Command("pactl", "set-sink-volume", "0", "-2%")
				err := lowerVolume.Run()
				if err != nil {
					panic(err)
				}
			}
		}
	})

}
