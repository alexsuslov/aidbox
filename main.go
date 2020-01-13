package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/alexsuslov/aidbox/api"
	"github.com/alexsuslov/godotenv"
	"io"
	"io/ioutil"
	"os"
)

var version string
var ver *bool
var help string
var debugger bool

var config string
var create string
var update string
var read string
var delete string

var ctype string

var contentType = map[string]string{
	"json": "application/json",
	"xml": "text/xml",
	"yml": "text/yaml",
	"yaml": "text/yaml",
}

func init() {
	// debugger
	flag.BoolVar(&debugger, "debugger", false, "Enable debugger")
	// Config
	flag.StringVar(&config, "config", ".env", "Config file")
	// Create
	flag.StringVar(&create, "create", "", "Create a new resource")
	// Read
	flag.StringVar(&read, "read", "", "Read the current state of the resource")
	// update
	flag.StringVar(&update, "update", "",
		"Update an existing resource by its id (or create it if it is new)")
	// Delete
	flag.StringVar(&delete, "delete", "", "Delete a resource")

	// content type
	flag.StringVar(&ctype, "ctype", "json", "Content Type json| xml| yml ")

	ver = flag.Bool("version", false, "Print current version")

	flag.Parse()
}

func main(){
	if err := godotenv.Load(config); err!= nil{
		panic(err)
	}
	if *ver {
      fmt.Println(version)
      os.Exit(0)
    }

	if err := api.Init(); err!= nil {
		panic(fmt.Errorf("Init:%v", err))
	}

	api.DEBUGGING=debugger

	// read
	if read!= ""{
		t, ok := contentType[ctype]
		if !ok{
			panic("Error Content-Type")
		}
		body, err := api.Read(read, &api.ReadOptions{ContentType:t})
		Done(body, err)
		os.Exit(0)
	}
	// Update
	if update!= ""{
		t, ok := contentType[ctype]
		if !ok{
			panic("Error Content-Type")
		}
		reader := bufio.NewReader(os.Stdin)
		body, err := api.Update(read, ioutil.NopCloser(reader), &api.UpdateOptions{ContentType:t})
		Done(body, err)
		os.Exit(0)
	}

	// create
	if create!= ""{
		t, ok := contentType[ctype]
		if !ok{
			panic("Error Content-Type")
		}
		reader := bufio.NewReader(os.Stdin)
		body, err := api.Create(create, ioutil.NopCloser(reader), &api.CreateOptions{
			ContentType:t,
		})
		Done(body, err)
		os.Exit(0)
	}

	// delete
	if delete!= ""{
		t, ok := contentType[ctype]
		if !ok{
			panic("Error Content-Type")
		}
		body, err := api.Delete(create, &api.DeleteOptions{
			ContentType:t,
		})
		Done(body, err)
		os.Exit(0)
	}

}

func Done(body io.ReadCloser, err error){
	if err != nil{
		panic(err)
	}
	defer body.Close()
	if _, err := io.Copy(os.Stdout, body); err!= nil{
		panic(err)
	}
}

