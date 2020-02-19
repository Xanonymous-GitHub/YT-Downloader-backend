package errorHandler

import (
	"fmt"
	"log"
)

func Handler(msg string, err error) {
	if err != nil {
		fmt.Printf("%s\n", msg)
		log.Fatalln(err)
	}
}
