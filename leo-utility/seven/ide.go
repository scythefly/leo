package seven

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	aw "github.com/deanishe/awgo"
)

type Option struct {
	Name string `xml:"name,attr"`
	List List   `xml:"list"`
}

type application struct {
	XMLName   xml.Name  `xml:"application"`
	Component Component `xml:"component"`
}

type Component struct {
	Name   string   `xml:"name,attr"`
	Option []Option `xml:"option"`
}

type List struct {
	Value []Value `xml:"option"`
}

type Value struct {
	Value string `xml:"value,attr"`
}

func BuildIDEFeedback(wf *aw.Workflow, path, project string) {
	reg := regexp.MustCompile(`^\$USER_HOME\$(.*)$`)
	userHome := os.Getenv("USER_HOME")
	if userHome == "" {
		userHome = "/Users/iuz"
	}
	conf, err := os.Open(path)
	if err != nil {
		wf.NewItem(fmt.Sprintf("> failed to open recent projects file %s", path))
		return
	}
	data, err := ioutil.ReadAll(conf)
	if err != nil {
		wf.NewItem("> read file error")
		return
	}
	app := application{}
	if err = xml.Unmarshal(data, &app); err != nil {
		wf.NewItem("> parse xml file error")
		return
	}

	for _, opt := range app.Component.Option {
		if opt.Name == "recentPaths" {
			for _, l := range opt.List.Value {
				idx := strings.LastIndex(l.Value, "/")
				if idx > 0 || project == "-" {
					projectName := l.Value[idx+1:]
					if strings.Index(projectName, project) >= 0 || project == "-" {
						ppath := reg.ReplaceAllString(l.Value, "/Users/iuz${1}")
						wf.NewItem("> open " + projectName).Arg(ppath).Valid(true)
					}
				}
			}
		}
	}
}
