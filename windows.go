package dm

type Window struct {
	game *DM
	hwnd int
}

func (g *DM) Window() *Window {
	return &Window{
		game: g,
		hwnd: g.HWND,
	}
}

func (w *Window) SetState(hwnd, state int) bool {
	return w.game.DmSoft.SetWindowState(hwnd, state)
}
