package main

import "C"
import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"log"
	"os/exec"
	"time"
)

var nameJavaIDE = "java"
var javaIDE []int32
var redButton = "c75450"

//var raiseVolume = exec.Command("pactl", "set-sink-volume", "0", "+2%")
//var lowerVolume = exec.Command("pactl", "set-sink-volume", "0", "+2%")

func main() {
	javaIDE = JavaIDE(nameJavaIDE)

	for {
		time.Sleep(500 * time.Millisecond)
		println(robotgo.GetPID())
	}

	xMax, yMax := robotgo.GetScreenSize()
	println(xMax, yMax)

	println(robotgo.GetHandle())
	//robotgo.MilliSleep(3000)

	println(TitleOfProcess("java"))
	println(robotgo.GetPID())

	println(CheckWindow())
	//for i := 0; i < 3; i++ {
	//	println(CheckWindow())
	//	robotgo.Sleep(1)
	//}

	robotgo.MilliSleep(500)

	x, y, w, h := robotgo.GetBounds(robotgo.GetPID())
	println(x, y, w, h)
	//start := time.Now().Nanosecond()
	println(checkColor(x+w-280, y+36, 10, 10, redButton, 3))
	//var duration = time.Now().Nanosecond() - start
	//println("duration ", duration)

	//findColor(x, y, w, h, redButton, 3)  // поиск красного цвета в прямоугольнике

	debug()
	addMouse()
	//lowLevel()
	//base()
	//add111()
	VolumeMouse()
	//addMouse()
	addKey()
	println("before <-")

	//cmd := exec.Command("amixer", "-D", "pulse", "sset", "Master", "10%+")
	//err := cmd.Run()
	//if err != nil {
	//	panic(err)
	//}
	/*
		// Подключаемся к звуковому серверу PulseAudio
		server, _ := pulse.NewClient()

		// Получаем список доступных звуковых устройств
		sinks, err := server.ListSinks()

		if err != nil {
			panic(err)
		}

		//Увеличиваем громкость звука на 10% для первого звукового устройства в списке
		_ = sinks[0]
		//volume, err := sink.ID().
		//	panic(err)
		//}
		//volume.Add(0.1)
		//err = sink.SetVolume(volume)
		//if err != nil {
		//	panic(err)
		//}

		// Закрываем соединение с звуковым сервером PulseAudio
		server.Close()
	*/

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
	println("Exit")
}

/*
XF86AudioLowerVolume
XF86AudioRaiseVolume
65299,
65297,
*/
func VolumeMouse() {
	robotgo.EventHook(hook.MouseWheel, []string{}, func(e hook.Event) {
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

// hook listen and return values using detailed examples
func add111() {
	s := hook.Start()
	defer hook.End()

	ct := false
	for {
		i := <-s

		if i.Kind == hook.KeyHold && i.Rawcode == 57 {
			ct = true
			println(ct)
			println("hook: ")
		}
		println(i.Rawcode, i.Rotation, i.Button)
		if i.Rawcode == 32 {
			println("break: ")
			break
		}
	}
}
func base() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		fmt.Println("hook: ", ev)
	}
}
func lowLevel() {
	////////////////////////////////////////////////////////////////////////////////
	// Global event listener
	////////////////////////////////////////////////////////////////////////////////
	fmt.Println("Press q to stop event gathering")
	evChan := robotgo.EventStart()
	for e := range evChan {
		fmt.Println(e)
		if e.Keychar == 'q' {
			robotgo.EventEnd()
			// break
		}
	}
}

/*
keychar 56
rawcode 65464
*/
func addKey() {
	robotgo.EventHook(hook.KeyDown, []string{"8", "ctrl"}, func(e hook.Event) {
		robotgo.KeyTap("f8")
		//fmt.Println("ctrl-shift-q")
		//fmt.Println(e)
	})
	robotgo.EventHook(hook.KeyDown, []string{"q"}, func(e hook.Event) {
		fmt.Println("EventEnd")
		robotgo.EventEnd()
	})
	//robotgo.EventHook(hook.KeyHold, []string{}, func(e hook.Event) {
	//	fmt.Println(e)
	//	if e.Keychar == 56 && e.Rawcode == 65464 {
	//		println("if e.Keychar == 56 && e.Rawcode == 65464")
	//	}
	//})
}

func addMouse() {
	robotgo.EventHook(hook.MouseDown, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["center"] { // правая
			robotgo.KeyTap("f7")
			println("center")
			log.Printf("%d\n",
				robotgo.GetPID())
		} else if e.Button == hook.MouseMap["right"] { // средняя
			robotgo.KeyTap("f9")
			println("right")
		}
	})
}

func debug() {

	for {
		time.Sleep(3 * time.Second)
		//println(time.Now().String())
		println(robotgo.GetTitle())

		//ch := make(chan int32)
		//go getpid(ch)
		//log.Printf("%d\n", <-ch)

		//log.Printf("%d\n",
		//	robotgo.GetPID())

		//log.Printf("   pid=%d  %t \n",
		//	robotgo.GetPID(),
		//	CheckWindow())
		//fmt.Printf("pid=%d  %t \n",
		//	robotgo.GetPID(),
		//	CheckWindow())
	}
}

func getpid(ch chan int32) {
	ch <- robotgo.GetPID()
}

func checkColor(x, y, w, h int, color string, step int) (bool, int) {
	var n = 1
	for y_ := y; y_ < y+h; y_ += step {
		for x_ := x; x_ < x+w; x_ += step {
			pixelColor := robotgo.GetPixelColor(x_, y_)
			if pixelColor == color {
				return true, n
			}
			n++
		}
	}
	return false, n
}

// x, расстояие справа   y, расстояние сверху,   в % от нчала окна
func findColor(x, y, w, h int, color string, step int) {
	println(x, y, w, h, color, step)
	for y_ := y; y_ < y+h-5; y_ += step {
		for x_ := x; x_ < x+w-5; x_ += step {
			pixelColor := robotgo.GetPixelColor(x_, y_)
			if pixelColor == color {
				fmt.Printf("x= %d -%d  \t y= %d +%d   %.2f%%  %.2f%%\n",
					x_, w-(x_-x), y_, y_-y,
					100*float32(x_-x)/float32(w),
					100*float32(y_-y)/float32(h),
				)
			}
		}
	}
}

func CheckWindow() bool {
	pid := robotgo.GetPID()
	for _, i2 := range javaIDE {
		if i2 == pid {
			return true
		}
	}
	return false
}

func JavaIDE(name string) []int32 {
	ids, _ := robotgo.FindIds(name)
	if len(ids) > 0 {
		return ids
	}
	//var int32s []int32
	//return len(ids) > 0?ids:int32s
	return nil
}

func TitleOfProcess(name string) string {
	//fmt.Println(robotgo.GetTitle())
	//ids, _ := robotgo.FindIds("java")
	var title string
	ids, _ := robotgo.FindIds(name)
	if len(ids) > 0 {
		title = robotgo.GetTitle(ids[0])
	}
	return title
}

func AllProcess() {
	names, err := robotgo.FindNames()
	if err != nil {
		fmt.Println(err)
	}
	for i, i2 := range names {
		println(i, i2)
	}
}

func w1() {
	x, y := robotgo.GetMousePos()
	fmt.Println("pos: ", x, y)

	color := robotgo.GetPixelColor(100, 200)
	fmt.Println("color---- ", color)

	fpid, err := robotgo.FindIds("Edge")
	if err == nil {
		fmt.Println("pids... ", fpid)

		if len(fpid) > 0 {
			robotgo.ActivePID(fpid[0])

			//robotgo.Kill(fpid[0])
		}
	}

	robotgo.ActiveName("chrome")

	isExist, err := robotgo.PidExists(100)
	if err == nil && isExist {
		fmt.Println("pid exists is", isExist)

		//robotgo.Kill(100)
	}

	abool := robotgo.ShowAlert("test", "robotgo")
	if abool {
		fmt.Println("ok@@@ ", "ok")
	}

	title := robotgo.GetTitle()
	fmt.Println("title@@@ ", title)

	// Получаем имя каждого окна и выводим его в консоль
	//for _, hwnd := range windows {
	//	title := robotgo.GetWindow(hwnd)
	//	fmt.Printf("Window %d: %s\n", hwnd, title)
	//}
}
