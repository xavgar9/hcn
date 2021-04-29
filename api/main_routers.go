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
	"net/http"

	"github.com/gorilla/mux"
)

// PingPong checks if the server is ok
func PingPong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}

// MainRouters are the collection of all URLs for the Main App.
func MainRouters(router *mux.Router) {
	router.HandleFunc("/", PingPong).Methods("GET")

	// Authentication URLs
	router.HandleFunc("/Authentication/Login", authentication.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/Authentication/IsValid", authentication.IsValid).Methods("POST", "OPTIONS")

	// Teachers URLs
	router.HandleFunc("/Teachers/GetAllTeachers", teachers.GetAllTeachers).Methods("GET")
	router.HandleFunc("/Teachers/GetTeacher", teachers.GetTeacher).Methods("GET")
	router.HandleFunc("/Teachers/UpdateTeacher", teachers.UpdateTeacher).Methods("PUT", "OPTIONS")
	router.HandleFunc("/Teachers/CreateTeacher", teachers.CreateTeacher).Methods("POST", "OPTIONS")
	router.HandleFunc("/Teachers/DeleteTeacher", teachers.DeleteTeacher).Methods("DELETE", "OPTIONS")

	// Students URLs
	router.HandleFunc("/Students/GetAllStudents", students.GetAllStudents).Methods("GET")
	router.HandleFunc("/Students/GetStudent", students.GetStudent).Methods("GET")
	router.HandleFunc("/Students/UpdateStudent", students.UpdateStudent).Methods("PUT", "OPTIONS")
	router.HandleFunc("/Students/CreateStudent", students.CreateStudent).Methods("POST", "OPTIONS")
	router.HandleFunc("/Students/DeleteStudent", students.DeleteStudent).Methods("DELETE", "OPTIONS")

	// Courses URLs
	router.HandleFunc("/Courses/GetAllCourses", courses.GetAllCourses).Methods("GET")
	router.HandleFunc("/Courses/GetCourse", courses.GetCourse).Methods("GET")
	router.HandleFunc("/Courses/UpdateCourse", courses.UpdateCourse).Methods("PUT", "OPTIONS")
	router.HandleFunc("/Courses/CreateCourse", courses.CreateCourse).Methods("POST", "OPTIONS")
	router.HandleFunc("/Courses/DeleteCourse", courses.DeleteCourse).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/Courses/AddHCN", courses.AddHCN).Methods("POST", "OPTIONS")
	router.HandleFunc("/Courses/GetAllHCN", courses.GetAllHCNCourse).Methods("GET")
	router.HandleFunc("/Courses/RemoveHCN", courses.RemoveHCN).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/Courses/VisibilityHCN", courses.VisibilityHCN).Methods("POST", "OPTIONS")

	router.HandleFunc("/Courses/AddClinicalCase", courses.AddClinicalCase).Methods("POST", "OPTIONS")
	router.HandleFunc("/Courses/GetAllClinicalCases", courses.GetAllClinicalCases).Methods("GET")
	router.HandleFunc("/Courses/RemoveClinicalCase", courses.RemoveClinicalCase).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/Courses/VisibilityClinicalCase", courses.VisibilityClinicalCase).Methods("POST", "OPTIONS")

	router.HandleFunc("/Courses/AddStudent", courses.AddStudent).Methods("POST", "OPTIONS")
	router.HandleFunc("/Courses/GetAllStudentsCourse", courses.GetAllStudentsCourse).Methods("GET")
	router.HandleFunc("/Courses/RemoveStudent", courses.RemoveStudent).Methods("DELETE", "OPTIONS")

	// Announcements URLs
	router.HandleFunc("/Announcements/GetAllAnnouncements", announcements.GetAllAnnouncements).Methods("GET")
	router.HandleFunc("/Announcements/GetAnnouncement", announcements.GetAnnouncement).Methods("GET")
	router.HandleFunc("/Announcements/UpdateAnnouncement", announcements.UpdateAnnouncement).Methods("PUT", "OPTIONS")
	router.HandleFunc("/Announcements/CreateAnnouncement", announcements.CreateAnnouncement).Methods("POST", "OPTIONS")
	router.HandleFunc("/Announcements/DeleteAnnouncement", announcements.DeleteAnnouncement).Methods("DELETE", "OPTIONS")

	// Activities URLs
	router.HandleFunc("/Activities/GetAllActivities", activities.GetAllActivities).Methods("GET")
	router.HandleFunc("/Activities/GetActivity", activities.GetActivity).Methods("GET")
	router.HandleFunc("/Activities/UpdateActivity", activities.UpdateActivity).Methods("PUT", "OPTIONS")
	router.HandleFunc("/Activities/CreateActivity", activities.CreateActivity).Methods("POST", "OPTIONS")
	router.HandleFunc("/Activities/DeleteActivity", activities.DeleteActivity).Methods("DELETE", "OPTIONS")

	// Feedbacks URLs
	/*
		router.HandleFunc("/Feedbacks/GetAllFeedbacks", feedbacks.GetAllFeedbacks).Methods("GET")
		router.HandleFunc("/Feedbacks/GetFeedback", feedbacks.GetFeedback).Methods("GET")
		router.HandleFunc("/Feedbacks/UpdateFeedback", feedbacks.UpdateFeedback).Methods("POST", "OPTIONS")
		router.HandleFunc("/Feedbacks/CreateFeedback", feedbacks.CreateFeedback).Methods("POST", "OPTIONS")
		router.HandleFunc("/Feedbacks/DeleteFeedback", feedbacks.DeleteFeedback).Methods("DELETE", "OPTIONS")
	*/

	// Clinical Cases URLs
	router.HandleFunc("/ClinicalCases/GetAllClinicalCases", ccases.GetAllClinicalCases).Methods("GET")
	router.HandleFunc("/ClinicalCases/GetClinicalCase", ccases.GetClinicalCase).Methods("GET")
	router.HandleFunc("/ClinicalCases/UpdateClinicalCase", ccases.UpdateClinicalCase).Methods("POST", "OPTIONS")
	router.HandleFunc("/ClinicalCases/CreateClinicalCase", ccases.CreateClinicalCase).Methods("POST", "OPTIONS")
	router.HandleFunc("/ClinicalCases/DeleteClinicalCase", ccases.DeleteClinicalCase).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/ClinicalCases/LinkHCN", ccases.LinkHCN).Methods("POST", "OPTIONS")
	router.HandleFunc("/ClinicalCases/UnlinkHCN", ccases.UnlinkHCN).Methods("DELETE", "OPTIONS")

	//router.HandleFunc("/ClinicalCases/DownloadPDF", ccases.DownloadPDF).Methods("GET")
	//router.HandleFunc("/ClinicalCases/UnlinkHCN", ccases.UnlinkHCN).Methods("DELETE", "OPTIONS")

	// HCN URLs
	router.HandleFunc("/HCN/GetAllHCN", hcn.GetAllHCN).Methods("GET")
	router.HandleFunc("/HCN/GetHCN", hcn.GetHCN).Methods("GET")
	router.HandleFunc("/HCN/UpdateHCN", hcn.UpdateHCN).Methods("POST", "OPTIONS")
	router.HandleFunc("/HCN/CreateHCN", hcn.CreateHCN).Methods("POST", "OPTIONS")
	router.HandleFunc("/HCN/DeleteHCN", hcn.DeleteHCN).Methods("DELETE", "OPTIONS")

	// HCN Mongo URLs
	router.HandleFunc("/HCN/GetAllHCNMongo", hcn.GetAllHCNMongo).Methods("GET")
	router.HandleFunc("/HCN/GetHCNMongo", hcn.GetHCNMongo).Methods("GET")
	router.HandleFunc("/HCN/UpdateHCNMongo", hcn.UpdateHCNMongo).Methods("POST", "OPTIONS")
	router.HandleFunc("/HCN/CreateHCNMongo", hcn.CreateHCNMongo).Methods("POST", "OPTIONS")
	router.HandleFunc("/HCN/DeleteHCNMongo", hcn.DeleteHCNMongo).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/HCN/DeleteAllHCNMongo", hcn.DeleteAllHCNMongo).Methods("DELETE", "OPTIONS")

	// SolvedHCN URLs
	router.HandleFunc("/SolvedHCN/CreateSolvedHCN", solvedhcn.CreateSolvedHCN).Methods("POST", "OPTIONS")
	router.HandleFunc("/SolvedHCN/GetAllSolvedHCN", solvedhcn.GetAllSolvedHCN).Methods("GET")
}
