package main

import "C"
import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/robotn/xgbutil"
	"image/color"
	"keyforwindow/src/api"
)

var X *xgbutil.XUtil
var nameJavaIDE = "java"
var javaIDE []int32
var redButton = "c75450"
var RedBtnColor = color.RGBA{
	R: 0xc7,
	G: 0x54,
	B: 0x50,
	A: 0xFF,
}

// var raiseVolume = exec.Command("pactl", "set-sink-volume", "0", "+2%")
// var lowerVolume = exec.Command("pactl", "set-sink-volume", "0", "+2%")

func init() {
	XUtil, err := xgbutil.NewConn()
	if err != nil {
		panic(err)
	}
	X = XUtil
}
func main() {

	//api.CheckBtn()

	//x, y, w, h := robotgo.GetBounds(robotgo.GetPID())
	//println(x, y, w, h)
	////start := time.Now().Nanosecond()
	//println(checkColor(x+w-280, y+36, 10, 10, redButton, 3))

	//var duration = time.Now().Nanosecond() - start
	//println("duration ", duration)

	//findColor(x, y, w, h, redButton, 3)  // поиск красного цвета в прямоугольнике

	s := hook.Start()
	defer hook.End()

	window := api.NewBuilder().SetWindow("jetbrains-idea")
	window.Register(hook.KeyDown, []string{"q"},
		func(event hook.Event) {
			fmt.Println("q!!!!!!!!!!!")
		})
	// window.Register(hook.KeyDown, []string{"esc"},
	// 	func(event hook.Event) {
	// 		hook.End()
	// 	})

	//window.RegisterMouse(hook.MouseUp, hook.MouseMap["left"],
	//	func(event hook.Event) {
	//		fmt.Println(event)
	//	})

	// center left right
	window.RegisterMouse(hook.MouseUp, hook.MouseMap["right"], // указательный
		func(event hook.Event) {
			if api.CheckBtn() {
				robotgo.KeyTap("f8")
				fmt.Println(event)
			}
		})

	window.Register(hook.KeyDown, []string{"ctrl", "8"}, // нижняя указательный
		func(event hook.Event) {
			if api.CheckBtn() {
				robotgo.KeyTap("f7")
			}
		})
	//window.Register(hook.KeyDown, []string{"l", "ctrl", "alt"}, // нижняя указательный
	//	func(event hook.Event) {
	//		if api.CheckBtn() {
	//			//robotgo.Sleep(500)
	//			robotgo.KeyTap(robotgo.F9)  // не работает
	//			//robotgo.KeyTap(robotgo.Shift, robotgo.F8)
	//			//robotgo.KeyTap("f8", "shift")
	//
	//			//robotgo.KeyToggle("shift", "down")
	//			//robotgo.KeyToggle("f8", "down")
	//			//robotgo.Sleep(100)
	//			//robotgo.KeyToggle("f8", "up")
	//			//robotgo.KeyToggle("shift", "up")
	//
	//			fmt.Println(event)
	//		}
	//		//else {
	//		//	robotgo.KeyTap("l", "ctrl", "shift")
	//		//}
	//	})

	//WM_CLASS(STRING) = "microsoft-edge", "Microsoft-edge"
	browser := api.NewBuilder().SetWindow("microsoft-edge")
	browser.RegisterMouse(hook.MouseUp, hook.MouseMap["right"], // указательный
		func(event hook.Event) {
			robotgo.KeyTap("w", "ctrl")
			//fmt.Println(event)
		})

	<-hook.Process(s)

	//for n := 1; ; n++ {
	//
	//	println(api.WinClass([]byte("jetbrains-idea")))
	//	time.Sleep(1050 * time.Millisecond)
	//	//println(n, win.Id, name, robotgo.GetPid())
	//	win := api.GetWin()
	//	//fmt.Printf("%d-%[1]x   %d \n", win.Id, robotgo.GetPid())
	//	fmt.Printf("%d-%[1]x \n", win.Id)
	//	//geom := win.Geom
	//	//println("window.geom", geom.X(), geom.Y(), geom.Width(), geom.Height())
	//
	//	//geometryReply, err := xproto.GetGeometry(X.Conn(), xproto.Drawable(win.Id)).Reply()
	//	//if err != nil {
	//	//	panic(err)
	//	//}
	//	//println("proto.GetGeometry", geometryReply.X, geometryReply.Y, geometryReply.Width, geometryReply.Height)
	//
	//	dgeom, err := xwindow.New(X, win.Id).DecorGeometry()
	//	if err != nil {
	//		panic(err)
	//	}
	//	println("DecorGeometry", dgeom.X(), dgeom.Y(), dgeom.Width(), dgeom.Height())
	//
	//	// Получить имя процесса текущего активного окна.
	//	//pidProp, err := xprop.GetProperty(X, win.Id, "_NET_WM_PID")
	//	//if err != nil {
	//	//	log.Fatalf("Could not get process ID: %s", err)
	//	//}
	//	//pid := pidProp.Value
	//
	//	//fmt.Printf("%[1]d  %[1]x", pid)
	//
	//	//nameProp, err := xprop.GetProperty(X, win.Id, "WM_CLASS")
	//	//if err != nil {
	//	//	log.Fatalf("Could not get process name: %s", err)
	//	//}
	//	//procName := nameProp.Value
	//
	//	// Вывести имя процесса текущего активного окна.
	//	//fmt.Printf("Window process ID: %d\n", pid)
	//	//fmt.Printf("Window process name: %s\n", procName)
	//
	//	// Получаем идентификатор процесса, связанного с активным окном
	//	pid1, err := ewmh.WmPidGet(X, win.Id)
	//	if err != nil {
	//		fmt.Println("Error getting process ID:", err)
	//		return
	//	}
	//	// Получаем имя окна по его идентификатору
	//	name, err := ewmh.WmNameGet(X, win.Id)
	//	if err != nil {
	//		fmt.Println("Error getting process name:", err)
	//		return
	//	}
	//	// Выводим имя процесса на экран
	//	fmt.Println("Active process:", name, pid1)
	//
	//}

	//for n := 1; ; n++ {
	//	time.Sleep(250 * time.Millisecond)
	//	println(n, robotgo.GetPid())
	//	//println(robotgo.GetHandle())
	//}
}

func add() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please press w---")
	hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("w")
	})

	s := hook.Start()
	<-hook.Process(s)
}

func low() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		fmt.Println("hook: ", ev)
	}
}

//func VolumeMouse() {
//	robotgo.EventHook(hook.MouseWheel, []string{}, func(e hook.Event) {
//		if e.X < 50 {
//			//fmt.Printf("mouse left @ %v - %v\n", e.X, e.Y)
//			if e.Rotation == -1 {
//				//robotgo.KeyTap("XF86AudioLowerVolume")
//				var raiseVolume = exec.Command("pactl", "set-sink-volume", "0", "+2%")
//				err := raiseVolume.Run()
//				if err != nil {
//					panic(err)
//				}
//			} else {
//				//robotgo.KeyTap("XF86AudioRaiseVolume")
//				var lowerVolume = exec.Command("pactl", "set-sink-volume", "0", "-2%")
//				err := lowerVolume.Run()
//				if err != nil {
//					panic(err)
//				}
//			}
//		}
//	})
//
//}
