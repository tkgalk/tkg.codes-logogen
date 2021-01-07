package drawer

import (
	"path/filepath"

	"github.com/fogleman/gg"
	"github.com/pkg/errors"
)

// ImageRootPath sets the path to where the fonts are.
var ImageRootPath = filepath.Join("assets", "images")

// DrawImage uses gg to render a specified image at a given location.
// The root folder from where the image is taken is configured by ImageRootPath.
func DrawImage(dc *gg.Context, imagepath string, x, y int) error {
	ip := filepath.Join(ImageRootPath, imagepath)

	image, err := gg.LoadImage(ip)
	if err != nil {
		return errors.Wrap(err, "\n\tfailed drawing an image")
	}

	dc.DrawImage(image, x, y)

	return nil
}
