package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	aw "github.com/deanishe/awgo"
	_ "github.com/mattn/go-sqlite3"
	"github.com/scythefly/leo/seven"
)

var (
	helpURL    = "https://www.google.com"
	maxResults = 200
	wf         *aw.Workflow
	db         *sql.DB

	icon = &aw.Icon{
		Value: "com.apple.safari.bookmark",
		Type:  aw.IconTypeFileType,
	}
)

func init() {
	filePath, _ := exec.LookPath(os.Args[0])
	filePath, _ = filepath.Abs(filePath)
	filePath = path.Join(filePath, "../leo.db")
	// log.Infof(">>> %s\n", filePath)
	wf = aw.New(aw.HelpURL(helpURL))

	// filePath := "./leo.db"
	exists, err := fileExists(filePath)
	if err != nil {
		// log.Infof("check file failed, err: %s\n", err.Error())
		os.Remove(filePath)
		exists = false
	}

	if !exists {
		db, err = sql.Open("sqlite3", filePath)
		if err != nil {
			log.Fatalf("cannot open sql, err: %s\n", err.Error())
			return
		}
		sqlString := `create table leo(id integer not null primary key, key text, value text not null, unique(key, value));`
		_, err = db.Exec(sqlString)
		if err != nil {
			log.Fatalf("create table failed, err: %s\n", err.Error())
			return
		}
	} else {
		db, err = sql.Open("sqlite3", filePath)
		if err != nil {
			log.Fatalf("cannot open sql, err: %s\n", err.Error())
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

func run() {
	var query string
	args := wf.Args()
	argsLen := len(args)

	if argsLen > 0 {
		query = args[0]
	}

	if query == "add" || query == "add " {
		it := wf.NewItem("Add ")
		it.Subtitle("insert value to database with key 'Shift'").Valid(true)
		if argsLen > 2 {
			title := fmt.Sprintf("add %s %s", args[1], args[2])
			it.Arg(title)
			it.NewModifier("shift").Subtitle(fmt.Sprintf("add %s -> %s", args[1], args[2])).Valid(true)
		} else if argsLen > 1 {
			title := fmt.Sprintf("add %s", args[1])
			it.Arg(title)
			it.NewModifier("shift").Subtitle(title).Valid(true)
		}
	} else if query == "rm" || query == "rm " {
		if argsLen > 1 {
			seven.BuildRmFeedback(wf, args[1], db)
		} else {
			wf.NewItem("Remove").Subtitle("remove from database with key 'Shift'").Valid(true)
		}
	} else if query != "" {
		// if argsLen > 1 {
		// 	sqlString = fmt.Sprintf("select key,value from leo where key LIKE '%%%s%%' and value LIKE '%%%s%%';", query, args[1])
		// } else {
		// 	sqlString = fmt.Sprintf("select key,value from leo where key LIKE '%%%s%%' or value LIKE '%%%s%%';", query, query)
		// }
		// seven.BuildSQLFeedback(wf, sqlString, db)
		if argsLen > 1 {
			seven.BuildQueryFeedback(wf, query, args[1], db)
		} else {
			seven.BuildQueryFeedback(wf, query, "", db)
		}
	}
	// wf.WarnEmpty("No matching!!!", "NO NO NO NO NO NO")
	if wf.IsEmpty() {
		wf.NewItem("Use add/rm to add or remove a query")
	}
	wf.SendFeedback()
}

func main() {
	defer db.Close()
	wf.Run(run)
}
