package journal

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/N3moAhead/endeavor/internal/model/activity"
	"github.com/N3moAhead/endeavor/internal/model/category"
	"github.com/N3moAhead/endeavor/internal/model/day"
	"github.com/N3moAhead/endeavor/internal/model/mood"
)

type Journal struct {
	DayID      int
	Mood       *mood.Mood
	Activities []activity.Activity
	Categories []category.Category
}

type PageData struct {
	Journal            *Journal
	Moods              []MoodOption
	Activities         []ActivityWithCategory
	SelectedMood       string
	SelectedActivities []int
}

type MoodOption struct {
	Name        string
	Description string
	Value       string
}

type ActivityWithCategory struct {
	activity.Activity
	CategoryName  string
	CategoryColor string
}

// Retrieves the data of the current day
// If no data has been retrieved today it
// will create a new day entry in the database
func GetToday() *Journal {
	dayID, err := day.CreateOrGetToday()
	if err != nil {
		// TODO there might be some better ways for error handling
		// than just crashing the process...
		panic(err)
	}
	todaysMood, err := mood.GetMoodForDay(dayID)
	if err != nil {
		// If just the row is not defined that's okay...
		if !errors.Is(err, sql.ErrNoRows) {
			// TODO again not that great...
			panic(err)
		}
	}

	todaysActivities, err := activity.GetActivitiesByDay(dayID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}

	categories, err := category.GetAll()
	if err != nil {
		panic(err)
	}

	return &Journal{
		DayID:      dayID,
		Mood:       todaysMood,
		Activities: todaysActivities,
		Categories: categories,
	}
}

func GetPageData() *PageData {
	journal := GetToday()

	moods := []MoodOption{
		{"Bad", "Bad", "bad"},
		{"Meh", "Meh", "meh"},
		{"Okay", "Okay", "okay"},
		{"Good", "Good", "good"},
		{"Amazing", "Amazing", "amazing"},
	}

	activities, err := activity.GetAll()
	if err != nil {
		panic(err)
	}

	categories, err := category.GetAll()
	if err != nil {
		panic(err)
	}

	categoryMap := make(map[int]category.Category)
	for _, cat := range categories {
		categoryMap[cat.ID] = cat
	}

	var activitiesWithCategory []ActivityWithCategory
	for _, act := range activities {
		cat := categoryMap[act.CategoryID]
		activitiesWithCategory = append(activitiesWithCategory, ActivityWithCategory{
			Activity:      act,
			CategoryName:  cat.Name,
			CategoryColor: cat.Color,
		})
	}

	selectedMood := ""
	if journal.Mood != nil {
		selectedMood = journal.Mood.Name
	}

	var selectedActivities []int
	for _, act := range journal.Activities {
		selectedActivities = append(selectedActivities, act.ID)
	}

	return &PageData{
		Journal:            journal,
		Moods:              moods,
		Activities:         activitiesWithCategory,
		SelectedMood:       selectedMood,
		SelectedActivities: selectedActivities,
	}
}

func SaveEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	journal := GetToday()

	// Check for new activity creation
	newActivityName := r.FormValue("new_activity")
	newActivityCategory := r.FormValue("new_activity_category")
	if newActivityName != "" && newActivityCategory != "" {
		if categoryID, err := strconv.Atoi(newActivityCategory); err == nil {
			err = activity.CreateNewActivity(newActivityName, categoryID)
			if err != nil {
				http.Error(w, "Failed to create new activity", http.StatusInternalServerError)
				return
			}
		}
	}

	// Save mood
	moodName := r.FormValue("mood")
	moodDescription := ""
	for _, mood := range []MoodOption{
		{"Bad", "Bad", "bad"},
		{"Meh", "Meh", "meh"},
		{"Okay", "Okay", "okay"},
		{"Good", "Good", "good"},
		{"Amazing", "Amazing", "amazing"},
	} {
		if mood.Value == moodName {
			moodDescription = mood.Description
			break
		}
	}

	if moodName != "" {
		err := mood.SaveMood(journal.DayID, moodName, moodDescription)
		if err != nil {
			http.Error(w, "Failed to save mood", http.StatusInternalServerError)
			return
		}
	}

	// Save activities
	activityValues := r.Form["activities"]
	var activityIDs []int
	for _, value := range activityValues {
		if id, err := strconv.Atoi(value); err == nil {
			activityIDs = append(activityIDs, id)
		}
	}

	err := activity.SaveActivitiesForDay(journal.DayID, activityIDs)
	if err != nil {
		http.Error(w, "Failed to save activities", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
