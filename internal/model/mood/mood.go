package mood

import (
	"github.com/N3moAhead/endeavor/internal/db"
)

type Mood struct {
	ID          int
	Name        string
	Description string
	DayID       int
}

func GetMoodForDay(dayId int) (*Mood, error) {
	query := `
		SELECT id, name, description, day_id
		FROM Moods
		WHERE day_id = $1;
	`

	mood := Mood{}
	err := db.DB.QueryRow(query, dayId).Scan(&mood.ID, &mood.Name, &mood.Description, &mood.DayID)
	if err != nil {
		return nil, err
	}

	return &mood, nil
}

func SaveMood(dayID int, name, description string) error {
	query := `
		INSERT INTO Moods (name, description, day_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (day_id) DO UPDATE SET
			name = EXCLUDED.name,
			description = EXCLUDED.description;
	`

	_, err := db.DB.Exec(query, name, description, dayID)
	return err
}

func GetAllMoods() ([]Mood, error) {
	query := `
		SELECT id, name, description, day_id
		FROM Moods
		ORDER BY day_id DESC;
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moods []Mood
	for rows.Next() {
		var mood Mood
		err := rows.Scan(&mood.ID, &mood.Name, &mood.Description, &mood.DayID)
		if err != nil {
			return nil, err
		}
		moods = append(moods, mood)
	}

	return moods, nil
}
