package models

import "fmt"

type Habit struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Goal        uint   `json:"goal"`
}

// Create a new Habit in the database.
// The name and the description are string, and the goal is an uint.
// This function returns an error.
func CreateHabit(name, description string, goal uint) error {
	statement := "INSERT INTO habits (name, description, goal) VALUES ($1, $2, $3)"
	_, err := db.Exec(statement, name, description, goal)
	return err
}

// Get a new Habit from the database.
// The id used to find the habits is an uint.
// This function returns a pointer to the Habit struct and an error.
func GetHabit(id uint) (*Habit, error) {
	var habit Habit
	statement := "SELECT id, name, description, goal FROM habits WHERE id = $1"
	row := db.QueryRow(statement, id)
	err := row.Scan(&habit.ID, &habit.Name, &habit.Description, &habit.Goal)
	if err != nil {
		return nil, err
	}
	return &habit, nil
}

// Get all of the Habits
// This function returns a slice of Habit and an error
func GetAllHabits() ([]Habit, error) {
	sliceHabit := []Habit{}

	statement := "SELECT id, name, description, goal FROM habits"
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, goal uint
		var name, description string
		err := rows.Scan(&id, &name, &description, &goal)
		if err != nil {
			return nil, err
		}

		currentHabit := Habit{
			ID:          id,
			Name:        name,
			Description: description,
			Goal:        goal,
		}
		sliceHabit = append(sliceHabit, currentHabit)
	}
	return sliceHabit, nil
}

// Update the name, the description and the goal of an Habit in the database
func UpdateHabit(id uint, name, description string, goal uint) error {
	statement := "UPDATE habits SET name = $1, description = $2, goal = $3 WHERE id = $4"
	result, err := db.Exec(statement, name, description, goal, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		fmt.Println("Habit is not found")
	}
	return nil
}

// Delete a Habit from the database
// The id is a uint
func Delete(id uint) error {
	statement := "DELETE FROM habits WHERE id = $1"
	_, err := db.Query(statement, id)
	return err
}
