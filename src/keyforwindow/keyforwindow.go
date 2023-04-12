package main

import (
	"fmt"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/xwindow"
	"time"
)

func main() {
	X, err := xgbutil.NewConn()
	if err != nil {
		panic(err)
	}
	
	for {
		time.Sleep(500 * time.Millisecond)

		active, err := ewmh.ActiveWindowGet(X)
		if err != nil {
			panic(err)
		}

		name, err := ewmh.WmNameGet(X, active)
		if err != nil {
			panic(err)
		}
		fmt.Println("Active window:", name, active)

		// Get the root window.
		//root := X.RootWin()

		// Get the ID of the active window.
		//activeWin, err := xgbutil.GetActiveWindow(X, root)
		//if err != nil {
		//	panic(err)
		//}

		// Get the geometry of the active window.
		winGeom, err := xwindow.New(X, active).Geometry()
		if err != nil {
			panic(err)
		}

		// Print the geometry of the active window.
		fmt.Printf("Active window geometry: %+v\n", winGeom)

	}

}
