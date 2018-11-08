package hippo

import (
	"log"
)

func CheckError(err error) {
    if err != nil {
	log.Fatalf("Error: %s", err)
    }
}
