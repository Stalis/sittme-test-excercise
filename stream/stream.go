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

// SetState пробует изменить состояние трансляции на указанное
func (s *Stream) SetState(state State) error {
	switch state {
	case Created:
		return errors.New("Cannot set stream state as Created")
	case Active:
		return s.Activate()
	case Interrupted:
		return s.Interrupt()
	case Finished:
		return s.Finish()
	default:
		return errors.New("Unknown state")
	}
}

// Activate активирует трансляцию после создания или прерывания
func (s *Stream) Activate() error {
	if s.State == Active {
		return errors.New("Stream is already Active")
	}
	if s.State != Created && s.State != Interrupted {
		return errors.New("Invalid stream State")
	}

	s.State = Active
	return nil
}

// Interrupt прерывает трансляцию
func (s *Stream) Interrupt() error {
	if s.State == Interrupted {
		return errors.New("Stream is already interrupted")
	}
	if s.State != Active {
		return errors.New("Stream is not Active")
	}

	s.State = Interrupted
	return nil
}

// Finish завершает трансляцию
func (s *Stream) Finish() error {
	if s.State == Finished {
		return errors.New("Stream is already Finished")
	}

	s.State = Finished
	return nil
}
