package history

import (
	"database/sql"
	"errors"
	"time"

	"github.com/N3moAhead/endeavor/internal/db"
	"github.com/N3moAhead/endeavor/internal/model/activity"
	"github.com/N3moAhead/endeavor/internal/model/category"
	"github.com/N3moAhead/endeavor/internal/model/mood"
)

type DayEntry struct {
	ID         int
	Date       time.Time
	Mood       *mood.Mood
	Activities []ActivityWithCategory
}

type ActivityWithCategory struct {
	activity.Activity
	CategoryName  string
	CategoryColor string
}

func GetAllDays() ([]DayEntry, error) {
	query := `
		SELECT d.id, d.date, m.id, m.name, m.description, m.day_id
		FROM Days d
		LEFT JOIN Moods m ON d.id = m.day_id
		ORDER BY d.date DESC;
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var days []DayEntry
	for rows.Next() {
		var day DayEntry
		var moodID sql.NullInt64
		var moodName, moodDescription sql.NullString
		var moodDayID sql.NullInt64

		err := rows.Scan(
			&day.ID, &day.Date,
			&moodID, &moodName, &moodDescription, &moodDayID,
		)
		if err != nil {
			return nil, err
		}

		if moodID.Valid {
			day.Mood = &mood.Mood{
				ID:          int(moodID.Int64),
				Name:        moodName.String,
				Description: moodDescription.String,
				DayID:       int(moodDayID.Int64),
			}
		}

		days = append(days, day)
	}

	return days, nil
}

func GetDayByID(dayID int) (*DayEntry, error) {
	// Get day info
	dayQuery := `
		SELECT id, date FROM Days WHERE id = $1;
	`
	var day DayEntry
	err := db.DB.QueryRow(dayQuery, dayID).Scan(&day.ID, &day.Date)
	if err != nil {
		return nil, err
	}

	// Get mood
	mood, err := mood.GetMoodForDay(dayID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	day.Mood = mood

	// Get activities with categories
	activities, err := activity.GetActivitiesByDay(dayID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	// Get all categories for mapping
	categories, err := category.GetAll()
	if err != nil {
		return nil, err
	}

	categoryMap := make(map[int]category.Category)
	for _, cat := range categories {
		categoryMap[cat.ID] = cat
	}

	for _, act := range activities {
		cat := categoryMap[act.CategoryID]
		day.Activities = append(day.Activities, ActivityWithCategory{
			Activity:      act,
			CategoryName:  cat.Name,
			CategoryColor: cat.Color,
		})
	}

	return &day, nil
}

func GetDaysByYear(year int) ([]DayEntry, error) {
	query := `
		SELECT d.id, d.date, m.id, m.name, m.description, m.day_id
		FROM Days d
		LEFT JOIN Moods m ON d.id = m.day_id
		WHERE EXTRACT(YEAR FROM d.date) = $1
		ORDER BY d.date DESC;
	`

	rows, err := db.DB.Query(query, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var days []DayEntry
	for rows.Next() {
		var day DayEntry
		var moodID sql.NullInt64
		var moodName, moodDescription sql.NullString
		var moodDayID sql.NullInt64

		err := rows.Scan(
			&day.ID, &day.Date,
			&moodID, &moodName, &moodDescription, &moodDayID,
		)
		if err != nil {
			return nil, err
		}

		if moodID.Valid {
			day.Mood = &mood.Mood{
				ID:          int(moodID.Int64),
				Name:        moodName.String,
				Description: moodDescription.String,
				DayID:       int(moodDayID.Int64),
			}
		}

		days = append(days, day)
	}

	return days, nil
}
