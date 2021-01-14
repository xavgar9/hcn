package api

import (
	"hcn/config"
	"hcn/myhandlers/activities"
	"hcn/myhandlers/announcements"
	"hcn/myhandlers/ccases"
	"hcn/myhandlers/courses"
	"hcn/myhandlers/feedbacks"
	"hcn/myhandlers/hcn"
	"hcn/myhandlers/students"
	"hcn/myhandlers/teachers"
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
	router.HandleFunc("/Teachers/GetAllTeachers", teachers.GetAllTeachers).Methods("GET")
	router.HandleFunc("/Teachers/GetTeacher/{id}", teachers.GetTeacher).Methods("GET")
	router.HandleFunc("/Teachers/UpdateTeacher", teachers.UpdateTeacher).Methods("POST")
	router.HandleFunc("/Teachers/CreateTeacher", teachers.CreateTeacher).Methods("POST")
	router.HandleFunc("/Teachers/DeleteTeacher", teachers.DeleteTeacher).Methods("DELETE")

	// Students URLs
	router.HandleFunc("/Students/GetAllStudents", students.GetAllStudents).Methods("GET")
	router.HandleFunc("/Students/GetStudent/{id}", students.GetStudent).Methods("GET")
	router.HandleFunc("/Students/UpdateStudent", students.UpdateStudent).Methods("POST")
	router.HandleFunc("/Students/CreateStudent", students.CreateStudent).Methods("POST")
	router.HandleFunc("/Students/DeleteStudent", students.DeleteStudent).Methods("DELETE")

	// Courses URLs
	router.HandleFunc("/Courses/GetAllCourses", courses.GetAllCourses).Methods("GET")
	router.HandleFunc("/Courses/GetCourse/{id}", courses.GetCourse).Methods("GET")
	router.HandleFunc("/Courses/UpdateCourse", courses.UpdateCourse).Methods("POST")
	router.HandleFunc("/Courses/CreateCourse", courses.CreateCourse).Methods("POST")
	router.HandleFunc("/Courses/DeleteCourse", courses.DeleteCourse).Methods("DELETE")

	router.HandleFunc("/Courses/AddHCN", courses.AddHCN).Methods("POST")
	router.HandleFunc("/Courses/GetHCN", courses.GetHCN).Methods("GET")
	router.HandleFunc("/Courses/RemoveHCN", courses.RemoveHCN).Methods("POST")
	router.HandleFunc("/Courses/VisibilityHCN", courses.VisibilityHCN).Methods("POST")

	//router.HandleFunc("/Courses/AddClinicalCase", myhandlers.AddClinicalCAse).Methods("POST")
	//router.HandleFunc("/Courses/GetAllClinicalCases/{id}", myhandlers.GetAllClinicalCases).Methods("GET")
	//router.HandleFunc("/Courses/RemoveClinicalCase", myhandlers.RemoveClinicalCase).Methods("POST")
	//router.HandleFunc("/Courses/VisibilityClinicalCase", myhandlers.VisibilityClinicalCase).Methods("POST")

	router.HandleFunc("/Courses/AddStudent", courses.AddStudent).Methods("POST")
	router.HandleFunc("/Courses/GetAllStudentsCourse/{id}", courses.GetAllStudentsCourse).Methods("GET")
	router.HandleFunc("/Courses/RemoveStudent", courses.RemoveStudent).Methods("DELETE")

	// Announcements URLs
	router.HandleFunc("/Announcements/GetAllAnnouncements", announcements.GetAllAnnouncements).Methods("GET")
	router.HandleFunc("/Announcements/GetAnnouncement/{id}", announcements.GetAnnouncement).Methods("GET")
	router.HandleFunc("/Announcements/UpdateAnnouncement", announcements.UpdateAnnouncement).Methods("POST")
	router.HandleFunc("/Announcements/CreateAnnouncement", announcements.CreateAnnouncement).Methods("POST")
	router.HandleFunc("/Announcements/DeleteAnnouncement", announcements.DeleteAnnouncement).Methods("DELETE")

	// Activities URLs
	router.HandleFunc("/Activities/GetAllActivities", activities.GetAllActivities).Methods("GET")
	router.HandleFunc("/Activities/GetActivity/{id}", activities.GetActivity).Methods("GET")
	router.HandleFunc("/Activities/UpdateActivity", activities.UpdateActivity).Methods("POST")
	router.HandleFunc("/Activities/CreateActivity", activities.CreateActivity).Methods("POST")
	router.HandleFunc("/Activities/DeleteActivity", activities.DeleteActivity).Methods("DELETE")

	// Feedbacks URLs
	router.HandleFunc("/Feedbacks/GetAllFeedbacks", feedbacks.GetAllFeedbacks).Methods("GET")
	router.HandleFunc("/Feedbacks/GetFeedback/{id}", feedbacks.GetFeedback).Methods("GET")
	router.HandleFunc("/Feedbacks/UpdateFeedback", feedbacks.UpdateFeedback).Methods("POST")
	router.HandleFunc("/Feedbacks/CreateFeedback", feedbacks.CreateFeedback).Methods("POST")
	router.HandleFunc("/Feedbacks/DeleteFeedback", feedbacks.DeleteFeedback).Methods("DELETE")

	// Clinical Cases URLs
	router.HandleFunc("/ClinicalCases/GetAllClinicalCases", ccases.GetAllClinicalCases).Methods("GET")
	router.HandleFunc("/ClinicalCases/GetClinicalCase/{id}", ccases.GetClinicalCase).Methods("GET")
	router.HandleFunc("/ClinicalCases/UpdateClinicalCase", ccases.UpdateClinicalCase).Methods("POST")
	router.HandleFunc("/ClinicalCases/CreateClinicalCase", ccases.CreateClinicalCase).Methods("POST")
	router.HandleFunc("/ClinicalCases/DeleteClinicalCase", ccases.DeleteClinicalCase).Methods("DELETE")

	router.HandleFunc("/ClinicalCases/AddHCN", ccases.AddHCN).Methods("POST")
	router.HandleFunc("/ClinicalCases/RemoveHCN", ccases.RemoveHCN).Methods("DELETE")

	// HCN URLs
	router.HandleFunc("/HCN/GetAllHCN", hcn.GetAllHCN).Methods("GET")
	router.HandleFunc("/HCN/GetHCN/{id}", hcn.GetHCN).Methods("GET")
	router.HandleFunc("/HCN/UpdateHCN", hcn.UpdateHCN).Methods("POST")
	router.HandleFunc("/HCN/CreateHCN", hcn.CreateHCN).Methods("POST")
	router.HandleFunc("/HCN/DeleteHCN", hcn.DeleteHCN).Methods("DELETE")
}
