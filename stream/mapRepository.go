package stream

import (
	"errors"

	"github.com/google/uuid"
)

// MapRepository репозиторий трансляций
type MapRepository struct {
	streams map[uuid.UUID]*Stream
}

// NewMapRepository возвращает новый экземпляр репозитория трансляций
func NewMapRepository() *MapRepository {
	return &MapRepository{
		make(map[uuid.UUID]*Stream),
	}
}

// CreateStream создает трансляцию и возвращает ее id
func (r *MapRepository) CreateStream() (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, nil
	}
	r.streams[id] = newStream()
	return id, nil
}

// GetInfo возвращает копию данных трансляции
func (r *MapRepository) GetInfo(id uuid.UUID) (Stream, error) {
	stream, ok := r.streams[id]
	if !ok {
		return Stream{}, errors.New("Not found stream")
	}
	return *stream, nil
}

// RemoveStream удаляет трансляцию
func (r *MapRepository) RemoveStream(id uuid.UUID) error {
	delete(r.streams, id)
	return nil
}
