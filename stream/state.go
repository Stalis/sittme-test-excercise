package stream

import "errors"

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
	switch str {
	case "Created":
		return Created, nil
	case "Active":
		return Active, nil
	case "Interrupted":
		return Interrupted, nil
	case "Finished":
		return Finished, nil
	default:
		return Created, errors.New("Invalid State value")
	}
}
