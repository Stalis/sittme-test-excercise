package stream

import (
	"errors"
	"strconv"
	"strings"
)

// State Тип состояния трансляции
type State int

const (
	// Created Трансляция создана
	Created State = iota
	// Active Трансляция активна
	Active
	// Interrupted Трансляция прервана
	Interrupted
	// Finished Трансляция завершена
	Finished
)

func (s State) String() string {
	switch s {
	case Created:
		return "Created"
	case Active:
		return "Active"
	case Interrupted:
		return "Interrupted"
	case Finished:
		return "Finished"
	default:
		return "Invalid"
	}
}

// ParseState возвращает значение состояния из строки
func ParseState(str string) (State, error) {
	res, err := strconv.Atoi(str)
	if err == nil {
		s := State(res)
		if s < Created || s > Finished {
			return s, errors.New("Invalid State value")
		}
		return s, nil
	}

	switch strings.ToLower(str) {
	case "created":
		return Created, nil
	case "active":
		return Active, nil
	case "interrupted":
		return Interrupted, nil
	case "finished":
		return Finished, nil
	}

	return State(-1), errors.New("Invalid State value")
}
