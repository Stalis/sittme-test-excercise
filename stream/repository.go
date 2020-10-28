package stream

import "github.com/google/uuid"

// Repository хранилище данных о трансляциях
type Repository interface {
	CreateStream() (uuid.UUID, error)
	SetState(uuid.UUID, State) error
	GetInfo(uuid.UUID) (Stream, error)
	RemoveStream(uuid.UUID) error
}
