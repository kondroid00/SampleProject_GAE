package model

import (
	"github.com/satori/go.uuid"
)

func getUUID() uuid.UUID {
	return uuid.NewV4()
}
