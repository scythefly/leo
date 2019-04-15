package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/gogf/gf/g/os/glog"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func init() {
	filePath, _ := exec.LookPath(os.Args[0])
	filePath, _ = filepath.Abs(filePath)
	filePath = path.Join(filePath, "../leo.db")
	// glog.Infof(">>> %s\n", filePath)

	// filePath := "./leo.db"
	exists, err := fileExists(filePath)
	if err != nil {
		// glog.Infof("check file failed, err: %s\n", err.Error())
		os.Remove(filePath)
		exists = false
	}

	if !exists {
		db, err = sql.Open("sqlite3", filePath)
		if err != nil {
			glog.Fatalf("cannot open sql, err: %s\n", err.Error())
			return
		}
		sqlString := `
		create table leo(id integer not null primary key, key text, value text not null, unique(key, value));`
		_, err = db.Exec(sqlString)
		if err != nil {
			glog.Fatalf("create table failed, err: %s\n", err.Error())
			return
		}
	} else {
		db, err = sql.Open("sqlite3", filePath)
		if err != nil {
			glog.Fatalf("cannot open sql, err: %s\n", err.Error())
			return
		}
	}
}

func fileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}

func main() {
	// glog.SetPath("/Users/iuz/temp")
	var sqlString, method string
	argsLen := len(os.Args)
	if argsLen > 3 {
		glog.Infof("%s %s %s\n", os.Args[1], os.Args[2], os.Args[3])
		method = os.Args[1]
		key := os.Args[2]
		value := os.Args[3]

		if method == "add" {
			sqlString = fmt.Sprintf("insert into leo(key,value) values('%s','%s');", key, value)
			db.Exec(sqlString)
		}
	} else if argsLen > 2 {
		glog.Infof("%s %s\n", os.Args[1], os.Args[2])
		method = os.Args[1]
		value := os.Args[2]
		if method == "add" {
			sqlString = fmt.Sprintf("insert into leo(key,value) values('','%s');", value)
			db.Exec(sqlString)
		} else if method == "rm" {
			sqlString = fmt.Sprintf("delete from leo where id=%s;", value)
			db.Exec(sqlString)
		}
	}
}
