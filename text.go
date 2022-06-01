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

type Text struct {
	game *DM

	str   []string
	color *color.Colors
	rect  *draw.Rect
	sim   float32
}

func (g *DM) Text() *Text {
	return &Text{
		game:  g,
		color: color.NewColors(),
		rect:  g.ClientRect(),
		sim:   S10,
	}
}

func (o *Text) Str(str string) *Text {
	strs := strings.Split(str, "|")
	for _, s := range strs {
		o.str = append(o.str, s)
	}
	return o
}

func (o *Text) Sim(s float32) *Text {
	o.sim = s
	return o
}

func (o *Text) ColorCfg(key string) *Text {
	o.Colors(c.FromCfg(key))
	return o
}

func (o *Text) Color(colors ...*color.Color) *Text {
	for _, color := range colors {
		o.color.Add(color)
	}
	return o
}

func (o *Text) Colors(colors []*color.Color) *Text {
	for _, color := range colors {
		o.color.Add(color)
	}
	return o
}

func (o *Text) RectCfg(rStr string) *Text {
	o.rect = rect.FromCfg(rStr)
	return o
}

func (o *Text) Rect(r *draw.Rect) *Text {
	o.rect = r
	return o
}

func (o *Text) Ocr() string {
	log.Debug("Text::Ocr() ", o.rect, o.str, o.color, o.sim)
	r := o.game.DmSoft.Ocr(o.rect, o.color, o.sim)
	log.Debugf("Ocr() Result: %s", r)
	return r
}

//func (o *Text) Find1() (*draw.Point, bool) {
//	log.Debug("Text::Find() ", o.rect, o.str, o.color, o.sim)
//	return o.game.DmSoft.FindStr(o.rect, o.str, o.color, o.sim)
//}

func (o *Text) Find() *utils.FindItemResult {
	log.Debug("Text::Find() ", o.rect, o.str, o.color, o.sim)
	return o.game.DmSoft.FindStrE(o.rect, o.str, o.color, o.sim)
}

func (o *Text) FindFast() *utils.FindItemResult {
	log.Debug("Text::FindFast() ", o.rect, o.str, o.color, o.sim)
	return o.game.DmSoft.FindStrFastE(o.rect, o.str, o.color, o.sim)
}

func (o *Text) FindAll() *utils.FindItemsResult {
	log.Debug("Text::FindAll() ", o.rect, o.str, o.color, o.sim)
	return o.game.DmSoft.FindStrEx(o.rect, o.str, o.color, o.sim)
}
