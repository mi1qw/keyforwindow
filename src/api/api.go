package api

import (
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var PAUSE = 250 * time.Millisecond

// m := Make(map[string]func(event hook.Event))
type Builder struct {
	ch          *chan hook.Event
	WindowClass []byte
	debug       bool
	bind        []string
	key         []string
	make        WindEvent
	Last        LastEvent
	left        bool
}

type WindEvent map[string]func(event hook.Event)

type LastEvent struct {
	key   string
	time  time.Time
	event hook.Event
}

func LastEventOf(key string, event hook.Event) LastEvent {
	return LastEvent{key: key,
		time:  time.Now(),
		event: event}
}

func (l LastEvent) Equals(other LastEvent) bool {
	return l.key == other.key
}

//type WindEvent struct {
//	WindEvent
//}

func NewBuilder(ch *chan hook.Event) *Builder {
	builder := Builder{}
	builder.ch = ch
	builder.make = make(map[string]func(event hook.Event))
	builder.debug = false
	return &builder
}

func (b *Builder) DoubleClick(other LastEvent) bool {
	if !other.Equals(b.Last) {
		return false
	}
	if other.time.Sub(b.Last.time) < PAUSE {
		return true
	}
	return false
	//return other.Equals(b.Last) &&
	//	other.time.Sub(b.Last.time) < 500*time.Millisecond
}

func (b *Builder) SetLast(last LastEvent) {
	b.Last = last
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

// принудительно поднимаем, иначе будут повторные нажатия
func keyUpFunc(cmds []string) func() {
	switch len(cmds) {
	case 1:
		return func() { robotgo.KeyTap(cmds[0], "up") }
	case 2:
		return func() { robotgo.KeyTap(cmds[0], "up", cmds[1]) }
	case 3:
		return func() { robotgo.KeyTap(cmds[0], "up", cmds[1], cmds[2]) }
	case 4:
		return func() { robotgo.KeyTap(cmds[0], "up", cmds[1], cmds[2], cmds[3]) }
	default:
		return nil
	}
}

// Register обёртка для func(event hook.Event)
func (b *Builder) Register1(when uint8, cmds []string, w WindEvent) *Builder {
	keys := " " + strings.Join(cmds, " ")
	b.addAll(w, keys)
	keyUp := keyUpFunc(cmds)
	cb := func(event hook.Event) {
		wind1 := b.findFuncByWind1(keys)
		thisEvent := LastEventOf(keys, event)
		if wind1 != nil && !b.DoubleClick(thisEvent) {
			if keyUp != nil {
				keyUp()
			}
			b.SetLast(thisEvent)
			wind1(event)
		}
	}

	hook.Register(when, cmds, cb)
	return b
}

// RegisterMouse обёртка для func(event hook.Event)
func (b *Builder) RegisterMouse(when uint8, comand uint16, f func(event hook.Event)) *Builder {
	cb := func(event hook.Event) {
		if IsWinClass(b.WindowClass) && event.Button == comand {
			f(event)
		}
	}
	hook.Register(when, []string{}, cb)
	return b
}

func (b *Builder) RegisterMouse1(when uint8, comand uint16, w WindEvent) *Builder {
	keys := " " + strconv.Itoa(int(comand))
	b.addAll(w, keys)
	cb := func(event hook.Event) {
		if event.Button == comand {
			wind1 := b.findFuncByWind1(keys)
			thisEvent := LastEventOf(keys, event)
			if wind1 != nil && !b.DoubleClick(thisEvent) {
				wind1(event)
				b.SetLast(thisEvent)
			}
		}
	}
	hook.Register(when, []string{}, cb)
	return b
}
func (b *Builder) RegisterMouseCtrl(when uint8, comand uint16, w WindEvent) *Builder {
	keys := " " + strconv.Itoa(int(comand))
	b.addAll(w, keys)
	cb := func(event hook.Event) {
		if event.Button == comand {
			wind1 := b.findFuncByWind1(keys)
			thisEvent := LastEventOf(keys, event)
			if wind1 != nil && b.HoldClick(hook.MouseMap["left"]) && !b.DoubleClick(thisEvent) {
				b.SetLast(thisEvent)
				wind1(event)
			}
		}
	}
	hook.Register(when, []string{}, cb)
	return b
}
func (b *Builder) HoldClick(comand uint16) bool {
	if comand == b.Last.event.Button &&
		(hook.MouseDown == b.Last.event.Kind ||
			hook.MouseHold == b.Last.event.Kind) {
		return true
	}
	return false
}

func (b *Builder) State() *chan hook.Event {
	events := make(chan hook.Event)
	go func() {
		for {
			ev := <-*b.ch
			if ev.Button == hook.MouseMap["left"] {
				//log.Println(ev)
				if ev.Kind == hook.MouseDown || ev.Kind == hook.MouseHold {
					b.left = true
				} else {
					b.left = false
				}
				b.Last.event = ev
			}
			events <- ev
		}
	}()
	return &events
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
