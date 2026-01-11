package router

import (
	"html/template"
	"net/http"

	historyController "github.com/N3moAhead/endeavor/internal/controller/history"
	"github.com/N3moAhead/endeavor/internal/controller/journal"
	historyModel "github.com/N3moAhead/endeavor/internal/model/history"
)

type Router struct {
	handler   *http.ServeMux
	templates *template.Template
}

func New() *Router {
	tpl := initTemplates()
	mux := initMux(tpl)

	return &Router{
		handler:   mux,
		templates: tpl,
	}
}

func initMux(tpl *template.Template) *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		data := journal.GetPageData()
		tpl.ExecuteTemplate(w, "index.html", data)
	})

	mux.HandleFunc("POST /{$}", func(w http.ResponseWriter, r *http.Request) {
		journal.SaveEntry(w, r)
	})

	mux.HandleFunc("GET /history", historyController.GetHistory)
	mux.HandleFunc("GET /day/detail", historyController.GetDayDetail)

	return mux
}

func initTemplates() *template.Template {
	funcMap := template.FuncMap{
		"groupActivitiesByCategory": groupActivitiesByCategory,
	}

	tpl, err := template.New("").Funcs(funcMap).ParseGlob("web/templates/*.html")
	if err != nil {
		panic(err)
	}
	return tpl
}

func groupActivitiesByCategory(activities []historyModel.ActivityWithCategory) []CategoryGroup {
	categoryMap := make(map[string]CategoryGroup)

	for _, act := range activities {
		if group, exists := categoryMap[act.CategoryName]; exists {
			group.Activities = append(group.Activities, act)
			categoryMap[act.CategoryName] = group
		} else {
			categoryMap[act.CategoryName] = CategoryGroup{
				Name:       act.CategoryName,
				Color:      act.CategoryColor,
				Activities: []historyModel.ActivityWithCategory{act},
			}
		}
	}

	var result []CategoryGroup
	for _, group := range categoryMap {
		result = append(result, group)
	}

	return result
}

type CategoryGroup struct {
	Name       string
	Color      string
	Activities []historyModel.ActivityWithCategory
}

func (r *Router) GetHandler() *http.ServeMux {
	return r.handler
}
