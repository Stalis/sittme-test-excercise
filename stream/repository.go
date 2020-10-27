package stream

import "github.com/google/uuid"

type Repository interface {
	CreateStream() (uuid.UUID, error)
	GetInfo(uuid.UUID) (Stream, error)
	RemoveStream(uuid.UUID) error
}
