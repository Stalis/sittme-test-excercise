package translation

import (
	"errors"

	"github.com/google/uuid"
)

// Translation описывает идентификацию через
type Translation struct {
	state State
}

func newTranslation(id uuid.UUID) *Translation {
	return &Translation{
		Created,
	}
}

// State возвращает состояние трансляции
func (t Translation) State() State {
	return t.state
}

// Activate активирует трансляцию после создания или прерывания
func (t *Translation) Activate() error {
	if t.state != Created && t.state != Interrupted {
		return errors.New("Invalid translation state")
	}
	t.state = Active
	return nil
}

// Interrupt прерывает трансляцию
func (t *Translation) Interrupt() error {
	if t.state != Active {
		return errors.New("Translation is not active")
	}
	t.state = Interrupted
	return nil
}

// Finish завершает трансляцию
func (t *Translation) Finish() error {
	t.state = Finished
	return nil
}
