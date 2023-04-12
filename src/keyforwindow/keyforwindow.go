package main

import (
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"log"
	"time"
)

func main() {
	// Connect to the X server.
	X, err := xgb.NewConn()
	if err != nil {
		panic(err)
	}
	defer X.Close()

	for {
		time.Sleep(900 * time.Millisecond)

		// Get the ID of the root window.
		root := xproto.Setup(X).DefaultScreen(X).Root

		// Get the ID of the active window.
		_, err = getActiveWindowID(X, root)
		if err != nil {
			panic(err)
		}
	}

	// Get the geometry of the active window.
	//geometry, err := getGeometry(X, active)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Print the dimensions and position of the active window.
	//fmt.Printf("Active window dimensions: %dx%d\n",
	//	geometry.Width, geometry.Height)
	//fmt.Printf("Active window position: (%d,%d)\n", geometry.X, geometry.Y)
}

func getActiveWindowID(conn *xgb.Conn, root xproto.Window) (xproto.Window, error) {
	// Get the atom id (i.e., intern an atom) of "_NET_WM_NAME".
	aname := "_NET_ACTIVE_WINDOW"
	activeAtom, err := xproto.InternAtom(conn,
		true, uint16(len(aname)),
		aname).Reply()

	// Query the X server for the ID of the active window.
	//reply, err := xproto.GetProperty(conn, false, root,
	//	xproto.Atom("_NET_ACTIVE_WINDOW"),
	//	xproto.Atom("WINDOW"), 0, 4).Reply()
	if err != nil {
		return 0, err
	}

	// Get the atom id (i.e., intern an atom) of "_NET_WM_NAME".
	aname = "_NET_WM_NAME"
	nameAtom, err := xproto.InternAtom(conn, true, uint16(len(aname)),
		aname).Reply()
	if err != nil {
		log.Fatal(err)
	}

	// Get the actual value of _NET_ACTIVE_WINDOW.
	// Note that 'reply.Value' is just a slice of bytes, so we use an
	// XGB helper function, 'Get32', to pull an unsigned 32-bit integer out
	// of the byte slice. We then convert it to an X resource id so it can
	// be used to get the name of the window in the next GetProperty request.
	reply, err := xproto.GetProperty(conn, false, root, activeAtom.Atom,
		xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
	if err != nil {
		log.Fatal(err)
	}
	// Parse the reply to get the active window ID.
	activeWinID := xproto.Window(xgb.Get32(reply.Value))
	fmt.Printf("Active window id: %X\n", activeWinID)

	// Now get the value of _NET_WM_NAME for the active window.
	// Note that this time, we simply convert the resulting byte slice,
	// reply.Value, to a string.
	reply, err = xproto.GetProperty(conn, false, activeWinID,
		nameAtom.Atom,
		xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Active window name: %s\n", string(reply.Value))

	//// Get the atom id (i.e., intern an atom) of "_NET_WM_NAME".
	//aname = "_NET_WM_NAME"
	//nameAtom, err := xproto.InternAtom(conn, true, uint16(len(aname)),
	//	aname).Reply()

	// Query the X server for the geometry of the given window.
	geometry, err := xproto.GetGeometry(conn, xproto.Drawable(activeWinID)).Reply()
	if err != nil {
		log.Fatal(err)
		//return xproto.GetGeometryReply{}, err
	}
	println(geometry.X, geometry.Y, geometry.Width, geometry.Height, geometry.BorderWidth)
	return activeWinID, nil
}

//func getGeometry(conn *xgb.Conn, win xproto.Window) (xproto.GetGeometryReply, error) {
//	// Query the X server for the geometry of the given window.
//	geometry, err := xproto.GetGeometry(conn, win).Reply()
//	if err != nil {
//		return xproto.GetGeometryReply{}, err
//	}
//
//	return *geometry, nil
//}
