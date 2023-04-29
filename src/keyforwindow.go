package main

import "C"
import (
	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
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

	////start := time.Now().Nanosecond()
	//println(checkColor(x+w-280, y+36, 10, 10, redButton, 3))

	//var duration = time.Now().Nanosecond() - start
	//println("duration ", duration)

	s := hook.Start()
	defer hook.End()

	window := api.NewBuilder(&s).SetWindow("jetbrains-idea")
	events := window.State()
	// center left right
	//WM_CLASS(STRING) = "microsoft-edge", "Microsoft-edge"
	window.Register1(hook.KeyDown, []string{"f1", "ctrl"}, // указательный
		api.WindEvent{
			"jetbrains-idea": func(event hook.Event) {
				if api.CheckBtn(window) {
					robotgo.KeyTap("f9")
					//log.Println("f9")
				} else {
					robotgo.KeyTap("f4", "ctrl")
					//log.Println("f4")
				}
			},
			"microsoft-edge": func(event hook.Event) {
				//robotgo.KeyTap("w", "ctrl")
				robotgo.KeyTap("w", "down", "ctrl")
				//robotgo.MilliSleep(10)
				robotgo.KeyTap("w", "up", "ctrl")
				//log.Println("microsoft-edge", "w + ctrl")
			}})

	window.Register1(hook.KeyDown, []string{"8", "ctrl"}, // нижняя указательный
		api.WindEvent{
			"jetbrains-idea": func(event hook.Event) {
				if api.CheckBtn(window) {
					robotgo.KeyTap("f8")
					//log.Println("f8")
				}
			}})

	window.Register1(hook.KeyDown, []string{"v", "ctrl"},
		api.WindEvent{
			"qterminal": func(event hook.Event) {
				text, _ := clipboard.ReadAll()
				robotgo.TypeStr(text)
				//robotgo.KeyTap("v", "ctrl", "shift") // вставляются ненужные символы
			}})

	window.RegisterMouse1(hook.MouseUp, hook.MouseMap["center"], // правый клик
		api.WindEvent{
			"jetbrains-idea": func(event hook.Event) {
				if api.CheckBtn(window) {
					if window.HoldClick(hook.MouseMap["left"]) {
						robotgo.KeyTap("f8", "shift")
						//log.Println("\"shift\" f8")
					} else {
						robotgo.KeyTap("f7")
						//log.Println("f7")
					}
				}
			}})
	//	window.RegisterMouseCtrl(hook.MouseUp, hook.MouseMap["center"], // правый клик
	//	api.WindEvent{
	//	"jetbrains-idea": func (event hook.Event){
	//	if api.CheckBtn(window){
	//	robotgo.KeyTap("f8", "shift")
	//	log.Println("\"shift\" f8")
	//}
	//}})

	//window.RegisterMouse1(hook.MouseUp, hook.MouseMap["left"], // клик
	//	api.WindEvent{
	//		"jetbrains-idea": func(event hook.Event) {
	//			if api.CheckBtn(window) {
	//				log.Println("MouseUp[\"left\"]")
	//			}
	//		}})
	//window.RegisterMouse1(hook.MouseDown, hook.MouseMap["left"], // клик
	//	api.WindEvent{
	//		"jetbrains-idea": func(event hook.Event) {
	//			if api.CheckBtn(window) {
	//				log.Println("MouseDown[\"left\"]")
	//			}
	//		}})

	// window.Register(hook.KeyDown, []string{"esc"},
	// 	func(event hook.Event) {
	// 		hook.End()
	// 	})

	<-hook.Process(*events)
	//<-hook.Process(s)

	//for n := 1; ; n++ {
	//
	//	println(api.IsWinClass([]byte("jetbrains-idea")))
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
