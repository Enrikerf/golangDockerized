package Core

import "github.com/google/uuid"

type Id interface {
	GetUUID() uuid.UUID
	GetStringValue() string
}

type id struct {
	uuid uuid.UUID
}

func NewId() Id {
	return &id{uuid.New()}
}

func (id *id) GetUUID() uuid.UUID {
	return id.uuid
}

func (id *id) GetStringValue() string {
	return id.uuid.String()
}
