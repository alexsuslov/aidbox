package main

import (
	"fmt"
	"os"
)

func helper(help string){
	switch help{

	case "read":
		fmt.Printf("Aidbox %v \n----", version)
		fmt.Println(helpReadText)
	}
	if help!=""{
		os.Exit(0)
	}
}



var helpReadText = `
Reader used to obtain a resource by a given id.

	aidbox -read /Patient/17b69d79-3d9b-45f8-af79-75f958502763

History used to obtain a resource by a given id and revision.

	aidbox -read /Patient/17b69d79-3d9b-45f8-af79-75f958502763/_history/13

`

