package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"hcn/mymodels"
	"hcn/tests/testhelpers"
)

// runTest basic test for running endpoints test.
func runTest(t *testing.T, allTest mymodels.AllTest) {
	for _, test := range allTest {
		req, err := http.NewRequest(test.Method, test.URL, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(test.Function)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != test.StatusCode {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, test.StatusCode)
		}
		responseBody := strings.TrimSuffix(rr.Body.String(), "\n") //deleting \n last char
		if responseBody != test.ExpectedBody {
			t.Errorf("Handler returned unexpected body: got \n%v want \n%v",
				responseBody, test.ExpectedBody)
		}
	}
}

// runUpdateTest basic test for running endpoints test.
func runTestWithBody(t *testing.T, allTest mymodels.AllTest) {
	for i, test := range allTest {
		req, err := http.NewRequest(test.Method, test.URL, bytes.NewBuffer([]byte(test.Body)))
		if err != nil {
			t.Fatal(err)
		}
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
	}
}

// Teachers test
func TestGetAllTeachers(t *testing.T) {
	runTest(t, testhelpers.CasesGetAllTeachers())
}

func TestGetTeacher(t *testing.T) {
	runTest(t, testhelpers.CasesGetTeacher())
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

// Students test
func TestGetAllStudents(t *testing.T) {
	runTest(t, testhelpers.CasesGetAllStudents())
}

func TestGetStudent(t *testing.T) {
	runTest(t, testhelpers.CasesGetStudent())
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

// Courses test
func TestGetAllCourses(t *testing.T) {
	runTest(t, testhelpers.CasesGetAllCourses())
}

func TestGetCourse(t *testing.T) {
	runTest(t, testhelpers.CasesGetCourse())
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
