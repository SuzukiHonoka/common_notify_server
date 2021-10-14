package res

const (
	CreateTableIfNotExist = "CREATE TABLE IF NOT EXISTS user(id integer primary key, email text, password text, " +
		"gid integer)"
	SelectAllUserFromDB = "SELECT email,password,gid FROM user"
	InsertUser          = "INSERT INTO user(email,password,gid) values(?,?,?)"
)
