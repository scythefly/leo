package seven

import (
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io"
    "os"
    "regexp"
    "strings"

    aw "github.com/deanishe/awgo"
)

type Option struct {
    Name string `xml:"name,attr"`
    Map  Map    `xml:"map"`
}

type application struct {
    XMLName   xml.Name  `xml:"application"`
    Component Component `xml:"component"`
}

type Component struct {
    Name   string `xml:"name,attr"`
    Option Option `xml:"option"`
}

type Map struct {
    Entry []Entry `xml:"entry"`
}

type Entry struct {
    Key string `xml:"key,attr"`
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
    data, err := io.ReadAll(conf)
    if err != nil {
        wf.NewItem("> read file error")
        return
    }
    app := application{}
    if err = xml.Unmarshal(data, &app); err != nil {
        wf.NewItem("> parse xml file error")
        return
    }

    for _, entry := range app.Component.Option.Map.Entry {
        idx := strings.LastIndex(entry.Key, "/")
        if idx > 0 || project == "-" {
            projectName := entry.Key[idx+1:]
            if strings.Index(projectName, project) >= 0 || project == "-" {
                ppath := reg.ReplaceAllString(entry.Key, "/Users/iuz${1}")
                wf.NewItem("> open " + projectName).Subtitle(ppath).Arg(ppath).Valid(true)
            }
        }
    }
}

type Storage struct {
    LastKnownMenubarData LastKnownMenubarData `json:"lastKnownMenubarData"`
}

type LastKnownMenubarData struct {
    Menus Menus `json:"menus"`
}

type Menus struct {
    File File `json:"file"`
}

type File struct {
    Items []Item `json:"items"`
}

type Item struct {
    ID      string  `json:"id"`
    Label   string  `json:"label"`
    Submenu Submenu `json:"submenu"`
}

type Submenu struct {
    Items []SItem `json:"items"`
}

type SItem struct {
    ID      string `json:"id"`
    Label   string `json:"label"`
    Enabled bool   `json:"enabled"`
    Uri     Uri    `json:"uri"`
}

type Uri struct {
    Path string `json:"path"`
}

func BuildCodeFeedback(wf *aw.Workflow, path, project string) {
    file, err := os.Open(path)
    if err != nil {
        wf.NewItem(fmt.Sprintf("> failed to open recent projects file %s", path))
        return
    }
    defer file.Close()
    decoder := json.NewDecoder(file)
    var storage Storage
    if err = decoder.Decode(&storage); err != nil {
        wf.NewItem("> decode json file error")
        return
    }
    items := storage.LastKnownMenubarData.Menus.File.Items
    for idx := range items {
        if items[idx].Label == "Open &&Recent" {
            subs := items[idx].Submenu.Items
            for i := range subs {
                if subs[i].Enabled && subs[i].ID == "openRecentFolder" {
                    if project != "-" && strings.Index(subs[i].Label, project) < 0 {
                        continue
                    }
                    wf.NewItem("> open " + subs[i].Label).Subtitle(subs[i].Uri.Path).Arg(subs[i].Uri.Path).Valid(true)
                }
            }
        }
    }
}
