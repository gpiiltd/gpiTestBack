package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gpitest/util"
)

var Conn *sql.DB
var ConnC *sql.DB

//TODO panic when conn is nil
func init() {
	conn, err := sql.Open("mysql", util.DB.Db_user + ":" + util.DB.Db_password+
		"@tcp("+ util.DB.Db_host+ ":"+ util.DB.Db_port+ ")/"+ util.DB.Db_database+ "?charset=utf8")
	if err != nil {
		fmt.Println("error while connect to mysql:")
		fmt.Println(err)
	}
	conn.SetMaxOpenConns(util.DB.Db_max_open)
	conn.SetMaxIdleConns(util.DB.Db_max_idle)
	fmt.Println("conn created:")
	fmt.Println(conn)
	Conn = conn

	connc, err := sql.Open("mysql", util.CDB.Db_user + ":" + util.CDB.Db_password+
		"@tcp("+ util.CDB.Db_host+ ":"+ util.CDB.Db_port+ ")/"+ util.CDB.Db_database+ "?charset=utf8")
	if err != nil {
		fmt.Println("error while connect to mysql:")
		fmt.Println(err)
	}
	connc.SetMaxOpenConns(util.DB.Db_max_open)
	connc.SetMaxIdleConns(util.DB.Db_max_idle)
	fmt.Println("ConnC created:")
	fmt.Println(connc)
	ConnC = connc
}
