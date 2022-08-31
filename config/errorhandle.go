package config

import (
	"log"
	"os"
)

func Errorhandle(err error, msm string, exit int) {
	if err != nil {
		log.Fatal(msm+"\n", err.Error())
		os.Exit(exit)
	}
}
