package database

import (
	"database/sql"
	"nfly/internal/utils"
)

func RowClose(row *sql.Rows) {
	utils.CheckErrors(row.Close())
}

func StmtClose(stmt *sql.Stmt) {
	utils.CheckErrors(stmt.Close())
}
