package functions

import "log"

func ErrManagement(err error) {
	if err != nil {
		log.Fatal("!! ERROR !!:", err)
	}
}
