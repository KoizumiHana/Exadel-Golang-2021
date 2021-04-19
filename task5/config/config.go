package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlDB() (db *sql.DB, err error) {
	driver := "mysql"
	user := "root"
	pass := "root"
	name := "todo"
	return sql.Open(driver, user+":"+pass+"@/"+name)
}
