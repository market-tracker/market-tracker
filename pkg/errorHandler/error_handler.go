package errorHandler

import "log"

func PanicError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}
