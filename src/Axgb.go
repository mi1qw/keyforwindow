package main

import (
	"fmt"
	"log"

	"github.com/robotn/xgb"
	"github.com/robotn/xgb/xproto"
)

func main() {
	X, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	setup := xproto.Setup(X)
	screen := setup.DefaultScreen(X)

	windowID, err := getActiveWindow(X)
	if err != nil {
		log.Fatal(err)
	}

	clientID, err := getClientID(X, windowID)
	if err != nil {
		log.Fatal(err)
	}

	clientName, err := getClientName(X, clientID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Application name: %s", clientName)
}

// Поиск идентификатора активного окна
func getActiveWindow(X *xgb.Conn) (xproto.Window, error) {
	reply, err := xproto.GetInputFocus(X).Reply()
	if err != nil {
		return 0, err
	}
	return reply.Focus, nil
}

// Получение идентификатора процесса, которому принадлежит окно
func getClientID(X *xgb.Conn, windowID xproto.Window) (xproto.Window, error) {
	reply, err := xproto.GetProperty(X, false, windowID, xproto.Atom("_NET_WM_PID"), xproto.AtomCardinal, 0, 4).Reply()
	if err != nil {
		return 0, err
	}
	if reply.Format != 32 || len(reply.Value) < 4 {
		return 0, fmt.Errorf("некорректный ответ от сервера")
	}
	pid := xproto.Window(xgb.Get32(reply.Value))
	return pid, nil
}

// Получение имени процесса по идентификатору процесса
func getClientName(X *xgb.Conn, clientID xproto.Window) (string, error) {
	reply, err := xproto.GetProperty(X, false, clientID, xproto.Atom("_NET_WM_NAME"), xproto.AtomUTF8String, 0, (1<<32)-1).Reply()
	if err != nil {
		return "", err
	}
	if reply.Format != 8 {
		return "", fmt.Errorf("некорректный формат ответа от сервера")
	}
	return string(reply.Value), nil
}
