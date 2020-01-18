package main

import (
	"fmt"
	"os"

	"github.com/alexsuslov/aidbox/api"
	. "github.com/logrusorgru/aurora"
)

func helper(help string) {
	switch help {
	case "patch":
		fmt.Println(api.HelpPatchText)
	case "read":
		fmt.Printf("Aidbox %v \n----", version)
		fmt.Printf(helpReadText, Green(exmRead), Green(exmVRead))
	}
	if help != "" {
		os.Exit(0)
	}
}

var helpReadText = `
Reader used to obtain a resource by a given id.

  %v

History used to obtain a resource by a given id and revision.

  %v

`
var exmRead = `aidbox -read /Patient/17b69d79-3d9b-45f8-af79-75f958502763`
var exmVRead = `aidbox -read /Patient/17b69d79-3d9b-45f8-af79-75f958502763/_history/13`
