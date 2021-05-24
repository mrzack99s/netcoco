package system

import "github.com/google/uuid"

func NewUUID() string {
	newUUID := uuid.New()
	return newUUID.String()
}
