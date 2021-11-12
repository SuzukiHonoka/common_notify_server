package database

import (
	"common_notify_server/internal/utils"
	"database/sql"
)

func RowClose(row *sql.Rows) {
	utils.CheckErrors(row.Close())
}

func StmtClose(stmt *sql.Stmt) {
	utils.CheckErrors(stmt.Close())
}
