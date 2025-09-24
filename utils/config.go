package utils

import (
	"log"
)

type SimpleLogger interface {
	Infow(message string)
}

func logger(message string) {
	log.Println(message)
}
