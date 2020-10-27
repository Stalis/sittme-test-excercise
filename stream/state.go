package stream

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
