package go_task

import (
	"database/sql"

	"fmt"

	"os"

	//"strings"

	"time"

	_ "github.com/mattn/go-sqlite3"
)

type appContext struct {
	db *sql.DB
}

//func add_TABLE()

// 检查文件或目录是否存在// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false

func Exist(filename string) bool {

	_, err := os.Stat(filename)

	return err == nil || os.IsExist(err)

}

func SQLite_open(name string) (*appContext, string) { //打开数据库

	if Exist(name) {

		fmt.Println("file SQLite ok")

		//os.Remove("abc.db") //删除文件

		db, err := sql.Open("sqlite3", name)

		if err != nil {

			return nil, err.Error()

		}

		if err = db.Ping(); err != nil {

			return nil, err.Error()

		}

		return &appContext{db}, ""

		//defer db.Close()

	} else {

		fmt.Println("file SQLite no")

		os.Create(name) //创建文件

		//////////////////////

		db, err := sql.Open("sqlite3", name)

		if err != nil {

			return nil, err.Error()

		}

		if err = db.Ping(); err != nil {

			return nil, err.Error()

		}

		//defer db.Close()

		///////////////////////////////////////////

		sqlStmt := `

                    CREATE TABLE "log" (

                    "id"  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,

                    "name"  TEXT(100) NOT NULL,

                    "megbox"  TEXT(100) NOT NULL,

                    "time"  TEXT(100) NOT NULL

                    );

                    `

		_, err = db.Exec(sqlStmt) //创建表

		if err != nil {

			fmt.Printf("%q: %s\n", err, sqlStmt)

			return nil, err.Error()

		}

		defer db.Close()

		/////////////////////////////////////////////////////////////

		dbx, err := sql.Open("sqlite3", name)

		if err != nil {

			return nil, err.Error()

		}

		if err = dbx.Ping(); err != nil {

			return nil, err.Error()

		}

		return &appContext{dbx}, ""

	}

}

func (c *appContext) SQLite_Exec(sql string) {

	_, err := c.db.Exec(sql)

	if err != nil {

		fmt.Printf("add error: %v", err)

		return

	}

}

func Time_Unix() string { //获取时间戳

	time_data := fmt.Sprintf("%d", time.Now().Unix())

	return time_data

}

func temp() {
	c, err := go_task.SQLite_open("abc.db")

	if err != "" {

		print(err)

	}

	sql := fmt.Sprintf("insert into log(name,megbox,time) VALUES('111','222','%s')", go_task.Time_Unix())

	c.SQLite_Exec(sql)

}
