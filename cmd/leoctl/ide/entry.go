package ide

import (
	"encoding/json"
	"fmt"
	"leo/internal/pb"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ide",
		Run: ideRun,
	}
	return cmd
}

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
	Path      string `json:"path"`
	External  string `json:"external"`
	Scheme    string `json:"scheme"`
	Authority string `json:"authority"`
}

func ideRun(_ *cobra.Command, args []string) {
	// args[0] path
	// args[1] project
	if len(args) < 2 || args[1] == "" {
		pb.Wf.NewItem("code $project").Valid(false)
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		pb.Wf.NewItem(fmt.Sprintf("> failed to open recent projects file %s", args[0])).Valid(false)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	var storage Storage
	if err = decoder.Decode(&storage); err != nil {
		pb.Wf.NewItem(fmt.Sprintf("> failed to decode recent projects file %s", args[0])).Valid(false)
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

		if !strings.Contains(fname, args[1]) {
			continue
		}

		// "vscode-remote://ssh-remote%2Bubuntu01/root/work/friday/LibreChat"
		uri := &subItems[idx].Uri
		if uri.External != "" {
			title = uri.External
		} else if uri.Scheme == "vscode-remote" {
			title = fmt.Sprintf("%s://%s%s",
				uri.Scheme,
				url.QueryEscape(uri.Authority),
				uri.Path,
			)
		}
		pb.Wf.NewItem("> open " + fname).Subtitle(title).Arg(title).Valid(true)
	}
}
