package cycle

import (
	"github.com/djlongarms/leftpad"
)
var DEFAULT_CHAR = ' '

func FormatDouble(s string, i int) string {
	return leftpad.Format(s + s, i)
}