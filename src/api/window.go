package api

import (
	"bytes"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/robotn/xgb/xproto"
	"github.com/robotn/xgbutil"
	"github.com/robotn/xgbutil/ewmh"
	"github.com/robotn/xgbutil/xgraphics"
	"github.com/robotn/xgbutil/xprop"
	"github.com/robotn/xgbutil/xwindow"
	"image/color"
	"strings"
)

var X *xgbutil.XUtil
var RedColor = "c75450"

func init() {
	XUtil, err := xgbutil.NewConn()
	if err != nil {
		panic(err)
	}
	X = XUtil
}

func GetWin() *xwindow.Window {
	activeWin, err := ewmh.ActiveWindowGet(X)
	if err != nil {
		panic(err)
	}
	win := xwindow.New(X, activeWin)
	//name, _ := ewmh.WmNameGet(X, activeWin)
	return win
}

// IsWinClass Получить имя класса текущего активного окна.
func IsWinClass(name []byte) bool {
	win := GetWin()
	nameProp, err := xprop.GetProperty(X, win.Id, "WM_CLASS")
	if err != nil {
		//log.Fatalf("Could not get process name: %s", err)
		return false
	}
	class := nameProp.Value
	//fmt.Printf("Window process name: %s\n", class)
	return bytes.HasPrefix(class, name)
}

func (b *Builder) findFuncByWind1() func(event hook.Event) {
	win := GetWin()
	nameProp, err := xprop.GetProperty(X, win.Id, "WM_CLASS")
	if err != nil {
		//log.Fatalf("Could not get process name: %s", err)
		return nil
	}
	class := nameProp.Value
	class = class[:len(class)/2-1]
	//fmt.Printf("Window process name: %s\n", class)
	return b.make[string(class)]
}

type RGBA struct {
	R, G, B, A uint8
}

// RedBtnColor c75450
var RedBtnColor = RGBA{
	R: 0xc7,
	G: 0x54,
	B: 0x50,
	A: 0xFF,
}

func (r *RGBA) equal(color color.RGBA) bool {
	return r.R == color.R &&
		r.G == color.G &&
		r.B == color.B &&
		r.A == color.A
}

func (r *RGBA) stringHEX() string {
	return fmt.Sprintf("%x%x%x", r.R, r.B, r.G)
}

func CheckBtn() bool {

	//for {
	//	win1 := GetWin()
	//	FindColor(1200, 100, 10, 10, RedBtnColor, 3, win1)
	//}

	//time.Sleep(500 * time.Millisecond)

	win := GetWin()
	x, y, w, _ := Geometry(win)
	//FindColor(x+w-280, y+57, 100, 100, RedBtnColor, 3, win)
	b, n := checkColor(x+w-280, y+57, 270, 100, RedColor, 3)
	if n > 1 { // если колличество попыток больше чем ...
		println(b, n, "checkColor")
	}
	return b
}

func Geometry(win *xwindow.Window) (int, int, int, int) {
	geom, err := xwindow.New(X, win.Id).DecorGeometry()
	if err != nil {
		panic(err)
	}
	return geom.X(), geom.Y(), geom.Width(), geom.Height()
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

// FindColor x, расстояие справа   y, расстояние сверху, в % от нчала окна
func FindColor(x, y, w, h int, color RGBA, step int, win *xwindow.Window) {
	colorStr := color.stringHEX()
	var n = 1
	fmt.Println(x, y, w, h, color, step)
	for y_ := y; y_ < y+h; y_ += step {
		for x_ := x; x_ < x+w; x_ += step {
			pixelColor := robotgo.GetPixelColor(x_, y_)
			//pixelColor := Color(x_, y_, win)
			fmt.Printf("")
			//if pixelColor == colorStr {
			if strings.HasPrefix(pixelColor, colorStr[:2]) {
				//if pixelColor == color {
				//if color.equal(pixelColor) {
				fmt.Printf("x= %d -%d  \t y= %d +%d   %.2f%%  %.2f%%  %s\n",
					x_, w-(x_-x), y_, y_-y,
					100*float32(x_-x)/float32(w),
					100*float32(y_-y)/float32(h),
					pixelColor,
				)
			}
			n++
		}
	}
}

func Color(x, y int, win *xwindow.Window) color.RGBA {
	im, err := xgraphics.NewDrawable(X, xproto.Drawable(win.Id))
	if err != nil {
		panic(err)
	}
	//r, g, b, a := im.At(x, y).RGBA()
	//println(r&0xff, g&0xff, b&0xff, a&0xff)

	i := im.PixOffset(x, y)
	rgba := color.RGBA{
		R: im.Pix[i+2],
		G: im.Pix[i+1],
		B: im.Pix[i],
		A: im.Pix[i+3],
	}
	//im.Set(x, y, color.RGBA{200, 100, 100, 100})
	//fmt.Println(rgba.R, rgba.G, rgba.B, rgba.A)
	return rgba
}
