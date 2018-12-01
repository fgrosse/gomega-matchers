package matchers

import (
	"fmt"
)

type color string

const (
	colorDefault color = "\x1b[0m"
	colorRed     color = "\x1b[31m"
	colorWhite   color = "\x1b[37m"
)

func colored(c color, s string) string {
	return fmt.Sprint(c, s, colorDefault)
}
