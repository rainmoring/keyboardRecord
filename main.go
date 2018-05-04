package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vence722/inputhook"
)

var keydb *sql.DB
var keyArray [VK_BOTTOM]int32

func hookCallback(keyEvent int, keyCode int) {
	keyArray[keyCode] += 1
	fmt.Println("keycode:", keyCode, "value", keyArray[keyCode])
}

func main() {
	var (
		err error
	)

	if keydb, err = sql.Open("sqlite3", "D:/code_dir/GO/src/persion_job/keyboardevent/db/keycode.db"); err != nil {
		fmt.Println("open keycode.db failed")
		fmt.Printf("hello")
		fmt.Print("world")
		return
	}
	keymap := mapinit()
	fmt.Println("bottom:", VK_BOTTOM, "VK_A:", VK_A)

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
					name TEXT(32) NULL,
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
			if stmt, errs := keydb.Prepare("INSERT INTO keytable(id,name,value) values(?,?,?)"); errs != nil {

				fmt.Println(errs.Error())
				return
			} else {
				stmt.Exec(i, keymap[i], 0)

			}

			keyArray[i] = 0
		} else {
			keyArray[i] = value
		}

		fmt.Println("key:", i, "value:", keyArray[i])
	}

	fmt.Println("OK")

	go inputhook.HookKeyboard(hookCallback)

	go func() {
		syncTimer := time.NewTimer(time.Second * 180)
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
				syncTimer.Reset(time.Second * 180)
			}

		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	for {
		<-c
		for i := 0; i < VK_BOTTOM; i++ {
			if stmt, errs := keydb.Prepare("update keytable set value =? where id=?"); errs != nil {
				fmt.Println(errs.Error())
			} else {
				stmt.Exec(keyArray[i], i)
				fmt.Println("key:", i, "value:", keyArray[i])
			}
		}

	}
}
