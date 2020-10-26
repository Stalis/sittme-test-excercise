package translation

import (
	"errors"

	"github.com/google/uuid"
)

type TranslationRepository struct {
	translations map[uuid.UUID]Translation
}


func New() *TranslationRepository {
	return &TranslationRepository{
		make(map[uuid.UUID]Translation),
	}
}

func (t *TranslationRepository) CreateTranslation() (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, nil
	}
	t.translations[id] = newTranslation(id)
}

func (t *TranslationRepository) GetState(id uuid.UUID) (State, error) {
	stream, ok := t.translations[id]
	if !ok {
		return 0, errors.New("Not found stream")
	}
	return stream.State(), nil
}

// RemoveStream удаляет трансляцию
func (t *TranslationRepository) RemoveStream(id uuid.UUID) {
	delete(t.translations, id)
}

