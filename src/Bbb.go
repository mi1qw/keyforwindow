package main

import "C"
import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	for n := 0; ; n++ {
		time.Sleep(250 * time.Millisecond)

		// Получение активного окна
		mdata := robotgo.GetActive()

		// Получение идентификатора процесса текущего окна
		fmt.Println("ID процесса текущего окна:", mdata.ProcessID)

		// Получение названия окна
		fmt.Println("Название активного окна:", robotgo.GetTitle())

		// Изменение координат мыши в позицию центра активного окна
		//robotgo.MoveMouseSmooth(mdata.X, mdata.Y)

		// Получение размеров активного окна
		//w, h := robotgo.GetWindowSize(mdata.ID)
		//fmt.Println("Размеры окна:", w, h)
		//
		//// Закрытие текущего окна
		//robotgo.CloseWindow(mdata.ID)

		//var cm C.MData
		//cm = robotgo.GetActive()
		//println(cm)
		fmt.Printf("%v | %s | %d \n",
			robotgo.GetPid(),
			robotgo.GetTitle(),
			n)
	}
}
