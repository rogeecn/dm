package dm

import (
	"fmt"
	"time"
)

type Keypad struct {
	game *DM
}

func (g *DM) Keypad() *Keypad {
	return &Keypad{game: g}
}

func (m *Keypad) Lock() {
	m.game.Locker.Keypad.Lock()
}

func (m *Keypad) UnLock() {
	m.game.Locker.Keypad.Unlock()
}

func (m *Keypad) Delay(d int) *Keypad {
	m.game.Sleep(time.Millisecond * time.Duration(d*100))
	return m
}

func (k *Keypad) Press(vkey int) *Keypad {
	k.game.DmSoft.KeyPress(vkey)
	return k
}

func (k *Keypad) Down(vkey int) *Keypad {
	k.game.DmSoft.KeyDown(vkey)
	k.game.Delay()

	return k
}

func (k *Keypad) ShortCut(vkeys ...int) *Keypad {
	//defer func() {
	//	for i := len(vkeys) - 1; i >= 0; i-- {
	//		//log.Debugf("up key: %d %d", i, vkeys[i])
	//		k.game.DmSoft.KeyUp(vkeys[i])
	//		k.Delay(2)
	//	}
	//}()

	for _, vkey := range vkeys {
		//log.Debugf("down key: %d", vkey)

		k.game.DmSoft.KeyDown(vkey)
		defer k.game.DmSoft.KeyUp(vkey)
		k.Delay(2)
	}

	return k
}

func (k *Keypad) Up(vkey int) *Keypad {
	k.game.DmSoft.KeyUp(vkey)
	k.game.Delay()
	return k
}

func (k *Keypad) Input(str string) *Keypad {
	k.game.DmSoft.SendString(k.game.HWND, str)
	return k
}

func (k *Keypad) KeyPressStr(str string) *Keypad {
	if !k.game.paid {
		return k.Input(str)
	}
	k.game.DmSoft.KeyPressStr(str, 50)
	return k
}

func (k *Keypad) InputInt(data int) *Keypad {
	return k.Input(fmt.Sprintf("%d", data))
}

func (k *Keypad) InputIme(str string) *Keypad {
	k.game.DmSoft.SendStringIme(str)
	return k
}
