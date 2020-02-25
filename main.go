package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/alexsuslov/aidbox/api"
	"github.com/alexsuslov/godotenv"
)

var version string
var ver *bool
var entity *bool
var help string
var debugger bool

var config string
var create string
var update string
var patch string

var read string
var delete string

var ctype string

var contentType = map[string]string{
	"json": "application/json",
	"xml":  "text/xml",
	"yml":  "text/yaml",
	"yaml": "text/yaml",
}

func init() {
	// debugger
	flag.BoolVar(&debugger, "debugger", false, "Enable debugger")
	// Config
	flag.StringVar(&config, "config", ".env", "Config file")
	// Help
	flag.StringVar(&help, "help", "", "Help")
	// Create
	flag.StringVar(&create, "create", "", "Create a new resource")
	// Read
	flag.StringVar(&read, "read", "", "Read the current state of the resource")
	// update
	flag.StringVar(&update, "update", "",
		"Update an existing resource by its id (or create it if it is new)")
	// Patch
	flag.StringVar(&patch, "patch", "", "Patch part of your resource")
	// Delete
	flag.StringVar(&delete, "delete", "", "Delete a resource")

	// content type
	flag.StringVar(&ctype, "ctype", "json", "Content Type json| xml| yml ")

	entity = flag.Bool("entity", false, "get list primitive")

	ver = flag.Bool("version", false, "Print current version")

	flag.Parse()
}

func main() {
	if err := godotenv.Load(config); err != nil {
		fmt.Printf("Warrning no .env file. %v", err)
	}
	if *ver {
		fmt.Println(version)
		os.Exit(0)
	}
	client, err := api.New(
		godotenv.GetPanic("AIDBOX_HOST"),
		godotenv.GetPanic("AIDBOX_CLIENT"),
		godotenv.GetPanic("AIDBOX_SECRET"),
		godotenv.GetPanic("AIDBOX_INSECURE") == "YES",
	)
	if err != nil {
		panic(fmt.Errorf("Init:%v", err))
	}

	api.DEBUGGING = debugger
	helper(help)
	// Read
	client.MainRead(ctype, read)
	// Update
	client.MainUpdate(ctype, update)
	// Patch
	client.MainPatch(ctype, patch)

	// create
	if create != "" {
		t, ok := contentType[ctype]
		if !ok {
			panic("Error Content-Type")
		}
		reader := bufio.NewReader(os.Stdin)
		body, err := client.Create(create, ioutil.NopCloser(reader), &api.CreateOptions{
			ContentType: t,
		})
		Done(body, err)
		os.Exit(0)
	}

	// delete
	if delete != "" {
		t, ok := contentType[ctype]
		if !ok {
			panic("Error Content-Type")
		}
		body, err := client.Delete(create, &api.DeleteOptions{
			ContentType: t,
		})
		Done(body, err)
		os.Exit(0)
	}

	client.MainEntity(*entity, ctype, ioutil.NopCloser(bufio.NewReader(os.Stdin)))

}

// Done Done
func Done(body io.ReadCloser, err error) {
	if err != nil {
		panic(err)
	}
	defer body.Close()
	if _, err := io.Copy(os.Stdout, body); err != nil {
		panic(err)
	}
}
