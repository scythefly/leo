package seven

import (
	"database/sql"
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/gogf/gf/g/container/gset"
	"github.com/gogf/gf/g/os/glog"
)

// BuildFeedback ...
func BuildFeedback(wf *aw.Workflow, query string, db *sql.DB) {
	sqlString := fmt.Sprintf("select key,value from leo where key LIKE '%%%s%%' or value LIKE '%%%s%%';", query, query)
	var key, value string
	rows, err := db.Query(sqlString)
	if err != nil {
		glog.Infof("query failed, err: %s\n", err.Error())
		return
	}
	i := 0
	defer rows.Close()

	for rows.Next() && i < 10 {
		if err = rows.Scan(&key, &value); err != nil {
			glog.Infof("read data failed, err: %s\n", err.Error())
			return
		}
		wf.NewItem("> " + value).Subtitle(fmt.Sprintf("%s -> %s", key, value)).Copytext(value).Arg(value).Valid(true)
		i++
	}
}

func buildFuzzyQuery(wf *aw.Workflow, arg1, arg2 string, db *sql.DB, s *gset.Set) {
	var key, value, sqlString string
	var idx int

	if arg2 == "" {
		sqlString = fmt.Sprintf("select id,key,value from leo where key LIKE '%%%s%%' or value LIKE '%%%s%%';", arg1, arg1)
	} else {
		sqlString = fmt.Sprintf("select id,key,value from leo where key LIKE '%%%s%%' and value LIKE '%%%s%%';", arg1, arg2)
	}

	rows, err := db.Query(sqlString)
	if err != nil {
		wf.NewItem(fmt.Sprintf("> Error: %s", err.Error())).Valid(true)
		return
	}
	i := 0
	defer rows.Close()

	for rows.Next() && i < 5 {
		if err = rows.Scan(&idx, &key, &value); err != nil {
			wf.NewItem(fmt.Sprintf("> Error: %s", err.Error())).Valid(true)
			return
		}
		if s.Contains(idx) {
			continue
		}
		wf.NewItem("> " + value).Subtitle(fmt.Sprintf("%s -> %s", key, value)).Copytext(value).Arg(value).Valid(true)
		i++
	}
}

// BuildQueryFeedback ...
func BuildQueryFeedback(wf *aw.Workflow, arg1, arg2 string, db *sql.DB) {
	var key, value, sqlString string
	var idx int

	if arg2 == "" {
		sqlString = fmt.Sprintf("select id,key,value from leo where key LIKE '%s%%' or value LIKE '%s%%';", arg1, arg1)
	} else {
		sqlString = fmt.Sprintf("select id,key,value from leo where key LIKE '%s%%' and value LIKE '%s%%';", arg1, arg2)
	}

	rows, err := db.Query(sqlString)
	if err != nil {
		wf.NewItem(fmt.Sprintf("> Error: %s", err.Error())).Valid(true)
		return
	}
	i := 0
	defer rows.Close()

	s := gset.New(true)

	for rows.Next() && i < 10 {
		if err = rows.Scan(&idx, &key, &value); err != nil {
			wf.NewItem(fmt.Sprintf("> Error: %s", err.Error())).Valid(true)
			return
		}
		wf.NewItem("> " + value).Subtitle(fmt.Sprintf("%s -> %s", key, value)).Copytext(value).Arg(value).Valid(true)
		s.Add(idx)
		i++
	}

	if i > 8 {
		return
	}
	buildFuzzyQuery(wf, arg1, arg2, db, s)
}

// BuildSQLFeedback ...
func BuildSQLFeedback(wf *aw.Workflow, sqlString string, db *sql.DB) {
	var key, value string
	rows, err := db.Query(sqlString)
	if err != nil {
		wf.NewItem(fmt.Sprintf("> Error: %s", err.Error())).Valid(true)
		return
	}

	i := 0
	defer rows.Close()

	for rows.Next() && i < 10 {
		if err = rows.Scan(&key, &value); err != nil {
			wf.NewItem(fmt.Sprintf("> Error: %s", err.Error())).Valid(true)
			return
		}
		wf.NewItem("> " + value).Subtitle(fmt.Sprintf("%s -> %s", key, value)).Copytext(value).Arg(value).Valid(true)
		i++
	}
}

// BuildRmFeedback ...
func BuildRmFeedback(wf *aw.Workflow, query string, db *sql.DB) {
	sqlString := fmt.Sprintf("select id,key,value from leo where key LIKE '%%%s%%' or value LIKE '%%%s%%';", query, query)
	var key, value string
	var id int
	rows, err := db.Query(sqlString)
	if err != nil {
		wf.NewItem("Error").Subtitle(err.Error()).Valid(true)
		return
	}
	i := 0
	defer rows.Close()

	for rows.Next() && i < 10 {
		if err = rows.Scan(&id, &key, &value); err != nil {
			glog.Infof("read data failed, err: %s\n", err.Error())
			wf.NewItem("Error").Subtitle(err.Error()).Valid(true)
			return
		}
		it := wf.NewItem("Remove")
		if key == "" {
			title := fmt.Sprintf("rm " + value)
			it.Arg(fmt.Sprintf("rm %d", id)).Subtitle(fmt.Sprintf("remove '%s' with key 'Shift'", value)).Valid(true)
			it.NewModifier("shift").Subtitle(title)
		} else {
			title := fmt.Sprintf("rm %s %s", key, value)
			it.Arg(fmt.Sprintf("rm %d", id)).Subtitle(fmt.Sprintf("remote '%s -> %s' with key 'Shift'", key, value)).Valid(true)
			it.NewModifier("shift").Subtitle(title)
		}
		i++
	}
	if i == 0 {
		wf.NewItem("Remove").Subtitle("remove from database with key 'Shift'").Valid(true)
	}
}
