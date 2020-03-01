package errorHandler

import (
	"fmt"
	"log"
)

func Handler(msg string, err error) {
	if err != nil {
		log.Printf("%T\n", err)
		fmt.Printf("%s\n", msg)
		log.Fatalln(err)
	}
}
