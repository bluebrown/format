package format

import (
	"fmt"
	"strings"
)

func ListOr(names ...string) string {
	return fmt.Sprintf("%s or %s", strings.Join(names[:len(names)-1], ", "), names[len(names)-1])
}
