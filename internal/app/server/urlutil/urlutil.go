package urlutil

import (
	"path"
	"strings"
)

// ShiftPath https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
