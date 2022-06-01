package dm

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rogeecn/draw"
	"github.com/rogeecn/gdm"
	"sync"
	"time"
	"tlbb/utils/plugin"
)

var once sync.Once

type locker struct {
	All    sync.Mutex
	Keypad sync.Mutex
	Mouse  sync.Mutex
}

type DM struct {
	paid bool

	HWND    int
	DmSoft  *gdm.DmSoft
	IsDebug bool
	Bound   bool

	M *Mouse
	K *Keypad

	Locker *locker
}

func New(dmRegDllPath string) (*DM, error) {

	var err error
	dm, err := gdm.New(dmRegDllPath, plugin.CallPluginName, plugin.PluginName())
	if err != nil {
		return nil, err
	}

	//once.Do(func() {
	//	dm.DmGuard(1, plugin.Guard)
	//})

	if err := plugin.PreConfig(dm); err != nil {
		return nil, err
	}

	g := &DM{DmSoft: dm, Locker: &locker{}, paid: plugin.Paid}

	g.M = g.Mouse()
	g.K = g.Keypad()

	return g, nil
}

func (g *DM) SetHWND(hwnd uint32) *DM {
	g.HWND = int(hwnd)
	return g
}

func (g *DM) SetBasePath(path string) *DM {
	g.DmSoft.SetPath(path)
	return g
}

func (g *DM) SetDict(name string) *DM {
	g.DmSoft.SetDefaultDict(name)
	return g
}

func (g *DM) Release() {
	if g.Bound {
		g.UnBindWindow()
	}
	g.DmSoft.Release()
}

func (g *DM) BaseDelay(delta time.Duration) time.Duration {
	return 200*time.Millisecond + delta
}

func (g *DM) Delay() {
	g.Sleep(g.BaseDelay(0))
}
func (g *DM) DelayN(n float32) {
	g.Sleep(time.Duration(n))
}

func (g *DM) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

func (g *DM) GetLastError() error {
	return g.DmSoft.GetLastError()
}

func (g *DM) ClientRect() *draw.Rect {
	if g.Bound {
		clientR := win.HWND(g.HWND).GetClientRect()
		return draw.NewRect(draw.NewSize(int(clientR.Right-clientR.Left), int(clientR.Bottom-clientR.Top)))
	}

	return g.DmSoft.GetClientRect(g.HWND)

}
