package helper

import (
	"log"
	"os"
)

func CheckMarshalError(err error) {
	if err != nil {
		log.Fatalf("cannot marshal response: %v", err)
	}
}

func CheckCommonError(err error, reason string) {
	if err != nil {
		log.Fatalf("%s: %v", reason, err)
	}
}

func CheckMissingFlag(flags string) {
	log.Printf("Error: missing required flag(s) %s\n", flags)
	log.Println("Run with --help to see usage.")
	os.Exit(1)
}
