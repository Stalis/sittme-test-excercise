package stream

import (
	"errors"
	"time"
)

// Stream описывает идентификацию через
type Stream struct {
	State   State     `json:"state"`
	Created time.Time `json:"created"`
}

func newStream() *Stream {
	return &Stream{
		State:   Created,
		Created: time.Now(),
	}
}

// // State возвращает состояние трансляции
// func (t Stream) State() State {
// 	return t.State
// }

// // Created возвращает время создания трансляция
// func (t Stream) Created() time.Time {
// 	return t.Created
// }

// Activate активирует трансляцию после создания или прерывания
func (t *Stream) Activate() error {
	if t.State != Created && t.State != Interrupted {
		return errors.New("Invalid stream State")
	}
	t.State = Active
	return nil
}

// Interrupt прерывает трансляцию
func (t *Stream) Interrupt() error {
	if t.State != Active {
		return errors.New("Stream is not active")
	}
	t.State = Interrupted
	return nil
}

// Finish завершает трансляцию
func (t *Stream) Finish() error {
	t.State = Finished
	return nil
}
