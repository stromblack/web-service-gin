package common

import (
	"log"
	"os"
)

func CheckErr(err error) {
	// check if on labmda log only
	env := os.Getenv("GIN_MODE")
	if env == "release" {
		panicOnError(err)
	} else {
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckErrMessage(msg string, err error) {
	if err != nil {
		log.Fatalln(msg, err.Error())
	}
}
