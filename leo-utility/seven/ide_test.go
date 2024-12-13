package seven_test

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"

	"leo-utility/seven"
)

type application struct {
	XMLName   xml.Name        `xml:"application"`
	Component seven.Component `xml:"component"`
}

func Test_IDEFeedback(t *testing.T) {
	path := "/Users/iuz/Library/ApplicationSupport/JetBrains/GoLand2020.3/options/recentProjects.xml"
	project := "u"
	reg := regexp.MustCompile(`^\$USER_HOME\$(.*)$`)
	userHome := os.Getenv("USER_HOME")
	if userHome == "" {
		userHome = "/Users/iuz"
	}
	conf, err := os.Open(path)
	if err != nil {
		t.Log(fmt.Sprintf("> failed to open recent projects file %s", path))
		return
	}
	data, err := ioutil.ReadAll(conf)
	if err != nil {
		t.Log("> read file error")
		return
	}
	app := application{}
	if err = xml.Unmarshal(data, &app); err != nil {
		t.Log("> parse xml file error")
		return
	}

	for _, entry := range app.Component.Option.Map.Entry {
		idx := strings.LastIndex(entry.Key, "/")
		if idx > 0 || project == "-" {
			projectName := entry.Key[idx+1:]
			if strings.Index(projectName, project) >= 0 || project == "-" {
				ppath := reg.ReplaceAllString(entry.Key, "/Users/iuz${1}")
				t.Log("> open " + projectName)
				t.Log(ppath)
			}
		}
	}
}
