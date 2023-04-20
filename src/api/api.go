package api

import (
	hook "github.com/robotn/gohook"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// m := Make(map[string]func(event hook.Event))
type Builder struct {
	WindowClass []byte
	debug       bool
	bind        []string
	key         []string
	make        WindEvent
}

type WindEvent map[string]func(event hook.Event)

//type WindEvent struct {
//	WindEvent
//}

func NewBuilder() *Builder {
	builder := Builder{}
	builder.make = make(map[string]func(event hook.Event))
	builder.debug = false
	return &builder
}

func (b *Builder) SetKey(consumer func()) *Builder {
	consumer()
	return b
}

func (b *Builder) SetWindow(name string) *Builder {
	b.WindowClass = []byte(name)
	return b
}

func (b *Builder) SetBind(bind []string) *Builder {
	b.bind = bind
	return b
}

// Register обёртка для func(event hook.Event)
func (b *Builder) Register(when uint8, cmds []string, f func(event hook.Event)) *Builder {
	cb := func(event hook.Event) {
		if IsWinClass(b.WindowClass) {
			f(event)
		}
	}
	hook.Register(when, cmds, cb)
	return b
}

// Register обёртка для func(event hook.Event)
func (b *Builder) Register1(when uint8, cmds []string, w WindEvent) *Builder {
	keys := " " + strings.Join(cmds, " ")
	b.addAll(w, keys)
	cb := func(event hook.Event) {
		wind1 := b.findFuncByWind1(keys)
		if wind1 != nil {
			wind1(event)
		}
	}
	hook.Register(when, cmds, cb)
	return b
}

// RegisterMouse обёртка для func(event hook.Event)
func (b *Builder) RegisterMouse(when uint8, comand uint16,
	f func(event hook.Event)) *Builder {
	cb := func(event hook.Event) {
		if IsWinClass(b.WindowClass) && event.Button == comand {
			f(event)
		}
	}
	hook.Register(when, []string{}, cb)
	return b
}

func (b *Builder) RegisterMouse1(when uint8, comand uint16,
	w WindEvent) *Builder {
	keys := " " + strconv.Itoa(int(comand))
	b.addAll(w, keys)
	cb := func(event hook.Event) {
		if event.Button == comand {
			wind1 := b.findFuncByWind1(keys)
			if wind1 != nil {
				wind1(event)
			}
		}
	}
	hook.Register(when, []string{}, cb)
	return b
}

func (b *Builder) addAll(w WindEvent, key string) {
	for s, f := range w {
		b.make[s+key] = f
	}
}

// OnRightClick включаем настройки
func OnRightClick() {
	log.Println("OnRightClick")
	//cmd := exec.Command("sh", "-c", "xmodmap ~/.Xmodmap.old")
	cmd := exec.Command("sh", "-c",
		"xmodmap -e 'pointer = 1 25 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24'")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// OffRightClick Отключаем правую кнопку мыши
func OffRightClick() {
	log.Println("OffRightClick")
	cmd := exec.Command("sh", "-c",
		"xmodmap -e 'pointer = 1 25 26 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24'")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (b *Builder) SetDebug(btn bool) {
	if btn {
		if !b.debug {
			b.debug = true
			OffRightClick()
		}
	} else {
		if b.debug {
			b.debug = false
			OnRightClick()
		}
	}
}
