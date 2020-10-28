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

// SetState пробует изменить состояние трансляции на указанное
func (r *MapRepository) SetState(id uuid.UUID, state State) error {
	str, err := r.getStream(id)
	if err != nil {
		return err
	}

	err = str.SetState(state)
	if err != nil {
		return err
	}

	return nil
}

// GetInfo возвращает копию данных трансляции
func (r *MapRepository) GetInfo(id uuid.UUID) (Stream, error) {
	stream, err := r.getStream(id)
	if err != nil {
		return Stream{}, err
	}

	return *stream, nil
}

// RemoveStream удаляет трансляцию
func (r *MapRepository) RemoveStream(id uuid.UUID) error {
	_, err := r.getStream(id)
	delete(r.streams, id)
	return err
}

func (r *MapRepository) getStream(id uuid.UUID) (*Stream, error) {
	str, ok := r.streams[id]
	if !ok {
		return nil, errors.New("Stream with id [" + id.String() + "] not found")
	}
	return str, nil
}
