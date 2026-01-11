package seed

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Init(db *sql.DB) {
	DB = db
}

func SeedCategories() error {
	query := `
		INSERT INTO Categories (name, description, color) VALUES
		('Sports', 'Physical activities and exercises', '#FF6B6B'),
		('Work', 'Work-related activities and tasks', '#4ECDC4'),
		('Social', 'Social interactions and relationships', '#45B7D1'),
		('Learning', 'Educational activities and skill development', '#96CEB4'),
		('Relaxation', 'Relaxation and self-care activities', '#FFEAA7'),
		('Entertainment', 'Entertainment and leisure activities', '#DDA0DD')
		ON CONFLICT DO NOTHING;
	`

	_, err := DB.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Categories seeded successfully")
	return nil
}

func SeedActivities() error {
	query := `
		INSERT INTO Activities (name, category_id) VALUES
		-- Sports
		('Running', 1), ('Walking', 1), ('Cycling', 1), ('Swimming', 1), ('Yoga', 1), ('Gym', 1),
		-- Work
		('Meetings', 2), ('Programming', 2), ('Answering emails', 2), ('Project work', 2), ('Planning', 2),
		-- Social
		('Meeting friends', 3), ('Visiting family', 3), ('Phone calls', 3), ('Chatting', 3), ('Date', 3),
		-- Learning
		('Reading books', 4), ('Online courses', 4), ('Language learning', 4), ('New skills', 4), ('Watching tutorials', 4),
		-- Relaxation
		('Meditating', 5), ('Listening to music', 5), ('Taking a bath', 5), ('Napping', 5), ('Coffee break', 5),
		-- Entertainment
		('Watching movies', 6), ('Watching series', 6), ('Gaming', 6), ('YouTube', 6), ('Podcasts', 6)
		ON CONFLICT DO NOTHING;
	`

	_, err := DB.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Activities seeded successfully")
	return nil
}

func SeedAll() error {
	if err := SeedCategories(); err != nil {
		return err
	}
	return SeedActivities()
}
