package dm

import (
	"github.com/rogeecn/gdm"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Bind struct {
	game *DM

	mouse  []string
	keypad []string
	screen []string
	public []string
	mode   int
}

func (g *DM) Bind() *Bind {
	return &Bind{
		game:   g,
		screen: []string{gdm.DISPLAY_NORMAL},
		mouse:  []string{gdm.MOUSE_NORMAL},
		keypad: []string{gdm.KEYPAD_NORMAL},
		public: []string{},
		mode:   gdm.MODE_0,
	}
}

func (g *DM) UnBindWindow() bool {
	log.Debug("UnbindWindow")
	return g.DmSoft.UnBindWindow()
}

func (b *Bind) Mouse(str string, replace bool) *Bind {
	if len(str) == 0 {
		return b
	}

	if replace {
		b.mouse = strings.Split(str, "|")
		return b
	}
	strs := strings.Split(str, "|")
	for _, s := range strs {
		b.mouse = append(b.mouse, s)
	}

	return b
}

func (b *Bind) KeyPad(str string, replace bool) *Bind {
	if len(str) == 0 {
		return b
	}

	if replace {
		b.keypad = strings.Split(str, "|")
		return b
	}
	strs := strings.Split(str, "|")
	for _, s := range strs {
		b.keypad = append(b.keypad, s)
	}

	return b
}

func (b *Bind) Screen(str string, replace bool) *Bind {
	if len(str) == 0 {
		return b
	}

	if replace {
		b.screen = strings.Split(str, "|")
		return b
	}

	strs := strings.Split(str, "|")
	for _, s := range strs {
		b.screen = append(b.screen, s)
	}

	return b
}

func (b *Bind) Public(str string, replace bool) *Bind {
	if len(str) == 0 {
		return b
	}

	if replace {
		b.public = strings.Split(str, "|")
		return b
	}

	strs := strings.Split(str, "|")
	for _, s := range strs {
		b.public = append(b.public, s)
	}

	return b
}

func (b *Bind) Mode(mode int) *Bind {
	b.mode = mode
	return b
}

func (b *Bind) Do() bool {
	log.Debugf("Bind::Do() Hwnd: %d, Screen: %s, Mouse: %s, Keypad: %s, Mode: %d", b.game.HWND,
		strings.Join(b.screen, "|"),
		strings.Join(b.mouse, "|"),
		strings.Join(b.keypad, "|"),
		b.mode)

	var ret bool
	if len(b.public) > 0 {
		ret = b.game.DmSoft.BindWindowEx(
			b.game.HWND,
			strings.Join(b.screen, "|"),
			strings.Join(b.mouse, "|"),
			strings.Join(b.keypad, "|"),
			strings.Join(b.public, "|"),
			b.mode)
	} else {
		ret = b.game.DmSoft.BindWindow(
			b.game.HWND,
			strings.Join(b.screen, "|"),
			strings.Join(b.mouse, "|"),
			strings.Join(b.keypad, "|"),
			b.mode)
	}
	log.Debugf("bind result: %+v", ret)

	if ret {
		b.game.Bound = ret
	}
	return ret
}
func (b *Bind) DoEx() bool {
	log.Debug("Bind::DoEx() ", b.game.HWND,
		strings.Join(b.screen, "|"),
		strings.Join(b.mouse, "|"),
		strings.Join(b.keypad, "|"),
		b.mode)

	return b.game.DmSoft.BindWindowEx(
		b.game.HWND,
		strings.Join(b.screen, "|"),
		strings.Join(b.mouse, "|"),
		strings.Join(b.keypad, "|"),
		strings.Join(b.public, "|"),
		b.mode)
}
