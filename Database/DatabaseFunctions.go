package Database

import (
	"database/sql"
	"fmt"
	"log"
)

func DBConnect() *sql.DB {
	conn, err := sql.Open("mysql", DBAddress)
	if err != nil {
		log.Fatal(err.Error())
	}
	return conn
}

func SelectFromDB(Column, Table, WhereColumn, WhereValue string, db *sql.DB) *sql.Rows {
	if WhereColumn == "1" {
		if Column == "*" {
			Query, err := db.Query("SELECT " + Column + " FROM `" + Table + "` WHERE 1")
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println(Column, Table)
			}
			return Query
		} else {
			Query, err := db.Query("SELECT `" + Column + "` FROM `" + Table + "` WHERE 1")
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println(Column, Table)
			}
			return Query
		}

	} else {
		if Column == "*" {
			Query, err := db.Query("SELECT " + Column + " FROM `" + Table + "` WHERE `" + WhereColumn + "` = \"" + WhereValue + "\"")

			if err != nil {
				fmt.Println(err.Error())
				fmt.Println(Column, Table)
			}
			return Query
		} else {
			Query, err := db.Query("SELECT `" + Column + "` FROM `" + Table + "` WHERE `" + WhereColumn + "` = \"" + WhereValue + "\"")
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println(Column, Table)
			}
			return Query
		}
	}
}
