package utils

import "log"

func CheckErrors(err ...error) {
	for _, v := range err {
		if v != nil {
			log.Fatal(err)
		}
	}
}
