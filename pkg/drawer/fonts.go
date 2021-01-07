package drawer

import (
	"path/filepath"

	"github.com/fogleman/gg"
	"github.com/pkg/errors"
)

// FontRootPath sets the path to where the fonts are.
var FontRootPath = filepath.Join("assets", "fonts")

func loadFont(dc *gg.Context, fontname string) error {
	fontpath := filepath.Join(FontRootPath, fontname)
	if err := dc.LoadFontFace(fontpath, 56); err != nil {
		return errors.Wrap(err, "\n\tfailed loading font")
	}

	return nil
}

// DrawText draws a string s with the color c in with a given fontname at the given lMargin while wrapping it at rMargin.
// Font root path is taken from FontRootPath variable.
func DrawText(dc *gg.Context, s, c, fontname string, lMargin, rMargin, yOffset float64) error {
	err := loadFont(dc, fontname)
	if err != nil {
		return errors.Wrap(err, "\n\tfailed drawing text")
	}

	y := float64(dc.Height())/2.0 + yOffset
	maxWidth := float64(dc.Width()) - lMargin - rMargin

	dc.SetHexColor(c)
	dc.DrawStringWrapped(s, lMargin, y, 0, 1, maxWidth, 1.5, gg.AlignLeft)

	return nil
}
