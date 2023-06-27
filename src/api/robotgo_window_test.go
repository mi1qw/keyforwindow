package api

import (
	"log"
	"testing"
)

/*
примеры заголовков окон

keyforwindow – Bbb.go

# Double Commander

go1 - [~/Java/go/go1] - SmartGit 21.2.4 Evaluation until Jul 9

jazyk_programmirovanija_go_2016.pdf  — Okular

~/Java/projects/conspekts/800.txt - Sublime Text
~/Java/projects/conspekts/Go.go - Sublime Text

ChatGPT Proxy и еще 196 страниц — Профиль 1: Microsoft Edge
WineHQ aimhero - Google Search и еще 196 страниц — Профиль 1: Microsoft Edge
*/
func TestHello(t *testing.T) {
	window := WindowData{}
	title := window.getTitle()
	log.Default().Println(title)

	expected := "Hello, world!"
	actual := "Hello, world!"
	//actual := hello("world")
	if actual != expected {
		t.Errorf("ошибка: получено '%s', ожидалось '%s'", actual, expected)
	}
}
