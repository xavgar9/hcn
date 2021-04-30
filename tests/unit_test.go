package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"hcn/mymodels"
	"hcn/tests/testhelpers"
)

func printCountCases(t *testing.T, cntCases int) {
	fmt.Println(" ")
	fmt.Printf("<-------------- Total cases tested: %d -------------->\n", cntCases)
	fmt.Println(" ")
}

var TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6NTAwMDMsIk5hbWUiOiJKaG9hbiBMb3phbm8gUm9qYXMiLCJFbWFpbCI6Impob2FuQGVtYWlsLmNvbSIsImV4cCI6MTYxOTc2MzY2MX0.njDKNgrX8Blm8xkd7tGZSk2erj4SWE2sppxa2YphAcE"

// runTest basic test for running endpoints test.
func runTest(t *testing.T, allTest mymodels.AllTest) {
	for i, test := range allTest {
		req, err := http.NewRequest(test.Method, test.URL, nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Token", TOKEN)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(test.Function)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != test.StatusCode {
			t.Errorf("Test #%v: Handler returned wrong status code: got %v want %v",
				i, status, test.StatusCode)
		}
		responseBody := strings.TrimSuffix(rr.Body.String(), "\n") //deleting \n last char
		if responseBody != test.ExpectedBody {
			t.Errorf("Test #%v: Handler returned unexpected body: got \n%v want \n%v",
				i, responseBody, test.ExpectedBody)
		}
		cntCases++
	}
}

// runUpdateTest basic test for running endpoints test.
func runTestWithBody(t *testing.T, allTest mymodels.AllTest) {
	for i, test := range allTest {
		req, err := http.NewRequest(test.Method, test.URL, bytes.NewBuffer([]byte(test.Body)))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Token", TOKEN)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(test.Function)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != test.StatusCode {
			t.Errorf("Test #%v: Handler returned wrong status code: got %v want %v",
				i, status, test.StatusCode)
		}
		responseBody := strings.TrimSuffix(rr.Body.String(), "\n") //deleting \n last char
		if responseBody != test.ExpectedBody {
			t.Errorf("Test #%v: Handler returned unexpected body: got \n%v want \n%v",
				i, responseBody, test.ExpectedBody)
		}
		cntCases++
	}
}

var cntCases int

// Announcements test
func TestGetAllAnnouncements(t *testing.T) {
	runTest(t, testhelpers.CasesGetAllAnnouncements())
}

func TestGetAnnouncement(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesGetAnnouncement())
}

func TestCreateAnnouncement(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesCreateAnnouncement())
}

func TestUpdateAnnouncements(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesUpdateAnnouncement())
}

func TestDeleteAnnouncement(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesDeleteAnnouncement())
}

// Students test
func TestGetAllStudents(t *testing.T) {
	runTest(t, testhelpers.CasesGetAllStudents())
}

func TestGetStudent(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesGetStudent())
}

func TestCreateStudent(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesCreateStudent())
}

func TestUpdateStudent(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesUpdateStudent())
}

func TestDeleteStudent(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesDeleteStudent())
}

// Activities test
func TestGetAllActivities(t *testing.T) {
	runTest(t, testhelpers.CasesGetAllActivities())
}

func TestGetActivity(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesGetActivity())
}

func TestCreateActivity(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesCreateActivity())
}

func TestUpdateUpdateActivity(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesUpdateActivity())
}

func TestDeleteActivity(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesDeleteActivity())
}

// Teachers test
func TestGetAllTeachers(t *testing.T) {
	runTest(t, testhelpers.CasesGetAllTeachers())
}

func TestGetTeacher(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesGetTeacher())
}

func TestCreateTeacher(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesCreateTeacher())
}

func TestUpdateTeacher(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesUpdateTeacher())
}

func TestDeleteTeacher(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesDeleteTeacher())
}

// Courses test
func TestGetAllCourses(t *testing.T) {
	runTest(t, testhelpers.CasesGetAllCourses())
}

func TestGetCourse(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesGetCourse())
}

func TestCreateCourse(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesCreateCourse())
}

func TestUpdateCourses(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesUpdateCourse())
}

func TestDeleteCourses(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesDeleteCourse())
}

// Students_Courses test
func TestGetAllStudentsCourse(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesGetAllStudentsCourse())
}

func TestAddStudent(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesAddStudent())
}

func TestRemoveStudent(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesRemoveStudent())
}

// HCN test
func TestGetAllHCN(t *testing.T) {
	runTest(t, testhelpers.CasesGetAllHCN())
}

func TestGetHCN(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesGetHCN())
}

func TestCreateHCN(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesCreateHCN())
}

func TestUpdateHCN(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesUpdateHCN())
}

func TestDeleteHCN(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesDeleteHCN())
}

// Clinical Case test
func TestGetAllClinicalCases(t *testing.T) {
	runTest(t, testhelpers.CasesGetAllClinicalCases())
}

func TestGetClinicalCase(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesGetClinicalCase())
}

func TestCreateClinicalCase(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesCreateClinicalCase())
}

func TestUpdateClinicalCase(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesUpdateClinicalCase())
}

func TestDeleteClinicalCase(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesDeleteClinicalCase())
}

// CCases_HCN test
func TestLinkHCN(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesLinkHCN())
}

func TestUnlinkHCN(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesUnlinkHCN())
}

func TestGetAllClinicalCasesCourse(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesGetAllClinicalCasesCourse())
}
func TestAddClinicalCase(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesAddClinicalCase())
}

func TestRemoveClinicalCase(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesRemoveClinicalCase())
}

func TestVisibilityClinicalCase(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesVisibilityClinicalCase())
}

// Courses_HCN test
func TestAddHCN(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesAddHCN())
}

func TestGetAllHCNCourse(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesGetAllHCNCourse())
}

func TestRemoveHCN(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesRemoveHCN())
}

func TestVisibilityHCN(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesVisibilityHCN())
}

////////////////////////////////////////////////////////////////////
/*
func TestDeleteAllHCNMongo(t *testing.T) {
	runTest(t, testhelpers.CasesDeleteAllHCNMongo())
}
func TestGetAllHCNMongo1(t *testing.T) {
	runTest(t, testhelpers.CasesGetAllHCNMongo1())
}

func TestCreateHCNMongo(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesCreateHCNMongo())
}

func TestGetAllHCNMongo2(t *testing.T) {
	runTest(t, testhelpers.CasesGetAllHCNMongo2())
}

func TestUpdateHCNMongo(t *testing.T) {
	runTestWithBody(t, testhelpers.CasesUpdateHCNMongo())
}
*/

////////////////////////////////////////////////////////////////////
func TestCountCases(t *testing.T) {
	printCountCases(t, cntCases)
}
