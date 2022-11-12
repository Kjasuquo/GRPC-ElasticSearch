package main

import (
	myService "gitlab.com/dh-backend/search-service/internal/services"
	"log"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	myService.Start()
}
