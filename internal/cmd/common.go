package cmd

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

func continueOrFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	ErrReceivedInterrupt = errors.New("received an interrupt")
)
