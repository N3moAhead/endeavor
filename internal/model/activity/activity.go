package activity

import (
	"fmt"

	"github.com/N3moAhead/endeavor/internal/db"
)

type Activity struct {
	ID         int
	Name       string
	CategoryID int
}

func GetAll() ([]Activity, error) {
	query := `
		SELECT id, name, category_id
		FROM Activities
		ORDER BY name;
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []Activity
	for rows.Next() {
		var act Activity
		err := rows.Scan(&act.ID, &act.Name, &act.CategoryID)
		if err != nil {
			return nil, err
		}
		activities = append(activities, act)
	}

	return activities, nil
}

func GetActivitiesByDay(dayID int) ([]Activity, error) {
	query := `
		SELECT a.id, a.name, a.category_id
		FROM Activities a
		INNER JOIN Day_Activities da ON a.id = da.activity_id
		WHERE da.day_id = $1
		ORDER BY a.name;
	`

	rows, err := db.DB.Query(query, dayID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []Activity
	for rows.Next() {
		var act Activity
		err := rows.Scan(&act.ID, &act.Name, &act.CategoryID)
		if err != nil {
			return nil, err
		}
		activities = append(activities, act)
	}

	return activities, nil
}

func SaveActivitiesForDay(dayID int, activityIDs []int) error {
	// First, delete existing activities for this day
	deleteQuery := `DELETE FROM Day_Activities WHERE day_id = $1`
	_, err := db.DB.Exec(deleteQuery, dayID)
	if err != nil {
		return err
	}

	// Then, insert new activities
	if len(activityIDs) == 0 {
		return nil
	}

	query := `INSERT INTO Day_Activities (day_id, activity_id) VALUES `
	for i := range activityIDs {
		if i > 0 {
			query += ", "
		}
		query += fmt.Sprintf("($1, $%d)", i+2)
	}

	args := make([]interface{}, len(activityIDs)+1)
	args[0] = dayID
	for i, activityID := range activityIDs {
		args[i+1] = activityID
	}

	_, err = db.DB.Exec(query, args...)
	return err
}

func CreateNewActivity(name string, categoryID int) error {
	query := `INSERT INTO Activities (name, category_id) VALUES ($1, $2)`
	_, err := db.DB.Exec(query, name, categoryID)
	return err
}
