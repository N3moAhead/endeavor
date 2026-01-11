package history

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/N3moAhead/endeavor/internal/model/history"
)

type HistoryPageData struct {
	Days        []history.DayEntry
	CurrentYear int
	Years       []int
}

type DayDetailPageData struct {
	Day *history.DayEntry
}

func GetHistory(w http.ResponseWriter, r *http.Request, templates *template.Template) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Get all days
	days, err := history.GetAllDays()
	if err != nil {
		http.Error(w, "Failed to load history", http.StatusInternalServerError)
		return
	}

	// Extract unique years
	yearMap := make(map[int]bool)
	currentYear := time.Now().Year()
	for _, day := range days {
		yearMap[day.Date.Year()] = true
	}

	var years []int
	for year := range yearMap {
		years = append(years, year)
	}

	data := HistoryPageData{
		Days:        days,
		CurrentYear: currentYear,
		Years:       years,
	}

	templates.ExecuteTemplate(w, "history.html", data)
}

func GetDayDetail(w http.ResponseWriter, r *http.Request, templates *template.Template) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Get day ID from URL parameter
	dayIDStr := r.URL.Query().Get("id")
	if dayIDStr == "" {
		http.Redirect(w, r, "/history", http.StatusSeeOther)
		return
	}

	dayID, err := strconv.Atoi(dayIDStr)
	if err != nil {
		http.Redirect(w, r, "/history", http.StatusSeeOther)
		return
	}

	// Get day details
	day, err := history.GetDayByID(dayID)
	if err != nil {
		http.Error(w, "Day not found", http.StatusNotFound)
		return
	}

	data := DayDetailPageData{
		Day: day,
	}

	templates.ExecuteTemplate(w, "day_detail.html", data)
}
