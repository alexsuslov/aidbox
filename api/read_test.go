package api_test

import (
	"github.com/alexsuslov/aidbox/api"
	"github.com/alexsuslov/godotenv"
	"io/ioutil"
	"log"
	"testing"
)

func TestRead(t *testing.T) {
	if err := godotenv.Load("../.env"); err!= nil{
		panic(err)
	}
	if err := api.Init(); err!= nil{
		panic(err)
	}
	type args struct {
		resource string
		options  *api.ReadOptions
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
	}{
		{
			"read Patient/0b72a6dd-5f83-4bb9-9aac-27a7786e2536",
			args{
				resource: "Patient/0b72a6dd-5f83-4bb9-9aac-27a7786e2536",
				options: &api.ReadOptions{
					ContentType: "application/json",
				},
			},
			false,
		},
		{
			"read Patient/0b72a6dd-5f83-4bb9-9aac-27a7786e2536_xxx",
			args{
				resource: "Patient/0b72a6dd-5f83-4bb9-9aac-27a7786e2536_xxx",
				options: &api.ReadOptions{
					ContentType: "application/json",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api.DEBUGGING=true
			gotBody, err := api.Read(tt.args.resource, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotBody!= nil{
				data, err := ioutil.ReadAll(gotBody)
				if err!= nil{
					panic(err)
				}
				log.Println(string(data))
			}

		})
	}
}
