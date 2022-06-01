package dm

import (
	"github.com/rogeecn/draw"
	"github.com/rogeecn/gdm/color"
	"github.com/rogeecn/gdm/utils"
	log "github.com/sirupsen/logrus"
	"strings"
	c "tlbb/utils/color"
	"tlbb/utils/rect"
)

func (g *DM) LoadPic(pic []string) *DM {
	g.DmSoft.LoadPic(pic)
	return g
}

type Pic struct {
	game *DM

	pic   []string
	color *color.Colors
	delta string
	rect  *draw.Rect
	sim   float32
	dir   int
}

func (g *DM) Pic() *Pic {
	return &Pic{game: g, color: color.NewColors(), delta: "000000", rect: g.ClientRect(), sim: S10, dir: 0}
}

func (p *Pic) Pic(str ...string) *Pic {
	for _, s := range str {
		if !strings.HasSuffix(s, ".bmp") {
			s = s + ".bmp"
		}
		p.pic = append(p.pic, s)
	}
	return p
}

func (p *Pic) ColorCfg(key string) *Pic {
	p.Color(c.FromCfg(key)...)
	return p
}

func (p *Pic) RectCfg(rStr string) *Pic {
	p.rect = rect.FromCfg(rStr)
	return p
}

func (p *Pic) Color(colors ...*color.Color) *Pic {
	for _, color := range colors {
		p.color.Add(color)
	}
	return p
}

func (p *Pic) Rect(r *draw.Rect) *Pic {
	p.rect = r
	return p
}

func (p *Pic) Dir(dir int) *Pic {
	p.dir = dir
	return p
}
func (p *Pic) Sim(sim float32) *Pic {
	p.sim = sim
	return p
}

func (p *Pic) Find() *utils.FindItemResult {
	if p.game.IsDebug {
		log.Debug("Pic::Find() ", p.rect, p.pic, p.delta, p.sim, p.dir)
	}
	return p.game.DmSoft.FindPic(p.rect, p.pic, p.delta, p.sim, p.dir)
}

func (p *Pic) FindAll() *utils.FindItemsResult {
	if p.game.IsDebug {
		log.Debug("Pic::FindAll() ", p.rect, p.pic, p.delta, p.sim, p.dir)
	}
	return p.game.DmSoft.FindPicEx(p.rect, p.pic, p.delta, p.sim, p.dir)
}

func (p *Pic) FindColor() (*draw.Point, bool) {
	if p.game.IsDebug {
		log.Debug("Pic::FindColor() ", p.rect, p.color, p.sim, p.dir)
	}
	return p.game.DmSoft.FindColor(p.rect, p.color, p.sim, p.dir)
}
