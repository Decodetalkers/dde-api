//Icon theme thumbnail generator
package icon

import (
	"fmt"
	"os"
	"path"

	. "pkg.deepin.io/dde/api/thumbnails/loader"
	"pkg.deepin.io/lib/mime"
	dutils "pkg.deepin.io/lib/utils"
)

var themeThumbDir = path.Join(os.Getenv("HOME"),
	".cache", "thumbnails", "appearance")

func init() {
	for _, ty := range SupportedTypes() {
		Register(ty, genIconThumbnail)
	}
}

func SupportedTypes() []string {
	return []string{
		mime.MimeTypeIcon,
	}
}

// GenThumbnail generate icon theme thumbnail
// src: the uri of icon theme index.theme
func GenThumbnail(src, bg string, width, height int, force bool) (string, error) {
	if width <= 0 || height <= 0 {
		return "", fmt.Errorf("Invalid width or height")
	}

	ty, err := mime.Query(src)
	if err != nil {
		return "", err
	}

	if ty != mime.MimeTypeIcon {
		return "", fmt.Errorf("Not supported type: %v", ty)
	}

	return genIconThumbnail(src, bg, width, height, force)
}

func ThumbnailForTheme(src, bg string, width, height int, force bool) (string, error) {
	if width <= 0 || height <= 0 {
		return "", fmt.Errorf("Invalid width or height")
	}

	return doGenThumbnail(src, dutils.DecodeURI(bg), width, height, force, true)
}

func genIconThumbnail(src, bg string, width, height int, force bool) (string, error) {
	return doGenThumbnail(src, dutils.DecodeURI(bg), width, height, force, false)
}

func getThumbDest(src string, width, height int, theme bool) (string, error) {
	dest, err := GetThumbnailDest(src, width, height)
	if err != nil {
		return "", err
	}

	if theme {
		dest = path.Join(themeThumbDir, "icon-"+path.Base(dest))
	}
	return dest, nil
}
