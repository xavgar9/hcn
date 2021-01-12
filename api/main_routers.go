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

	// Teachers URLs
	router.HandleFunc("/Teachers/GetAllTeachers", myhandlers.GetAllTeachers).Methods("GET")
	router.HandleFunc("/Teachers/GetTeacher/{id}", myhandlers.GetTeacher).Methods("GET")
	router.HandleFunc("/Teachers/UpdateTeacher", myhandlers.UpdateTeacher).Methods("POST")
	router.HandleFunc("/Teachers/CreateTeacher", myhandlers.CreateTeacher).Methods("POST")
	router.HandleFunc("/Teachers/DeleteTeacher", myhandlers.DeleteTeacher).Methods("DELETE")

	// Students URLs
	router.HandleFunc("/Students/GetAllStudents", myhandlers.GetAllStudents).Methods("GET")
	router.HandleFunc("/Students/GetStudent/{id}", myhandlers.GetStudent).Methods("GET")
	router.HandleFunc("/Students/UpdateStudent", myhandlers.UpdateStudent).Methods("POST")
	router.HandleFunc("/Students/CreateStudent", myhandlers.CreateStudent).Methods("POST")
	router.HandleFunc("/Students/DeleteStudent", myhandlers.DeleteStudent).Methods("DELETE")

	// Courses URLs
	router.HandleFunc("/Courses/GetAllCourses", myhandlers.GetAllCourses).Methods("GET")
	router.HandleFunc("/Courses/GetCourse/{id}", myhandlers.GetCourse).Methods("GET")
	router.HandleFunc("/Courses/UpdateCourse", myhandlers.UpdateCourse).Methods("POST")
	router.HandleFunc("/Courses/CreateCourse", myhandlers.CreateCourse).Methods("POST")
	router.HandleFunc("/Courses/DeleteCourse", myhandlers.DeleteCourse).Methods("DELETE")

	//router.HandleFunc("/Courses/AddHCN", myhandlers.AddHCN).Methods("POST")
	//router.HandleFunc("/Courses/GetAllHCNCourse/{id}", myhandlers.GetAllHCNCourse).Methods("GET")
	//router.HandleFunc("/Courses/RemoveHCN", myhandlers.RemoveHCN).Methods("POST")
	//router.HandleFunc("/Courses/VisibilityHCN", myhandlers.VisibilityHCN).Methods("POST")

	//router.HandleFunc("/Courses/AddClinicalCAse", myhandlers.AddClinicalCAse).Methods("POST")
	//router.HandleFunc("/Courses/GetAllClinicalCases/{id}", myhandlers.GetAllClinicalCases).Methods("GET")
	//router.HandleFunc("/Courses/RemoveClinicalCase", myhandlers.RemoveClinicalCase).Methods("POST")
	//router.HandleFunc("/Courses/VisibilityClinicalCase", myhandlers.VisibilityClinicalCase).Methods("POST")

	router.HandleFunc("/Courses/AddStudent", myhandlers.AddStudent).Methods("POST")
	router.HandleFunc("/Courses/GetAllStudentsCourse/{id}", myhandlers.GetAllStudentsCourse).Methods("GET")
	router.HandleFunc("/Courses/RemoveStudent", myhandlers.RemoveStudent).Methods("DELETE")

	// Announcements URLs
	router.HandleFunc("/Announcements/GetAllAnnouncements", myhandlers.GetAllAnnouncements).Methods("GET")
	router.HandleFunc("/Announcements/GetAnnouncement/{id}", myhandlers.GetAnnouncement).Methods("GET")
	router.HandleFunc("/Announcements/UpdateAnnouncement", myhandlers.UpdateAnnouncement).Methods("POST")
	router.HandleFunc("/Announcements/CreateAnnouncement", myhandlers.CreateAnnouncement).Methods("POST")
	router.HandleFunc("/Announcements/DeleteAnnouncement", myhandlers.DeleteAnnouncement).Methods("DELETE")

	// Activities URLs
	router.HandleFunc("/Activities/GetAllActivities", myhandlers.GetAllActivities).Methods("GET")
	router.HandleFunc("/Activities/GetActivity/{id}", myhandlers.GetActivity).Methods("GET")
	router.HandleFunc("/Activities/UpdateActivity", myhandlers.UpdateActivity).Methods("POST")
	router.HandleFunc("/Activities/CreateActivity", myhandlers.CreateActivity).Methods("POST")
	router.HandleFunc("/Activities/DeleteActivity", myhandlers.DeleteActivity).Methods("DELETE")

	// Feedbacks URLs
	router.HandleFunc("/Feedbacks/GetAllFeedbacks", myhandlers.GetAllFeedbacks).Methods("GET")
	router.HandleFunc("/Feedbacks/GetFeedback/{id}", myhandlers.GetFeedback).Methods("GET")
	router.HandleFunc("/Feedbacks/UpdateFeedback", myhandlers.UpdateFeedback).Methods("POST")
	router.HandleFunc("/Feedbacks/CreateFeedback", myhandlers.CreateFeedback).Methods("POST")
	router.HandleFunc("/Feedbacks/DeleteFeedback", myhandlers.DeleteFeedback).Methods("DELETE")

	// Clinical Cases URLs
	router.HandleFunc("/ClinicalCases/GetAllClinicalCases", myhandlers.GetAllClinicalCases).Methods("GET")
	router.HandleFunc("/ClinicalCases/GetClinicalCase/{id}", myhandlers.GetClinicalCase).Methods("GET")
	router.HandleFunc("/ClinicalCases/UpdateClinicalCase", myhandlers.UpdateClinicalCase).Methods("POST")
	router.HandleFunc("/ClinicalCases/CreateClinicalCase", myhandlers.CreateClinicalCase).Methods("POST")
	router.HandleFunc("/ClinicalCases/DeleteClinicalCase", myhandlers.DeleteClinicalCase).Methods("DELETE")

	router.HandleFunc("/ClinicalCases/AddHCN", myhandlers.AddHCN).Methods("POST")
	router.HandleFunc("/ClinicalCases/RemoveHCN", myhandlers.RemoveHCN).Methods("DELETE")

	// HCN URLs
	router.HandleFunc("/HCN/GetAllHCN", myhandlers.GetAllHCN).Methods("GET")
	router.HandleFunc("/HCN/GetHCN/{id}", myhandlers.GetHCN).Methods("GET")
	router.HandleFunc("/HCN/UpdateHCN", myhandlers.UpdateHCN).Methods("POST")
	router.HandleFunc("/HCN/CreateHCN", myhandlers.CreateHCN).Methods("POST")
	router.HandleFunc("/HCN/DeleteHCN", myhandlers.DeleteHCN).Methods("DELETE")
}
