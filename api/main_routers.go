package api

import (
	"fmt"
	"hcn/myhandlers/activities"
	"hcn/myhandlers/announcements"
	authentication "hcn/myhandlers/authentication"
	"hcn/myhandlers/ccases"
	"hcn/myhandlers/courses"
	"hcn/myhandlers/hcn"
	solvedhcn "hcn/myhandlers/solvedHCN"
	"hcn/myhandlers/students"
	"hcn/myhandlers/teachers"
	middleware "hcn/myhelpers/middlewareHelper"
	"net/http"

	"github.com/gorilla/mux"
)

// PingPong checks if the server is ok
func PingPong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}

// MainRouters are the collection of all URLs for the Main App.
func MainRouters(router *mux.Router) {
	router.HandleFunc("/", PingPong).Methods("GET", "OPTIONS")

	// Authentication URLs
	router.HandleFunc("/Authentication/Login", authentication.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/Authentication/IsValid", authentication.IsValid).Methods("POST", "OPTIONS")

	// Activities URLs
	router.HandleFunc("/Activities/GetAllActivities", middleware.Middleware(activities.GetAllActivities)).Methods("GET", "OPTIONS")
	router.HandleFunc("/Activities/GetActivity", middleware.Middleware(activities.GetActivity)).Methods("GET", "OPTIONS")
	router.HandleFunc("/Activities/UpdateActivity", middleware.Middleware(activities.UpdateActivity)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/Activities/CreateActivity", middleware.Middleware(activities.CreateActivity)).Methods("POST", "OPTIONS")
	router.HandleFunc("/Activities/DeleteActivity", middleware.Middleware(activities.DeleteActivity)).Methods("DELETE", "OPTIONS")

	// Students URLs
	router.HandleFunc("/Students/GetAllStudents", middleware.Middleware(students.GetAllStudents)).Methods("GET", "OPTIONS")
	router.HandleFunc("/Students/GetStudent", middleware.Middleware(students.GetStudent)).Methods("GET", "OPTIONS")
	router.HandleFunc("/Students/UpdateStudent", middleware.Middleware(students.UpdateStudent)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/Students/CreateStudent", middleware.Middleware(students.CreateStudent)).Methods("POST", "OPTIONS")
	router.HandleFunc("/Students/DeleteStudent", middleware.Middleware(students.DeleteStudent)).Methods("DELETE", "OPTIONS")

	// Teachers URLs
	router.HandleFunc("/Teachers/GetAllTeachers", middleware.Middleware(teachers.GetAllTeachers)).Methods("GET", "OPTIONS")
	router.HandleFunc("/Teachers/GetTeacher", middleware.Middleware(teachers.GetTeacher)).Methods("GET", "OPTIONS")
	router.HandleFunc("/Teachers/UpdateTeacher", middleware.Middleware(teachers.UpdateTeacher)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/Teachers/CreateTeacher", middleware.Middleware(teachers.CreateTeacher)).Methods("POST", "OPTIONS")
	router.HandleFunc("/Teachers/DeleteTeacher", middleware.Middleware(teachers.DeleteTeacher)).Methods("DELETE", "OPTIONS")

	// Courses URLs
	router.HandleFunc("/Courses/GetAllCourses", middleware.Middleware(courses.GetAllCourses)).Methods("GET", "OPTIONS")
	router.HandleFunc("/Courses/GetCourse", middleware.Middleware(courses.GetCourse)).Methods("GET", "OPTIONS")
	router.HandleFunc("/Courses/UpdateCourse", middleware.Middleware(courses.UpdateCourse)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/Courses/CreateCourse", middleware.Middleware(courses.CreateCourse)).Methods("POST", "OPTIONS")
	router.HandleFunc("/Courses/DeleteCourse", middleware.Middleware(courses.DeleteCourse)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/Courses/AddHCN", middleware.Middleware(courses.AddHCN)).Methods("POST", "OPTIONS")
	router.HandleFunc("/Courses/GetAllHCN", middleware.Middleware(courses.GetAllHCNCourse)).Methods("GET", "OPTIONS")
	router.HandleFunc("/Courses/RemoveHCN", middleware.Middleware(courses.RemoveHCN)).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/Courses/VisibilityHCN", middleware.Middleware(courses.VisibilityHCN)).Methods("POST", "OPTIONS")

	router.HandleFunc("/Courses/AddClinicalCase", middleware.Middleware(courses.AddClinicalCase)).Methods("POST", "OPTIONS")
	router.HandleFunc("/Courses/GetAllClinicalCases", middleware.Middleware(courses.GetAllClinicalCases)).Methods("GET", "OPTIONS")
	router.HandleFunc("/Courses/RemoveClinicalCase", middleware.Middleware(courses.RemoveClinicalCase)).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/Courses/VisibilityClinicalCase", middleware.Middleware(courses.VisibilityClinicalCase)).Methods("POST", "OPTIONS")

	router.HandleFunc("/Courses/AddStudent", middleware.Middleware(courses.AddStudent)).Methods("POST", "OPTIONS")
	router.HandleFunc("/Courses/GetAllStudentsCourse", middleware.Middleware(courses.GetAllStudentsCourse)).Methods("GET", "OPTIONS")
	router.HandleFunc("/Courses/RemoveStudent", middleware.Middleware(courses.RemoveStudent)).Methods("DELETE", "OPTIONS")

	// Announcements URLs
	router.HandleFunc("/Announcements/GetAllAnnouncements", middleware.Middleware(announcements.GetAllAnnouncements)).Methods("GET", "OPTIONS")
	router.HandleFunc("/Announcements/GetAnnouncement", middleware.Middleware(announcements.GetAnnouncement)).Methods("GET", "OPTIONS")
	router.HandleFunc("/Announcements/UpdateAnnouncement", middleware.Middleware(announcements.UpdateAnnouncement)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/Announcements/CreateAnnouncement", middleware.Middleware(announcements.CreateAnnouncement)).Methods("POST", "OPTIONS")
	router.HandleFunc("/Announcements/DeleteAnnouncement", middleware.Middleware(announcements.DeleteAnnouncement)).Methods("DELETE", "OPTIONS")

	// Feedbacks URLs
	/*
		router.HandleFunc("/Feedbacks/GetAllFeedbacks", feedbacks.GetAllFeedbacks)).Methods("GET")
		router.HandleFunc("/Feedbacks/GetFeedback", feedbacks.GetFeedback)).Methods("GET")
		router.HandleFunc("/Feedbacks/UpdateFeedback", feedbacks.UpdateFeedback)).Methods("POST", "OPTIONS")
		router.HandleFunc("/Feedbacks/CreateFeedback", feedbacks.CreateFeedback)).Methods("PUT", "OPTIONS")
		router.HandleFunc("/Feedbacks/DeleteFeedback", feedbacks.DeleteFeedback)).Methods("DELETE", "OPTIONS")
	*/

	// Clinical Cases URLs
	router.HandleFunc("/ClinicalCases/GetAllClinicalCases", middleware.Middleware(ccases.GetAllClinicalCases)).Methods("GET", "OPTIONS")
	router.HandleFunc("/ClinicalCases/GetClinicalCase", middleware.Middleware(ccases.GetClinicalCase)).Methods("GET", "OPTIONS")
	router.HandleFunc("/ClinicalCases/UpdateClinicalCase", middleware.Middleware(ccases.UpdateClinicalCase)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/ClinicalCases/CreateClinicalCase", middleware.Middleware(ccases.CreateClinicalCase)).Methods("POST", "OPTIONS")
	router.HandleFunc("/ClinicalCases/DeleteClinicalCase", middleware.Middleware(ccases.DeleteClinicalCase)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/ClinicalCases/LinkHCN", middleware.Middleware(ccases.LinkHCN)).Methods("POST", "OPTIONS")
	router.HandleFunc("/ClinicalCases/UnlinkHCN", middleware.Middleware(ccases.UnlinkHCN)).Methods("DELETE", "OPTIONS")

	//router.HandleFunc("/ClinicalCases/DownloadPDF", ccases.DownloadPDF)).Methods("GET")
	//router.HandleFunc("/ClinicalCases/UnlinkHCN", ccases.UnlinkHCN)).Methods("DELETE", "OPTIONS")

	// HCN URLs
	router.HandleFunc("/HCN/GetAllHCN", middleware.Middleware(hcn.GetAllHCN)).Methods("GET", "OPTIONS")
	router.HandleFunc("/HCN/GetHCN", middleware.Middleware(hcn.GetHCN)).Methods("GET", "OPTIONS")
	router.HandleFunc("/HCN/UpdateHCN", middleware.Middleware(hcn.UpdateHCN)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/HCN/CreateHCN", middleware.Middleware(hcn.CreateHCN)).Methods("POST", "OPTIONS")
	router.HandleFunc("/HCN/DeleteHCN", middleware.Middleware(hcn.DeleteHCN)).Methods("DELETE", "OPTIONS")

	// HCN Mongo URLs
	router.HandleFunc("/HCN/GetAllHCNMongo", middleware.Middleware(hcn.GetAllHCNMongo)).Methods("GET", "OPTIONS")
	router.HandleFunc("/HCN/GetHCNMongo", middleware.Middleware(hcn.GetHCNMongo)).Methods("GET", "OPTIONS")
	router.HandleFunc("/HCN/UpdateHCNMongo", middleware.Middleware(hcn.UpdateHCNMongo)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/HCN/CreateHCNMongo", middleware.Middleware(hcn.CreateHCNMongo)).Methods("POST", "OPTIONS")
	router.HandleFunc("/HCN/DeleteHCNMongo", middleware.Middleware(hcn.DeleteHCNMongo)).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/HCN/DeleteAllHCNMongo", middleware.Middleware(hcn.DeleteAllHCNMongo)).Methods("DELETE", "OPTIONS")

	// SolvedHCN URLs
	router.HandleFunc("/SolvedHCN/CreateSolvedHCN", middleware.Middleware(solvedhcn.CreateSolvedHCN)).Methods("POST", "OPTIONS")
	router.HandleFunc("/SolvedHCN/GetAllSolvedHCN", middleware.Middleware(solvedhcn.GetAllSolvedHCN)).Methods("GET", "OPTIONS")
	router.HandleFunc("/SolvedHCN/UpdateSolvedHCN", middleware.Middleware(solvedhcn.UpdateSolvedHCN)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/SolvedHCN/DeleteSolvedHCN", middleware.Middleware(solvedhcn.DeleteSolvedHCN)).Methods("DELETE", "OPTIONS")
}
