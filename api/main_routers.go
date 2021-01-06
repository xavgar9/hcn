package api

import (
	"hcn/config"
	"hcn/myhandlers"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

// contextData are the most widely use common variables for each pages to load.
type contextData map[string]interface{}

// Home function is to render the homepage page.
func Home(w http.ResponseWriter, r *http.Request) {
	//tmpl := template.Must(template.ParseFiles(config.SiteRootTemplate+"front/index.html", config.SiteHeaderTemplate, config.SiteFooterTemplate))
	tmpl := template.Must(template.ParseFiles(config.SiteRootTemplate + "front/index.html"))

	data := contextData{
		"PageTitle":    config.SiteFullName,
		"PageMetaDesc": config.SiteSlogan,
		"CanonicalURL": r.RequestURI,
		"CsrfToken":    csrf.Token(r),
		"Settings":     config.SiteSettings,
	}
	tmpl.Execute(w, data)
}

/*
func testPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(config.SiteRootTemplate + "front/index.html"))
	data := contextData{
		"PageTitle":    "POST",
		"PageMetaDesc": config.SiteSlogan,
		"CanonicalURL": r.RequestURI,
		"CsrfToken":    csrf.Token(r),
		"Settings":     config.SiteSettings,
	}
	tmpl.Execute(w, data)
}
*/

// MainRouters are the collection of all URLs for the Main App.
func MainRouters(router *mux.Router) {
	router.HandleFunc("/", Home).Methods("GET")

	// Students URLs
	router.HandleFunc("/Students/GetStudents", myhandlers.GetStudents).Methods("GET")
	router.HandleFunc("/Students/GetStudent/{id}", myhandlers.GetStudent).Methods("GET")
	router.HandleFunc("/Students/UpdateStudent", myhandlers.UpdateStudent).Methods("POST")
	router.HandleFunc("/Students/CreateStudent", myhandlers.CreateStudent).Methods("POST")
	router.HandleFunc("/Students/DeleteStudent", myhandlers.DeleteStudent).Methods("DELETE")

	// Teachers URLs
	router.HandleFunc("/Teachers/GetTeachers", myhandlers.GetTeachers).Methods("GET")
	router.HandleFunc("/Teachers/GetTeacher/{id}", myhandlers.GetTeacher).Methods("GET")
	router.HandleFunc("/Teachers/UpdateTeacher", myhandlers.UpdateTeacher).Methods("POST")
	router.HandleFunc("/Teachers/CreateTeacher", myhandlers.CreateTeacher).Methods("POST")
	router.HandleFunc("/Teachers/DeleteTeacher", myhandlers.DeleteTeacher).Methods("DELETE")
}
