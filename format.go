package format

import (
	"fmt"
	"strings"
)

var OrTerm = "or"

func ListOr(names ...string) string {
	return fmt.Sprintf("%s %s %s", strings.Join(names[:len(names)-1], ", "), OrTerm, names[len(names)-1])
}
