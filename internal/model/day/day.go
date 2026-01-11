package day

import "github.com/N3moAhead/endeavor/internal/db"

func CreateOrGetToday() (int, error) {
	var id int
	query := `
        INSERT INTO Days (date)
        VALUES (CURRENT_DATE)
        ON CONFLICT (date) DO UPDATE SET date = EXCLUDED.date
        RETURNING id;
    `

	err := db.DB.QueryRow(query).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
