package rbdebug

import (
	"encoding/json"
	"fmt"
	"strings"
)

func D(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		if strings.Contains(err.Error(), "unsupported type") {
			return fmt.Sprintf("go[ %v ]", v)
		}
		return "<ERR: " + err.Error() + ">"
	}
	return string(b)
}
