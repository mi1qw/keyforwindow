package api

import (
	hook "github.com/robotn/gohook"
)

// m := Make(map[string]func(event hook.Event))
type Builder struct {
	WindowClass []byte
	bind        []string
	key         []string
	//make        map[string]func(event hook.Event)
	make WindEvent
}

type WindEvent map[string]func(event hook.Event)

//type WindEvent struct {
//	WindEvent
//}

func NewBuilder() *Builder {
	return &Builder{}
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
	b.make = w
	cb := func(event hook.Event) {
		wind1 := b.findFuncByWind1()
		if wind1 != nil && event.Button == comand {
			wind1(event)
		}
	}
	hook.Register(when, []string{}, cb)
	return b
}
