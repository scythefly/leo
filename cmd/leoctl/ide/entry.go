package ide

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	aw "github.com/deanishe/awgo"
)

type Storage struct {
	LastKnownMenubarData LastKnownMenubarData `json:"lastKnownMenubarData"`
}
type LastKnownMenubarData struct {
	Menus Meuns `json:"menus"`
}
type Meuns struct {
	File File `json:"File"`
}

type File struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID      string  `json:"id"`
	Submenu Submenu `json:"submenu"`
}

type Submenu struct {
	Items []SubItem `json:"items"`
}

type SubItem struct {
	ID  string `json:"id"`
	Uri Uri    `json:"uri"`
}

type Uri struct {
	Path     string `json:"path"`
	External string `json:"external"`
}

func BuildFeedback(wf *aw.Workflow) {
	// args[1] path
	// args[2] project
	args := wf.Args()
	if len(args) < 3 || args[2] == "" {
		wf.NewItem("code $project").Valid(false)
		return
	}

	file, err := os.Open(args[1])
	if err != nil {
		wf.NewItem(fmt.Sprintf("> failed to open recent projects file %s", args[1])).Valid(false)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	var storage Storage
	if err = decoder.Decode(&storage); err != nil {
		wf.NewItem(fmt.Sprintf("> failed to decode recent projects file %s", args[1])).Valid(false)
		return
	}

	items := storage.LastKnownMenubarData.Menus.File.Items
	var subItems []SubItem

	for idx := range items {
		if items[idx].ID == "submenuitem.MenubarRecentMenu" {
			subItems = items[idx].Submenu.Items
			break
		}
	}

	for idx := range subItems {
		if subItems[idx].ID != "openRecentFolder" {
			continue
		}

		title := subItems[idx].Uri.Path
		fname := title
		if index := strings.LastIndexByte(title, '/'); index > 0 {
			fname = title[index+1:]
		}

		if !strings.Contains(fname, args[2]) {
			continue
		}

		if subItems[idx].Uri.External != "" {
			title = subItems[idx].Uri.External
		}
		wf.NewItem("> open " + fname).Subtitle(title).Arg(title).Valid(true)
	}
}
