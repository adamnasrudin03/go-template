package helpers

import "log"

func PanicRecover(opName string) {
	if r := recover(); r != nil {
		log.Printf("%v panic recover: %v \n", opName, r)
	}
}
