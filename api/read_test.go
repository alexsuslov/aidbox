package api_test

import (
	"github.com/alexsuslov/godotenv"
	"github.com/alexsuslov/aidbox/api"
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
			"read /entities/Entity/markdown",
			args{
				resource: "/entities/Entity/markdown",
				options: &api.ReadOptions{
					ContentType: "application/json",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBody, err := api.Read(tt.args.resource, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
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