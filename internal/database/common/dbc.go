package database

import (
	"common_notify_server/internal/utils"
	"database/sql"
)

func RowClose(row *sql.Rows) {
	utils.CheckErrors(row.Close())
}
