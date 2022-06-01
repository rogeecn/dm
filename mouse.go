package dm

import (
	"github.com/rogeecn/draw"
	"time"
)

type Mouse struct {
	game *DM
}

func (g *DM) Mouse() *Mouse {
	return &Mouse{game: g}
}

func (m *Mouse) Lock() {
	m.game.Locker.Mouse.Lock()
}

func (m *Mouse) UnLock() {
	m.game.Locker.Mouse.Unlock()
}

func (m *Mouse) Delay(d float32) *Mouse {
	m.game.Sleep(time.Millisecond * time.Duration(d*100))
	return m
}

func (m *Mouse) MoveToWithShake(pt *draw.Point) *Mouse {
	if m.game.paid {
		return m.MoveTo(pt)
	}
	m.MoveTo(pt).Delay(0.5).MoveR(1, 1).Delay(0.5).MoveR(-1, -1).Delay(0.5)
	return m
}

func (m *Mouse) MoveTo(pt *draw.Point) *Mouse {
	m.game.DmSoft.MoveTo(pt)
	return m
}

func (m *Mouse) MoveR(x, y int) *Mouse {
	m.game.DmSoft.MoveR(x, y)
	m.game.Delay()
	return m
}

func (m *Mouse) MoveToR(pt *draw.Point, x, y int) *Mouse {
	pt.Offset(x, y)
	return m.MoveTo(pt)
}

func (m *Mouse) LeftClick() *Mouse {
	m.game.DmSoft.LeftClick()
	m.Delay(2)
	return m
}

func (m *Mouse) LeftClickPoint(pt *draw.Point) *Mouse {
	return m.MoveTo(pt).
		MoveR(1, 1).
		Delay(1).MoveR(-1, -1).
		Delay(1).
		LeftClick().Delay(2)
}

func (m *Mouse) LeftClickRect(r *draw.Rect) *Mouse {
	return m.MoveTo(r.RegionMiddleCenter().Center()).
		Delay(1).MoveR(1, 1).
		Delay(1).MoveR(-1, -1).
		LeftClick()
}

func (m *Mouse) LeftDblClick() *Mouse {
	m.game.DmSoft.LeftDoubleClick()
	return m
}

func (m *Mouse) LeftDblClickPoint(pt *draw.Point) *Mouse {
	return m.MoveToWithShake(pt).LeftDblClick()
}

func (m *Mouse) LeftDown() *Mouse {
	m.game.DmSoft.LeftDown()
	m.game.Delay()
	return m
}

func (m *Mouse) LeftUp() *Mouse {
	m.game.DmSoft.LeftUp()
	m.game.Delay()
	return m
}

func (m *Mouse) RightClick() *Mouse {
	m.game.DmSoft.RightClick()
	m.game.Delay()
	return m
}

func (m *Mouse) RightDown() *Mouse {
	m.game.DmSoft.RightDown()
	m.game.Delay()
	return m
}

func (m *Mouse) RightUp() *Mouse {
	m.game.DmSoft.RightUp()
	m.game.Delay()
	return m
}

func (m *Mouse) RightClickPoint(pt *draw.Point) *Mouse {
	return m.MoveToWithShake(pt).
		RightClick()
}

func (m *Mouse) WheelDown() *Mouse {
	m.game.DmSoft.WheelDown()
	m.game.Delay()
	return m
}

func (m *Mouse) WheelDownTimes(times int) *Mouse {
	for i := 0; i < times; i++ {
		m.game.DmSoft.WheelDown()
		m.game.Delay()
	}

	return m
}

func (m *Mouse) WheelUp() *Mouse {
	m.game.DmSoft.WheelUp()
	m.game.Delay()
	return m
}

func (m *Mouse) Pos() *draw.Point {
	pos, _ := m.game.DmSoft.GetCursorPos()
	return pos
}

func (m *Mouse) Shape() string {
	return m.game.DmSoft.GetCursorShape()
}

func (m *Mouse) ShapeEx() string {
	return m.game.DmSoft.GetCursorShapeEx(1)
}

func (m *Mouse) CompareShape(pt *draw.Point, shape string, f func(pt *draw.Point)) {
	m.CompareShapeInPoints([]*draw.Point{pt}, shape, f, true)
}

func (m *Mouse) CompareShapeInPoints(points []*draw.Point, shape string, f func(pt *draw.Point), br bool) {
	for _, pt := range points {
		m.MoveTo(pt)
		if m.ShapeEx() == shape {
			f(pt)
			if br {
				break
			}
		}
	}
}

func (m *Mouse) CompareShapeInPointsUntil(points []*draw.Point, shape string, f func(pt *draw.Point) bool) {
	for _, pt := range points {
		m.MoveTo(pt).Delay(1)
		if m.ShapeEx() == shape {
			if f(pt) {
				break
			}
		}
	}
}
