package models

type HabitEntry struct {
	ID      int    `json:"id"`
	HabitID int    `json:"habit_id"`
	Date    string `json:"date"`
	Status  bool   `json:"status"`
	Habit   Habit  `json:"habit"`
}
