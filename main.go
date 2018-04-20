package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vence722/inputhook"
)

var keyS [VK_BOTTOM]int
var db *sql.DB

type KeyRecord struct {
	Id    sql.NullInt64
	value sql.NullInt64
}

func hookCallback(keyEvent int, keyCode int) {
	keyS[keyCode] += 1
}

func main() {
	var (
		err error
	)
	if db, err = sql.Open("sqlite3", "keycode.db"); err != nil {
		log.Fatal("open keycode.db failed")
	}

	if table, err = db.Exec("CREATE DATABASE IF NOT EXIST keycount( value int primary key not null ,count int)"); err != nil {
	}
	if err = db.Ping(); err != nil {
		log.Fatal("connect sqlite3 failed")
	}

	rows, err := db.Query("select id, name from user")
	inputhook.HookKeyboard(hookCallback)
	ch := make(chan bool)
	<-ch

}
