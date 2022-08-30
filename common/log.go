package common

import (
	"log"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func CheckErrMessage(msg string, err error) {
	if err != nil {
		log.Fatalln(msg, err.Error())
	}
}
