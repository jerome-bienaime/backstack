package config

import (
	"log"
	"strconv"
	"time"
)

func RequestTimeout(duration string) time.Duration {
	timeout, err := strconv.Atoi(duration)
	if err != nil {
		log.Fatal(err)
	}

	return time.Duration(timeout) * time.Millisecond
}
