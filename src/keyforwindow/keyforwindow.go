package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

/*
Пример, который я привел выше, использует внешние команды xprop и xwininfo для получения идентификатора и имени активного окна соответственно. В то время как предыдущие примеры использовали библиотеки X11 для работы с окнами непосредственно из Go кода.

Использование внешних команд может быть не так эффективно, как работа с X11 непосредственно из Go, и может приводить к дополнительным накладным расходам на запуск и обработку вывода команд.

Кроме того, для работы этого примера необходимо наличие установленных на системе команд xprop и xwininfo, что может быть проблемой на некоторых системах.

Тем не менее, этот пример все еще может быть полезным, если вы не хотите использовать библиотеки X11 и вам необходимо получить имя активного окна на Linux без дополнительных настроек.
*/
func main() {
	for {
		time.Sleep(500 * time.Millisecond)

		// Получить идентификатор активного окна
		active, err := getActiveWindowID()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Получить имя активного окна
		name, err := getWindowName(active)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Active window:", name, active)
	}

}

func getActiveWindowID() (string, error) {
	// Получить идентификатор активного окна
	out, err := exec.Command("xprop", "-root", "_NET_ACTIVE_WINDOW").Output()
	if err != nil {
		return "", fmt.Errorf("xprop failed")
	}

	// Обработать вывод команды и вернуть идентификатор окна
	re := regexp.MustCompile(`0x[0-9a-f]+`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) < 1 {
		return "", fmt.Errorf("active window not found")
	}
	return matches[0], nil
}

func getWindowName(id string) (string, error) {
	// Получить имя окна с помощью команды xwininfo
	out, err := exec.Command("xwininfo", "-id", id).Output()
	if err != nil {
		return "", fmt.Errorf("xwininfo failed")
	}

	// Обработать вывод команды и вернуть имя окна
	re := regexp.MustCompile(`"(.*)"`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) < 2 {
		return "", fmt.Errorf("window name not found")
	}
	return strings.TrimSpace(matches[1]), nil
}
