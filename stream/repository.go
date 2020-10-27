package stream

import (
	"errors"

	"github.com/google/uuid"
)

// Repository репозиторий трансляций
type Repository struct {
	streams map[uuid.UUID]*Stream
}

// NewRepository возвращает новый экземпляр репозитория трансляций
func NewRepository() *Repository {
	return &Repository{
		make(map[uuid.UUID]*Stream),
	}
}

// CreateStream создает трансляцию и возвращает ее id
func (r *Repository) CreateStream() (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, nil
	}
	r.streams[id] = newStream()
	return id, nil
}

// GetInfo возвращает копию данных трансляции
func (r *Repository) GetInfo(id uuid.UUID) (Stream, error) {
	stream, ok := r.streams[id]
	if !ok {
		return Stream{}, errors.New("Not found stream")
	}
	return *stream, nil
}

// RemoveStream удаляет трансляцию
func (r *Repository) RemoveStream(id uuid.UUID) {
	delete(r.streams, id)
}
