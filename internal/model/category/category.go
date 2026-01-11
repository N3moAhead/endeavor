package category

import (
	"github.com/N3moAhead/endeavor/internal/db"
)

type Category struct {
	ID          int
	Name        string
	Description string
	Color       string
}

func GetAll() ([]Category, error) {
	query := `
		SELECT id, name, description, color
		FROM Categories
		ORDER BY name;
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var cat Category
		err := rows.Scan(&cat.ID, &cat.Name, &cat.Description, &cat.Color)
		if err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}

	return categories, nil
}
