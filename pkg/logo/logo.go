package logo

import (
	"errors"
	"github.com/fogleman/gg"
	"github.com/galkowskit/generate-cover-image/pkg/drawer"
)

// Width is the whole rectangle width.
const Width = 1280

// Height is the whole rectangle height.
const Height = 668

// TextRightMargin tells how much (in pixels) the text will end at from the right side.
const TextRightMargin = 60.0

// DrawLogo draws the actual tkg.codes logo.
func DrawLogo(title, subtitle string) error {
	dc := gg.NewContext(Width, Height)

	dc.DrawRectangle(0, 0, Width, Height)
	dc.SetHexColor("#FFFFFF")
	dc.Fill()

	err := drawTitle(dc, title)
	if err != nil {
		return errors.Wrap(err, "draw title error")
	}

	err = drawSubtitle(dc, subtitle)
	if err != nil {
		return errors.Wrap(err, "draw subtitle error")
	}

	err = drawLogo(dc)
	if err != nil {
		return errors.Wrap(err, "draw image error")
	}

	err = drawRectangles(dc)
	if err != nil {
		return errors.Wrap(err, "draw rectangles error")
	}

	err = dc.SavePNG("output.png")
	if err != nil {
		return errors.Wrap(err, "save png error")
	}

	return nil
}

func drawLogo(dc *gg.Context) error {
	err := drawer.DrawImage(dc, "logo.png", 80, 210)
	if err != nil {
		return errors.Wrap(err, "\n\tfailed drawing logo image")
	}

	return nil
}

func drawRectangles(dc *gg.Context) error {
	err := drawer.DrawImage(dc, "bottom-rects.png", 145, dc.Height()-125)
	if err != nil {
		return errors.Wrap(err, "\n\tfailed drawing bottom rectangles image")
	}

	err = drawer.DrawImage(dc, "top-rects.png", 0, 0)
	if err != nil {
		return errors.Wrap(err, "\n\tfailed drawing top rectangles image")
	}

	return nil
}

func drawTitle(dc *gg.Context, s string) error {
	err := drawer.DrawText(dc, s, "#5A67D8", "Ubuntu-Medium.ttf", 480.0, TextRightMargin, 0.0, 1)
	if err != nil {
		return errors.Wrap(err, "\n\tfailed drawing title")
	}

	return nil
}

func drawSubtitle(dc *gg.Context, s string) error {
	err := drawer.DrawText(dc, s, "#000000", "Ubuntu-Light.ttf", 480.0, TextRightMargin, 45.0, 0)
	if err != nil {
		return errors.Wrap(err, "\n\tfailed drawing subtitle")
	}

	return nil
}
