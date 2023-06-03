package checkerror

import "log"

func Checkerror(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
