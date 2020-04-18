package mysqlStu

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func connect() {
	db, err := sql.Open("mysql", "mysqluser:mysqluser@tcp(127.0.0.1:3306)/test")
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if query, err := db.Query("SELECT * FROM `lqh_two`"); err != nil {
		panic(err)
	} else {
		columns, _ := query.Columns()
		fmt.Println(columns)
		for query.Next() {
			var (
				id       int
				name     string
				age      int
				proctime string
			)
			_ = query.Scan(&id, &name, &age, &proctime)
			fmt.Printf("Id: %d, Name: %s, Age: %d, Time: %s \n", id, name, age, proctime)
		}
	}

}
