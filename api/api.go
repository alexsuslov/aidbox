package api

import (
	"github.com/alexsuslov/godotenv"
	"log"
)

var(
	DEBUGGING bool
	_InsecureSkipVerify bool
	_host string
	_client string
	_secret string
)

func Init()error{
	_host = godotenv.GetPanic("AIDBOX_HOST")
	_client = godotenv.GetPanic("AIDBOX_CLIENT")
	_secret = godotenv.GetPanic("AIDBOX_SECRET")
	_InsecureSkipVerify = godotenv.GetPanic("AIDBOX_INSECURE")=="YES"
	return nil
}

func Print(opts ...interface{}) {
	if DEBUGGING {
		log.Println(opts...)
	}
}
