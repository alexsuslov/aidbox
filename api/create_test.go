package api_test

import (
	"github.com/alexsuslov/aidbox/api"
	"github.com/alexsuslov/godotenv"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	if err := godotenv.Load("../.env"); err!= nil{
		panic(err)
	}
	if err := api.Init(); err!= nil{
		panic(err)
	}
	f, err := os.Open("create_test.txt")
	if err!= nil{
		panic(err)
	}
	type args struct {
		resource string
		Req      io.ReadCloser
		options  *api.CreateOptions
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
	}{
		{
			"create patient",
			args{
				"Patient",
				f,
				&api.CreateOptions{
					ContentType: "text/yaml",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api.DEBUGGING=true
			gotBody, err := api.Create(tt.args.resource, tt.args.Req, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			data, err := ioutil.ReadAll(gotBody)
			if err!= nil{
				panic(err)
			}
			log.Println(string(data))
		})
	}
}
