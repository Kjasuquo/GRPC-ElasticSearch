package services

import "gitlab.com/dh-backend/search-service/internal/rabbitMQ"

func InsertMenu() {
	rabbitMQ.ReadFromQueueToInsertInES("menu", "menuCreationChannel")
}
