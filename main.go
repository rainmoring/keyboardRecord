package main

import (
	"database/sql"
	"fmt"
	"sync/atomic"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vence722/inputhook"
)

var keydb *sql.DB
var keyArray [VK_BOTTOM]uint64

func hookCallback(keyEvent int, keyCode int) {
	atomic.AddUint64(&keyArray[keyCode], 1)
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

                    CREATE TABLE "keytable" (

                    "id"  INTEGER PRIMARY KEY,
                    "value"  INTEGER

                    );`
	//if _, err = keydb.Exec(`CREATE DATABASE IF NOT EXIST keytable(
	if _, err = keydb.Exec(sqlStmt); err != nil {
		fmt.Println(err.Error())
	}

	if err = keydb.Ping(); err != nil {
		fmt.Println("open keycode.db failed")
		return
	}

	fmt.Println("ok")
	for i := 0; i < VK_BOTTOM; i++ {
		var value uint64
		if err = keydb.QueryRow("select value from keytable where id = ?", i).Scan(); err != nil {
			stmt, _ := keydb.Prepare("INSERT INTO keytable(id,count),values(?,?)")
			stmt.Exec(i, 0)
			keyArray[i] = 0
		} else {
			keyArray[i] = value
		}

		fmt.Println("key:", i, "value:", keyArray[i])
	}

	syncTimer := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-syncTimer.C:
			for i := 0; i < VK_BOTTOM; i++ {
				stmt, _ := keydb.Prepare("update table set value =? where id=?")
				stmt.Exec(keyArray[i], i)
				fmt.Println("key:", i, "value:", keyArray[i])
			}
			syncTimer.Reset(time.Second * 10)
		}

	}

	go inputhook.HookKeyboard(hookCallback)
	ch := make(chan bool)
	<-ch
}
