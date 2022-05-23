package uuid

import "github.com/satori/go.uuid"

func New() string {
	return uuid.NewV4().String()
}
