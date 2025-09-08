package database

import . "shared/option"

func OptionalToSQL(opt Option[string]) string {
	return opt.UnwrapOr("NULL")
}
