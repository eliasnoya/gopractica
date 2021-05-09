package model

import (
	"core/config"
	"database/sql"
)

func GetConn() *sql.DB {

	var c config.Config = config.Get()
	var conString string = c.DbUsr + ":" + c.DbPsw + "@" + "tcp(" + c.DbAddress + ":" + c.DbPort + ")/" + c.Dbname
	db, err := sql.Open(c.DbEngine, conString)
	if err != nil {
		panic(err.Error())
	}
	return db

}
