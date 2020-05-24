package main

import (
	"fmt"
	aw "github.com/deanishe/awgo"
	"path"
)

var (
	modKeys = []aw.ModKey {
		aw.ModCmd,
		aw.ModCtrl,
		aw.ModFn,
		aw.ModOpt,
		aw.ModShift,
	}
)

func getIcon(repo string) *aw.Icon {
	return &aw.Icon{Value: path.Join(fmt.Sprintf("resources/%s", getIconName(repo)))}
}

func filter(query string) {
	wf.Filter(query)
}

func addItem(repo string) {
	item := wf.NewItem(getRepoName(repo, true)).
		Arg(getRepoURL(repo)).
		Subtitle(getDomain(repo)).
		Icon(getIcon(repo)).
		Valid(true)

	for _, key := range modKeys {
		item.SetModifier(createModifier(repo, key))
	}
}

func createModifier(repo string, key aw.ModKey) *aw.Modifier {
	mod := &aw.Modifier{Key: key}
	return mod.
		Arg(getArgWithModifier(repo, key)).
		Subtitle(getSubWithModifier(key)).
		Valid(true)
}

func getArgWithModifier(repo string, key aw.ModKey) string {
	switch key {
	case aw.ModCmd:
		return repo
	case aw.ModCtrl:
		return repo
	case aw.ModFn:
		return getRepoName(repo, true)
	case aw.ModOpt:
		return repo
	case aw.ModShift:
		return getRepoURL(repo)
	}
	return getRepoURL(repo)
}

func getSubWithModifier(key aw.ModKey) string {
	switch key {
	case aw.ModCmd:
		return "Reveal in Finder."
	case aw.ModCtrl:
		return "Browse in terminal."
	case aw.ModFn:
		return "Search in browser."
	case aw.ModOpt:
		return "Open files in VSCode."
	case aw.ModShift:
		return "Open URL."
	}
	return "Open URL."
}
