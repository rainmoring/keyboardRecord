package main

import (
	"database/sql"
	"fmt"
	"sync/atomic"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var keydb *sql.DB
var keyArray [VK_BOTTOM]int32

func hookCallback(keyEvent int, keyCode int) {
	atomic.AddInt32(&keyArray[keyCode], 1)
}

func main() {
	var (
		err error
	)
	if keydb, err = sql.Open("sqlite3", "./keycode.db"); err != nil {
		fmt.Println("open keycode.db failed")
		return
	}
	/*
		`

		                    CREATE TABLE "keytable" (

		                    "id"  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,

		                    "name"  TEXT(100) NOT NULL,

		                    "megbox"  TEXT(100) NOT NULL,

		                    "time"  TEXT(100) NOT NULL

		                    );`

	*/
	sqlStmt := `
                CREATE TABLE IF NOT EXISTS keytable(
					uid INTEGER PRIMARY KEY AUTOINCREMENT,
                    id INTEGER ,
                    value INTEGER
                );
				`
	//if _, err = keydb.Exec(`CREATE DATABASE IF NOT EXIST keytable(
	if _, err = keydb.Exec(sqlStmt); err != nil {
		fmt.Println(err.Error())
	}

	if err = keydb.Ping(); err != nil {
		fmt.Println("open keycode.db failed")
		return
	}

	fmt.Println("ok, VK_BOTTOM", VK_BOTTOM)
	for i := 0; i < VK_BOTTOM; i++ {

		var value int32
		fmt.Println("query i is ", i)
		if err = keydb.QueryRow("select value from keytable where id = ?", i).Scan(&value); err != nil {
			fmt.Println(err.Error())
			if stmt, errs := keydb.Prepare("INSERT INTO keytable(id,value) values(?,?)"); errs != nil {

				fmt.Println(errs.Error())
				return
			} else {
				stmt.Exec(i, 0)

			}

			keyArray[i] = 0
		} else {
			keyArray[i] = value
		}

		fmt.Println("key:", i, "value:", keyArray[i])
	}

	fmt.Println("OK")

	syncTimer := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-syncTimer.C:
			for i := 0; i < VK_BOTTOM; i++ {
				if stmt, errs := keydb.Prepare("update keytable set value =? where id=?"); errs != nil {
					fmt.Println(errs.Error())
				} else {
					stmt.Exec(keyArray[i], i)
					fmt.Println("key:", i, "value:", keyArray[i])
				}
			}
			syncTimer.Reset(time.Second * 10)
		}

	}

	//go inputhook.HookKeyboard(hookCallback)
	ch := make(chan bool)
	<-ch
}

func tmp() {
	//打开数据库，如果不存在，则创建
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	//创建表
	sql_table := `
    CREATE TABLE IF NOT EXISTS userinfo(
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR(64) NULL,
        departname VARCHAR(64) NULL,
        created DATE NULL
    );
    `

	db.Exec(sql_table)

	// insert
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("wangshubo", "国务院", "2017-04-21")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("wangshubo_new", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	rows.Close()

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
