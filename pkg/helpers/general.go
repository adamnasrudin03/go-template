package helpers

import (
	"log"
	"time"
)

func TimeZoneJakarta() *time.Location {
	jakarta, err := time.LoadLocation(AsiaJakarta)
	if err != nil {
		log.Println("should not error on generating current time in jakarta")
	}

	return jakarta
}
