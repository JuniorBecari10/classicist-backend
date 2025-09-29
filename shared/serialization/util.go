package serialization

import (
	"fmt"
	. "shared/option"
	"strings"
)

func OptionalToSQL[T any](opt Option[T]) string {
	if val, ok := opt.TryUnwrap(); ok {
		switch v := any(val).(type) {
			case string:
				return fmt.Sprintf("'%s'", v)
			default:
				return fmt.Sprintf("%v", v)
		}
	} else {
		return "NULL"
	}
}

func StringToSQL(s string) string {
    escaped := strings.ReplaceAll(s, "'", "''")
    return fmt.Sprintf("'%s'", escaped)
}
